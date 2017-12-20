'''
协程也是一个函数，它在返回的同时，还可以保存返回前的运行上下文(本地变量，以及下一条指令)，需要的时候可以重新加载上下文从上次离开的下一条命令继续执行。这种方式的return一般叫做yielding。

Python中yield是一个关键词，它可以用来创建协程。
1.当调用 yield value 的时候，这个value就被返回给调用方，同时CPU控制权也交给协程的调用方。调用yield之后，如何想要重新返回协程，需要调用Python中内置的next方法。
2.当调用 y = yield x 的时候，x被返回给调用方。要继续返回协程上下文，调用方需要再执行协程的send方法。在这个例子中，给send方法的参数会被传入协程作为这个表达式的值(在本例中，这个值会被y接收到)。

这意味着我们可以用协程来写异步代码，当程序等待异步操作的时候，只需要使用yield把控制权交出去就行了，当异步操作完成了再进入协程继续执行。
'''

from bisect import insort
from collections import deque
from fib import timed_fib
from functools import partial
from time import time
import selectors
import sys
import types

class sleep_for_seconds(object):
    """
    Yield an object of this type from a coroutine to have it "sleep" for
    the given numbers of seconds.
    """
    def __init__(self, wait_time):
        self._wait_time = wait_time

class EventLoop(object):
    """
    Implements a simplified coroutine-based event loop as a demonstration.
    Very similar to the "Tramploline" example in PEP 342, with exception
    handling taken out for simplicity, and selectors added to handle file
    to IO.
    """
    def __init__(self, *tasks):
        self._running = False
        self._selector = selectors.DefaultSelector()

        # Queue of functions scheduled to run
        self._tasks = deque(tasks)

        # (coroutine, stack) pair of tasks waiting for input from stdin
        self._tasks_waiting_on_stdin = []

        # List of (time_to_run, task) pairs, in sorted order
        self._timers = []

        # Register for polling stdin for input to read
        self._selector.register(sys.stdin, selectors.EVENT_READ)

    def resume_task(self, coroutine, value=None, stack=()):
        result = coroutine.send(value)
        if isinstance(result, types.GeneratorType):
            self.schedule(result, None, (coroutine, stack))
        elif isinstance(result, sleep_for_seconds):
            self.schedule(coroutine, None, stack, time()+result._wait_time)
        elif result is sys.stdin:
            self._tasks_waiting_on_stdin.append((coroutine, stack))
        elif stack:
            self.schedule(stack[0], result, stack[1])

    def schedule(self, coroutine, value=None, stack=(), when=None):
        task = partial(self.resume_task, coroutine, value, stack)
        if when:
            insort(self._timers, (when, task))
        else:
            self._tasks.append(task)

    def stop(self):
        self._running = False

    def do_on_next_tick(self, func, *args, **kwargs):
        self._tasks.appendleft(partial(func, *args, **kwargs))

    def run_forever(self):
        self._running = True
        while self._running:
            # First check for available IO input
            for key, mask in self._selector.select(0):
                line = key.fileobj.readline().strip()
                for task, stack in self._tasks_waiting_on_stdin:
                    self.schedule(task, line, stack)
                self._tasks_waiting_on_stdin.clear()

            # Next, run the next task
            if self._tasks:
                task = self._tasks.popleft()
                task()

            # Finally run time scheduled tasks
            while self._timers and self._timers[0][0] < time():
                task = self._timers[0][1]
                del self._timers[0]
                task()

        self._running = False

def print_every(message, interval):
    """
    Coroutine task to repeatedly print the message at the given interval
    (in seconds)
    """
    while True:
        print('{} - {}'.format(int(time()), message))
        yield sleep_for_seconds(interval)

def read_input(loop):
    """
    Coroutine task to repeatedly read new lines of input from stdin,
    treat the input as a number n, and calculate and display fib(n).
    """
    while True:
        line = yield sys.stdin
        if line == 'exit':
            loop.do_on_next_tick(loop.stop)
            continue
        n = int(line)
        print('fib({}) = {}'.format(n, timed_fib(n)))

def main():
    loop = EventLoop()
    hello_task = print_every('Hello world!', 3)
    fib_task = read_input(loop)
    loop.schedule(hello_task)
    loop.schedule(fib_task)
    loop.run_forever()

if __name__ == '__main__':
    main()
