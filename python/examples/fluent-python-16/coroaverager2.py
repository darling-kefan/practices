from collections import namedtuple
from inspect import getgeneratorstate
import sys

Result = namedtuple('Result', 'count average')

def averager():
    total = 0.0
    count = 0
    average = 0.0
    while True:
        term = yield
        if term is None:
            break
        total += term
        count += 1
        average = total / count
    return Result(count, average)

if __name__ == '__main__':
    aver = averager()
    print(getgeneratorstate(aver))

    # 预激协程
    next(aver)

    aver.send(10)
    aver.send(15)
    aver.send(20)
    aver.send(25)
    aver.send(30)

    print(getgeneratorstate(aver))

    try:
        aver.send(None)
    except StopIteration as e:
        # 捕获异常中的返回值e.value
        print(sys.exc_info())
        print(type(e.value), e.value, e.args, e.with_traceback)
        print(e.value.count, e.value.average)
        print(e, type(e))

    print(getgeneratorstate(aver))
