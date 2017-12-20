'''
该例子中，我们写的程序是顺序执行的，虽然使用了协程，但互不相关的协程并没有完美地并发。
也就是说，协程中的每一行代码都依赖于前一行的执行完毕。有时候我们需要一些互不相关的协程并发执行、等待它们的完成结果，并不在意它们的执行顺序。

协程可以让我们用同步的方式编写异步的代码，但是对于处理互不相关的任务不论是完成后马上处理抑或是最后统一处理，回调的方式看上去都是最好的选择。
'''

import asyncio
import json

host = 'api.ipify.org'
request_headers = {'User-Agent': 'python/3.4',
                   'host': host,
                   'Accept': 'application/json',
                   'Accept-Charset': 'UTF-8'}

@asyncio.coroutine
def write_headers(writer):
    for key, value in request_headers.items():
        writer.write((key + ':' + value + '\r\n').encode())
    writer.write(b'\r\n')
    yield from writer.drain()

@asyncio.coroutine
def read_headers(reader):
    response_headers = {}
    while True:
        line_bytes = yield from reader.readline()
        line = line_bytes.decode().strip()
        if not line:
            break
        key, value = line.split(':', 1)
        response_headers[key.strip()] = value.strip()
    return response_headers

@asyncio.coroutine
def get_my_ip_address(verbose):
    reader, writer = yield from asyncio.open_connection(host, 80)
    writer.write(b'GET /?format=json HTTP/1.1\r\n')
    yield from write_headers(writer)
    status_line = yield from reader.readline()
    status_line = status_line.decode().strip()
    http_version, status_code, status = status_line.split(' ')
    if verbose:
        print('Got status {} {}'.format(status_code, status))
    response_headers = yield from read_headers(reader)
    if verbose:
        print('Response headers:')
        for key, value in response_headers.items():
            print(key + ': ' + value)
    # Assume the content length is sent by the server, which is the case
    # with ipify
    content_length = int(response_headers['Content-Length'])
    response_body_bytes = yield from reader.read(content_length)
    response_body = response_body_bytes.decode()
    response_object = json.loads(response_body)
    writer.close()
    return response_object['ip']

@asyncio.coroutine
def print_my_ip_address(verbose):
    try:
        ip_address = yield from get_my_ip_address(verbose)
        print('My IP address is:')
        print(ip_address)
    except Exception as e:
        print('Error: ', e)

def main():
    loop = asyncio.get_event_loop()
    try:
        loop.run_until_complete(print_my_ip_address(verbose=True))
    finally:
        loop.close()

if __name__ == '__main__':
    main()

