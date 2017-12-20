from events_coroutine import EventLoop
from events_coroutine import print_every
import sys

def fib(n):
    if n <= 1:
        yield n
    else:
        a = yield fib(n - 1)
        b = yield fib(n - 2)
        yield a + b

def read_input(loop):
    while True:
        line = yield sys.stdin
        n = int(line)
        fib_n = yield fib(n)
        print('fib({}) = {}'.format(n, fib_n))

def main():
    loop = EventLoop()
    hello_task = print_every('Hello world!', 3)
    fib_task = read_input(loop)
    loop.schedule(hello_task)
    loop.schedule(fib_task)
    loop.run_forever()

if __name__ == "__main__":
    main()
