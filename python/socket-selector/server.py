#import socket
#
## 创建socket对象，绑定IP端口，监听
#sk = socket.socket()
#sk.bind(('127.0.0.1', 1559))
#sk.listen(5)
#
## 循环接受每一个连接池中的连接
#while True:
#    # 接受客户端连接
#    conn, address = sk.accept()
#    # 向客户端发送欢迎信息
#    conn.sendall(bytes('hello', encoding='utf8'))
#
#    # 进入收发消息的循环中
#    while True:
#
#        # Windows客户端在异常断开后抛出异常，这里是处理Windows的断开情况
#        try:
#            recv = conn.recv(1024)
#
#            # Linux客户端断开recv会是空值，这里处理Linux的断开情况
#            if not recv:
#                break
#
#            # 这里处理客户端主动发出断开请求的情况
#            if str(recv, encoding='utf-8') == 'exit':
#                break
#        except Exception as ex:
#            break
#
#        # 向客户端发送数据
#        conn.sendall(recv)

import socket
import select

# 创建socket对象，绑定IP端口，监听
sk = socket.socket()
sk.bind(('127.0.0.1', 1559))
sk.listen(5)

inputs = [sk]
outputs = []
while True:
    # 监听服务端socket对象sk
    rList, wList, e = select.select(inputs, outputs, [], 1)
    print("---" * 20)
    print("select当前监听inputs对象的数量>", len(inputs), " | 发生变化的socket数量>", len(rList))
    print("select当前监听outputs对象的数量>", len(outputs), " | 需要回复客户端消息的数量>", len(wList))

    # 遍历rList中的每一个socket对象
    # 目前rList中只出现服务端的socket对象
    for s in rList:
        # 判断socket对象如果是服务端socket对象的话
        if s == sk:
            conn, address = s.accept()
            # conn也是一个socket对象
            # 当服务端socket接收到客户的请求后，会分配一个新的socket对象专门用来和
            # 这个客户端进行通信

            # 当服务端分配新的socket对象给新连接进来的客户端的时候
            # 我们也需要监听这个客户端的socket对象是否会发生变化
            # 一旦发生变化，意味着客户端向服务端发来了消息
            inputs.append(conn)
            conn.sendall(bytes('hello', encoding='utf8'))
        # 其它的就是客户端的socket对象了
        else:
            try:
                msg = s.recv(1024)
                print(type(msg), msg)

                # 意味着客户端给服务端发送消息了
                if not msg:
                    raise Exception('客户端已断开连接')
                else:
                    # 将消息发送给客户端
                    outputs.append(s)
                    outputs.append(s)
                    outputs.append(s)
                    print(msg)

            except Exception as ex:
                # Windows平台下的处理
                inputs.remove(s)

    # 遍历wList(遍历给服务端发送过消息的客户端)
    for s in wList:
        # 给所有客户端统一回复内容
        s.sendall(bytes('server response', encoding='utf8'))

        # 回复完成后，一定要将outputs中该socket对象移除
        outputs.remove(s)
