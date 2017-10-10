
# 杨辉三角
def triangles(max):
    n = 1
    l = []
    while n <= max:
        if n == 1:
            l = [1]
        elif n == 2:
            l = [1, 1]
        else:
            l = [l[i] + l[i+1] for i in range(len(l)-1)]
            l.insert(0, 1)
            l.insert(len(l), 1)

        yield l
        n += 1

t = triangles(10)
for k in t:
    print(k)


exit()

def fib(max):
    n, a, b = 0, 0, 1
    while n < max:
        print(b)
        a, b = b, a + b
        n += 1
    return 'done'

fib(10)


def fib(max):
    n, a, b = 0, 0, 1
    while n < max:
        yield b
        a, b = b, a + b
        n += 1
    return 'done'

f = fib(10)
for n in f:
    print(n)

while True:
    try:
        x = next(f)
        print('g:', x)
    except StopIteration as e:
        print('Generator return value:', e.value)
        break

exit()

## 关键字参数
def person(name, age, **kw):
    print('name:', name, 'age:', age, 'other:', kw)

person('Michael', 30)
person('Bob', 35, city='Beijing')
person('Adam', 45, gender='M', job='Engineer')
args = {'gender': 'F', 'job': 'worker', 'city': 'shanghai'}
person('shouqiang', 31, **args)
print(args)


exit()

## 可变参数

args1 = [1, 2, 3, 4, 5]
args2 = (1, 2, 3, 4)

# list/tuple作为参数
def calc(numbers):
    sum = 0
    for num in numbers:
        sum += num * num
    return sum

r1 = calc(args1)
r2 = calc(args2)
print(r1, r2)

# 可变参数
def calcNew(*numbers):
    sum = 0
    for num in numbers:
        sum += num*num
    return sum

r1 = calcNew()
r2 = calcNew(1)
r3 = calcNew(1, 2, 3)
r4 = calcNew(*args1)
r5 = calcNew(*args2)
print(r1, r2, r3, r4, r5)

# 总结：定义可变参数和定义一个list/tuple作为参数相比，仅仅在参数前面加一个*号；传递list/tuple作为可变参数的话，需在其前加上*号；

exit()

import math

# ax2 + bx + c = 0
def quadratic(a, b, c):
    if b*b - 4*a*c >= 0:
        x1 = (-b + math.sqrt(-4*a*c + b*b)) / (2*a)
        x2 = (-b - math.sqrt(-4*a*c + b*b)) / (2*a)
        return x1, x2
    else:
        raise TypeError('bad operand type')

print(quadratic(2, 3, 1))
print(quadratic(1, -3, 4))
