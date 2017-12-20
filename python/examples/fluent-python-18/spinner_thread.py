import threading
import itertools
import time
import sys

# 定义一个简单的可变对象；其中有个go属性，用于在外部控制线程。
class Signal: # 1
    go  = True

def spin(msg, signal): # 2
    write, flush = sys.stdout.write, sys.stdout.flush
    # 这其实是个无限循环，因为itertools.cycle函数会从指定的序列中反复不断地生成元素
    for char in itertools.cycle('|/-\\'): # 3
        status = char + ' ' + msg
        write(status)
        flush()
        # 这是显示文本式动画的诀窍所在：使用退格符(\x08)把光标移回来
        write('\x08' * len(status)) # 4
        time.sleep(.1)
        if not signal.go: # 5
            break

    write(' ' * len(status) + '\x08' * len(status)) # 6

def slow_function(): # 7
    # 调用sleep函数会堵塞主线程，不过一定要这么做，以便释放GIL，创建从属线程
    time.sleep(3) # 8 
    return 42

# 这个函数设置从属线程，显示线程对象，运行耗时的计算，最后杀死线程
def supervisor(): # 9
    signal = Signal()
    spinner = threading.Thread(target=spin,
                               args=('thinking!', signal))
    print('spinner object:', spinner) # 10
    # 启动从属线程
    spinner.start() # 11
    # 堵塞主线程。同时，从属线程以动画形式显示旋转指针
    result = slow_function() # 12
    signal.go = False # 13
    # 等待spinner线程结束
    spinner.join() # 14
    return result

def main():
    result = supervisor() # 15
    print('Answer:', result)

if __name__ == '__main__':
    main()

# 注意，Python没有提供终止线程的API，这是有意为之的。若想关闭线程，必须给线程发送消息。这里，我使用的是signal.go属性: 在主线程中把它设为False后，spinner线程最终会注意到，然后干净地退出。
