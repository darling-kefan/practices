from threading import Thread
from time import sleep
from time import time
from fib import timed_fib

'''
事件循环和线程

对于线程和事件循环我们需要有一个简单的认识，来理解上面两种解答的区别。先从线程说起，可以把线程理解成指令的序列以及CPU执行的上下文(CPU上下文就是寄存器的值，也就是下一条指令的寄存器)。

一个同步的程序总是在一个线程中运行的，这也是为什么在等待，比如说等待IO或定时器的时候，整个程序会被阻塞。最简单的挂起操作是sleep，它会把当前运行的线程挂起一段给定的时间。一个进程可以有多个线程，同一个进程中的线程共享了进程的一些资源，比如说内存、地址空间、文件描述符等。

线程是由操作系统的调度器来调度的，调度器统一负责管理调度进程中的线程，它来决定什么时候该把当前线程挂起，并把CPU的控制权交给另一个线程来处理。这称为上下文切换，包括对于当前线程上下文的保存、对目标线程上下文的加载。上下文切换会对性能造成一定的影响，因为它本身也需要CPU周期来执行。

操作系统切换线程有很多原因：
1. 另一个优先级更高的线程需要马上被执行(比如处理硬件中断的代码)
2. 线程自己想要被挂起一段时间(比如sleep)
3. 线程已经用完了自己的时间片，这个时候线程就不得不再次进入队列，供调度器调度

回到我们之前的代码，Python的解答是多线程的。这也解释了两个任务可以并行的原因，也就是在计算fibonaci这样的CPU密集型任务的时候，没有把其它的线程阻塞住。

再来看Node.js的解答，从计算fibonaci把定时线程阻塞住可以看出它是单线程的，这也就是Node.js实现的方式。从操作系统的角度，你的Node.js程序是在单线程上运行的(事实上，根据操作系统的不同，libuv库在处理一些IO事件的时候可能会使用线程池的方式，但这并不影响你的Javascript代码是跑在单线程上的事实)。

基于一些原因，你可能会考虑避免多线程的方式：
1. 线程在计算和资源消耗的角度是较为昂贵的。
2. 线程并发所带来的问题，比如因为共享的内存空间而带来的死锁和竞态条件。这些又会导致更加复杂的代码，在编写代码的时候需要时不时地注意一些线程安全的问题。

'''

def print_hello():
    while True:
        print('{} - Hello world!'.format(int(time())))
        sleep(3)

def read_and_process_input():
    while True:
        n = int(input())
        print('fib({}) = {}'.format(n, timed_fib(n)))

def main():
    t = Thread(target=print_hello)
    t.daemon = True
    t.start()
    # Main thread will read and process input
    read_and_process_input()

if __name__ == '__main__':
    main()
