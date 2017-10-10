# Fibonacci 数列模块

def fib(n): # 打印小于 n 的 Fibonacci 数列
    a, b = 0, 1
    while b < n:
        print(b, end=' ')
        a, b = b, a+b
    print()

def fib2(n):
    result = []
    a, b = 0, 1
    while b < n:
        result.append(b)
        a, b = b, a+b
    return result

# 把模块当作脚本执行
if __name__ == "__main__":
    import sys
    fib(int(sys.argv[1]))
