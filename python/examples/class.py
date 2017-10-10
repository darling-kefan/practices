#! /usr/bin/env python3.5
# -*- coding: utf-8 -*-

# 生成器
# Generator(生成器)是一个用于创建迭代器简单而强大的工具。它们和普通的函数很像，但是当他们需要返回值时，则使用yield。每次next()被调用时，生成器会从它上次离开的地方继续执行。

print(sum(i*i for i in range(10)))
xvec = [10, 20, 30]
yvec = [7, 5, 3]
print(sum(x*y for x,y in zip(xvec, yvec)))

from math import pi, sin
sine_table = {x: sin(x*pi/180) for x in range(0, 91)}
print(sine_table)

data = 'golf'
print(list(data[i] for i in range(len(data)-1, -1, -1)))

exit()

def reverse(data):
    for index in range(len(data) - 1, -1, -1):
        yield data[index]

for char in reverse('golf'):
    print(char)

for index in range(len('golf')-1, -1, -1):
    print(index)

exit()

# 迭代器
for element in [1, 2, 3]: #列表
    print(element)
for element in (1, 2, 3): #元组
    print(element)
for key in {'one': 1, 'two': 2}:
    print(key)
for char in "123":
    print(char)
#for line in open("example.py"):
#    print(line)

# 迭代器的使用遍布于Python之中。在这个外表之下，for语句对容器对象调用了iter(). 这个函数返回一个迭代器对象，它定义了__next__()方法，用以在每次访问时的到一个元素。当没有任何元素时，__next__()将产生StopIteration异常，它告诉for停止循环。

s = 'abc'
it = iter(s)
print(it)
print(next(it))
print(next(it))
print(next(it))
# print(next(it)) # 异常会终止脚本的执行，即发生异常时，其后代码将不会被执行

class Reverse:
    "Iterator for looping over a sequence backwards"
    def __init__(self, data):
        self.data = data
        self.index = len(data)
    def __iter__(self):
        return self
    def __next__(self):
        if self.index == 0:
            raise StopIteration # raise语句会终止脚本执行
        self.index = self.index - 1
        return self.data[self.index]

rev = Reverse('spam')
print(iter(rev))
for char in rev:
    print(char)


exit()

class B(Exception):
    pass
class C(B):
    pass
class D(C):
    pass

for c in [B, C, D]:
    try:
        raise c()
    except D:
        print("D")
    except C:
        print("C")
    except B:
        print("B")

# try抛出异常后，except语句将与该异常的类或基类匹配，因此如下代码会输出B B B
for c in [B, C, D]:
    try:
        raise c()
    except B:
        print("B")
    except C:
        print("C")
    except D:
        print("D")


exit()

def scope_test():
    def do_local():
        spam = "local spam"
    def do_nonlocal():
        nonlocal spam
        spam = "nonlocal spam"
    def do_global():
        global spam
        spam = "global spam"

    spam = "test spam"
    do_local()
    print("After local assignment:", spam)
    do_nonlocal()
    print("After nonlocal assignment:", spam)
    do_global()
    print("After global assignment:", spam)

scope_test()
print("In global scope:", spam)
