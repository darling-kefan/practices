#! /usr/bin/env python3.5
# -*- coding: utf-8 -*-

'''
用来向生成器函数送入一个异常，可以结束系统定义的异常，或者自定义的异常。
throw()后直接抛出异常并结束程序，或者消耗掉一个yield，或者在没有下一个yield的时候直接进行到程序的结尾。
'''

def gen():
    while True:
        try:
            yield 'normal value'
            yield 'normal value 2'
            print('here')
        except ValueError:
            print('we got ValueError here')
        except TypeError:
            break

g = gen()
print(next(g))
print(g.throw(ValueError))
print(next(g))
print(g.throw(TypeError))
