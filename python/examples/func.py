class Student(object):
    def __init__(self, name, score):
        self.__name = name
        self.__score = score
    def print_score(self):
        print('%s: %s' %s (self.__name, self.__score))
    def getName(self):
        return self.__name
    def getScore(self):
        return self.__score

bart = Student('Bart Simpson', 98)
print(bart.getName())

bart.__name = 'New Name'
print(bart.__name)
print(bart.getName())


exit()

#def log(text):
#    def decorator(func):
#        def wrapper(*args, **kw):
#            print('%s %s():' % (text, func.__name__))
#            return func(*args, **kw)
#        return wrapper
#    return decorator
#
#@log('execute')
#def now():
#    print('2015-3-25')
#
#now()
#
#exit()

import functools

def log(func):
    @functools.wraps(func)
    def wrapper(*args, **kw):
        print('call %s():' % func.__name__)
        return func(*args, **kw)
    return wrapper

@log
def now():
    print('2015-3-25')

now()

print(now.__name__)

exit()


def count():
    def f(i):
        def g():
            return i*i
        return g
    fs = []
    for i in range(1, 4):
        fs.append(f(i))
    return fs

f1, f2, f3 = count()
print(f1(), f2(), f3())

def count():
    fs = []
    for i in range(1, 4):
        def f():
            return i*i
        fs.append(f)
    return fs

f1, f2, f3 = count()
print(f1(), f2(), f3())

exit()

def calc_sum(*args):
    ax = 0
    for n in args:
        ax = ax + n
    return ax

def lazy_sum(*args):
    def sum():
        ax = 0
        for n in args:
            ax = ax + n
        return ax
    return sum

print(calc_sum(1, 2, 3, 4, 5))
f1 = lazy_sum(1, 2, 3, 4, 5)
f2 = lazy_sum(1, 2, 3, 4, 5)
print(f1, f2)
print(f1 == f2)

exit()

def is_palindrome(n):
    ns = str(n)
    rt = ""
    for k in range(len(ns)-1, -1, -1):
        rt = rt + ns[k]
    if rt == ns:
        return True
    return False

def is_palindrome2(n):
    s = str(n)
    for i in range(len(s) // 2):
        if s[i] != s[len(s)-1-i]:
            return False
    return True

output = filter(is_palindrome2, range(1, 1000))
print(list(output))

exit()

def _odd_iter():
    n = 1
    while True:
        n = n + 2
        yield n

def _not_divisible(n):
    return lambda x: x % n > 0

def primes():
    yield 2
    it = _odd_iter() #初始化序列
    while True:
        n = next(it)
        yield n
        it = filter(_not_divisible(n), it)

for n in primes():
    if n < 1000:
        print(n)
    else:
        break

exit()

from functools import reduce

CHAR_TO_FLOAT = {
    '0': 0,
    '1': 1,
    '2': 2,
    '3': 3,
    '4': 4,
    '5': 5,
    '6': 6,
    '7': 7,
    '8': 8,
    '9': 9,
    '.': -1
}

def str2float(s):
    nums = map(lambda ch: CHAR_TO_FLOAT[ch], s)
    point = 0
    def to_float(f, n):
        nonlocal point
        if n == -1:
            point = 1
            return f
        if point == 0:
            return f * 10 + n
        else:
            point = point * 10
            return f + n / point
    return reduce(to_float, nums, 0.0)

print(str2float('0'))
print(str2float('123.456'))
print(str2float('123.45600'))
print(str2float('0.1234'))
print(str2float('.1234'))
print(str2float('120.0034'))
