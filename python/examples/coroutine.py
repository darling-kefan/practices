#! /usr/bin/env python3.5
# -*- coding: utf-8 -*-

def consumer():
    r = ''
    while True:
        n = yield r
        if not n:
            return
        print('[CONSUMER] Consuming %s...' % n)
        r = '200 OK'

def produce(c):
    c.send(None) # 启动generator:consumer. 同时代码跳至consumer函数并执行，遇到yield后，consumer函数停止执行，并返回r('')
    n = 0
    while n < 5:
        n = n + 1
        print('[PRODUCER] Producing %s...' % n)
        r = c.send(n) # 跳转到generator:consumer的yield处继续执行，此时yield接收send参数n并赋给yield语句(n yield r)里的n，第一次n=1，consumer输出[CONSUMER] Consuming 1...，又到yield语句，consumer执行暂停并返回r='200 OK'赋给该处的变量r，以此循环执行...
        print('[PRODUCER] Consumer return: %s' % r)
    c.close()

c = consumer()
produce(c)
