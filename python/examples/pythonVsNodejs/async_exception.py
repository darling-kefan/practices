def coroutine():
    '''
    这个特性使得用异常处理问题有一个统一的处理方式，不管是在同步还是异步的代码中，因为事件
    循环可以合理地捕获以及传递异常。
    '''
    print('Starting')
    try:
        yield "Let's pause until continued."
        print('Continuing')
    except Exception as e:
        yield "Got an exception: " + str(e)

def main():
    c = coroutine()
    # 输出第一个yield处后面的消息，并在yield处暂停，控制权转交给调用方
    msg = next(c)
    print(msg)
    value = c.throw(Exception('Have an exceptional day!'))
    print(value)

if __name__ == '__main__':
    main()
