import contextlib

@contextlib.contextmanager
def looking_glass():
    '''
    @contextmanager装饰器能减少创建上下文管理器的样板代码量，因为不用编写一个完整的类，定义__enter__和__exit__方法，而只需实现有一个yield语句的生成器，生成想让__enter__方法返回的值

    在使用@contextmanger装饰的生成器中，yield语句的作用是把函数的定义体分成两部分：yield语句前面的所有代码在with块开始时（即解释器调用__enter__方法时）执行，yield语句后面的代码在with块结束时（即调用__exit__方法时）执行。

    其实，contextlib.contextmanager装饰器会把函数包装成实现__enter__和__exit__方法的类。
    '''
    import sys
    original_write = sys.stdout.write

    def reverse_write(text):
        original_write(text[::-1])

    sys.stdout.write = reverse_write
    yield 'JABBERWOCKY'
    sys.stdout.write = original_write
