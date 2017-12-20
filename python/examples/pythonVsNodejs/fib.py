import time

def clock(func):
    def foo(*args):
        t0 = time.perf_counter()
        result = func(*args)
        elapsed = time.perf_counter() - t0
        name = func.__name__
        arg_str = ', '.join(repr(arg) for arg in args)
        print('[%0.8fs] %s(%s) -> %r' % (elapsed, name, arg_str, result))
        return result
    return foo


def log_execution_time(func):
    def exection(*args):
        t0 = time.perf_counter()
        result = func(*args)
        elapsed = time.perf_counter() - t0
        name = func.__name__
        print('Executing {} took {:.6f} seconds'.format(name, elapsed))
        return result
    return exection

def fib(n):
    return fib(n-1) + fib(n-2) if n > 1 else n

timed_fib = log_execution_time(fib)

if __name__ == '__main__':
    print(timed_fib(30))
