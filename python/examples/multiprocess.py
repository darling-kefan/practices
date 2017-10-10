#! /usr/bin/env python3.5
# -*- coding: utf-8 -*-

#from multiprocessing import Process
from multiprocessing import Process
import os

# 子进程要执行的代码
def run_proc(name):
    print('Run child process %s (%s)...' % (name, os.getpid()))

if __name__ == '__main__':
    print('Parent process %s.' % os.getpid())
    p = Process(target=run_proc, args=('test',))
    print('Child process will start.')
    # 启动进程
    p.start()
    # 等待子进程结束后再继续往下运行，通常用于进程间的同步
    p.join()
    print('Child process end.')
