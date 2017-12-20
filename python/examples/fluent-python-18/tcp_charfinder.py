import sys
import asyncio

from charfinder import UnicodeNameIndex # 1

CRLF = b'\r\n'
PROMPT = b'?> '

index = UnicodeNameIndex() # 2

@asyncio.coroutine
def handle_queries(reader, writer): # 3
    while True: # 4
        writer.write(PROMPT) # 不能用yield from! 5
        yield from writer.drain() # 必须使用yield from! 6
        data = yield from reader.readline() # 7
        try:
            query = data.decode().strip()
        except UnicodeDecodeError: # 8
            query = '\x00'
        client = writer.get_extra_info('peername') # 9
        print('Received from {}: {!r}'.format(client, query)) # 10
        if query:
            if ord(query[:1]) < 32: # 11
                break
            lines = list(index.find_description_strs(query)) # 12
            if lines:
                writer.writelines(line.encode() + CRLF for line in lines) # 13
                writer.write(index.status(query, len(lines)).encode() + CRLF) # 14

            yield from writer.drain() # 15
            print('Sent {} results'.format(len(lines))) # 16

    print('Close the client socket') # 17
    writer.close() # 18

# 注意，run_until_complete方法的参数是一个协程(start_server方法返回的结果)或一个Future对象(server.wait_closed方法返回的结果)。如果传给run_until_complete方法的参数是协程，会把协程包装在Task对象中。
def main(address='127.0.0.1', port=2323): # 1
    port = int(port)
    loop = asyncio.get_event_loop()
    server_coro = asyncio.start_server(handle_queries, address, port,
                                       loop=loop) # 2
    server = loop.run_until_complete(server_coro) # 驱动server_coro协程，启动服务器(server)
    host = server.sockets[0].getsockname() # 4
    print('Serving on {}. Hit CTRL-C to stop.'.format(host))
    try:
        loop.run_forever() # 运行事件循环；main函数在这里阻塞，直到在服务器的控制台中按CTRL-C关闭。
    except KeyboardInterrupt: # 按CTRL-C键
        pass

    print('Server shutting down.')
    server.close() # 7
    # server.wait_closed()方法返回一个期物；调用loop.run_until_complete方法，运行期物。
    loop.run_until_complete(server.wait_closed())
    loop.close()

if __name__ == '__main__':
    main(*sys.argv[1:]) # 10
