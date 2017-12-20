import asyncio

'''
Python中的协程允许异常在协程调用栈(yield from)中传递，在协程挂起的地方捕获到异常状态。
事件循环可以合理的捕获以及传递异常。
'''

@asyncio.coroutine
def A():
    raise Exception('Something went wrong in A!')

@asyncio.coroutine
def B():
    a = yield from A()
    yield a + 1

@asyncio.coroutine
def C():
    try:
        b = yield from B()
        print(b)
    except Exception as e:
        print('C got exception:', e)

def main():
    loop = asyncio.get_event_loop()
    loop.run_until_complete(C())

if __name__ == '__main__':
    main()
