import contextlib

@contextlib.contextmanager
def looking_glass():
    '''
    前面说过，为了告诉解释器异常已经处理了，__exit__方法会返回True，此时解释器会压制异常。如果__exit__方法没有显式返回一个值，那么解释器的到的是None，然后向上冒泡异常。
    使用@contextmanager装饰器时，默认的行为是相反的：装饰器提供的__exit__方法假定发给生成器的所有异常都得到处理了，因此应该压制异常。如果不想让@contextmanager压制异常，必须在被装饰的函数中显式重新抛出异常。

    如果在with块中抛出了异常，Python解释器会将其捕获，然后在looking_glass函数的yield表达式里再次抛出。
    '''
    import sys
    original_write = sys.stdout.write

    def reverse_write(text):
        original_write(text[::-1])

    sys.stdout.write = reverse_write

    msg = ''
    try:
        yield 'JABBERWOCKY'
    except ZeroDivisionError:
        msg = 'Please DO NOT divide by zero!'
    finally:
        sys.stdout.write = original_write
        if msg:
            print(msg)
