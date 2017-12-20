mem = {}
def fib(n):
    if n < 2: return n
    if n not in mem:
        mem[n] = fib(n-1) + fib(n-2)
    return mem[n]

import functools

@functools.lru_cache()
def fibonacci(n):
    if n < 2: return n
    return fibonacci(n-1) + fibonacci(n-2)

if __name__ == '__main__':
    print(fib(100))
    print(fibonacci(100))

