#! /usr/bin/env python3.5
# -*- coding:utf-8 -*-

'''
子程序，或者称为函数，在所有语言中都是顶层调用，比如A调用B，B的执行过程中又调用了C，C执行完毕返回，B执行完毕返回，最后是A执行完毕。

所以子程序调用是通过栈实现的，一个线程就是执行一个子程序。

子程序调用总是一个入口，一次调用，调用顺序是明确的。而协程的调用和子程序不同。

协程看上去也是子程序，但执行过程中，在子程序内部可中断，然后转而执行别的子程序，在适当的时候返回来接着执行。

注意，在一个子程序中中断，去执行其他子程序，不是函数调用，有点类似CPU的中断。

将yield from视为提供了一个调用者和子生成器之间的透明的双向通道。包括从子生成器获取数据以及向子生成器传送数据。
'''


def accumulate():
    tally = 0
    while 1:
        next = yield
        if next is None:
            return tally
        tally += next

def gather_tallies(tallies):
    while 1:
        tally = yield from accumulate()
        tallies.append(tally)

tallies = []
acc = gather_tallies(tallies)
next(acc) # 使累加生成器准备好接收传入值,等价于acc.send(None)
for i in range(4):
    acc.send(i)
acc.send(None)

for i in range(5):
    acc.send(i)
acc.send(None)

print(tallies)


exit()

class SpamException(Exception):
    pass

def writer():
    while True:
        try:
            w = (yield)
        except SpamException:
            print('***')
        else:
            print('>> ', w)

def writer_wrapper(coro1):
    coro1.send(None)
    while True:
        try:
            #x = (yield) # yield只能够引发异常，然后停止运行
            #coro1.send(x)
            try:
                x = (yield)
            except Exception as e:
                coro1.throw(e)
            else:
                coro1.send(x)
        except StopIteration:
            pass

w = writer()
wrap = writer_wrapper(w)
wrap.send(None)
for i in [0, 1, 2, 'spam', 4]:
    if i == 'spam':
        wrap.throw(SpamException)
    else:
        wrap.send(i)

exit()

# 利用yield from语句向生成器(协程)
def writer():
    # 读取send传进的数据，并模拟写进套接字，文件等
    while True:
        w = (yield) # w接收send传进的数据
        print('>> ', w)

#def writer_wrapper(coro):
#    pass

#def writer_wrapper(coro1):
#    coro1.send(None) # 生成准备好接收数据
#    while True:
#        try:
#            x = (yield) # x接收send传进的数据
#            coro1.send(x) # 然后将x在send给writer子生成器
#        except StopIteration: # 处理子生成器返回的异常
#            pass

def writer_wrapper(coro2):
    yield from coro2

w = writer()
wrap = writer_wrapper(w)
wrap.send(None) # 生成器准备好接收数据
for i in range(4):
    wrap.send(i)



exit()


# 利用yield from从生成器读取数据
def reader():
    for i in range(4):
        yield '<< %s' % i

def reader_wrapper(g):
    #for v in g:
    #    yield v
    yield from g

wrap = reader_wrapper(reader())
for i in wrap:
    print(i)

