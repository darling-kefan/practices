class DemoException(Exception):
    pass

def demo_exc_handling():
    print('Starting...')
    while True:
        try:
            y = yield
        except DemoException as e:
            print(e)
        else:
            print(y)

    raise(RuntimeError('hello world'))

if __name__ == '__main__':
    dh = demo_exc_handling()
    # 预激协程
    dh.send(None)
    dh.send(10)
    dh.throw(DemoException('demo exception'))
    dh.send(12)
    dh.throw(RuntimeError('hel'))
    dh.send(13)
    dh.send(14)
