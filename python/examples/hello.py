
print(list(range(3, 6)))
args = [3, 6]
print(list(range(*args)))

def parrot(voltage, state='a stiff', action='voom'):
    print("-- This parrot wouldn't", action, end=' ')
    print("if you put", voltage, "volts through it.", end=' ')
    print("E's", state, "!")

d = {"voltage": "four million", "state": "bleedin' demised", "action": "VOOM"}
parrot(**d)

exit()

def concat(*arg, sep="/"):
    return sep.join(arg)

print(concat("abc", "def", "ghi", "jkl", sep=","))

def concat2(sep="/", *arg):
    return sep.join(arg)

print(concat2("|", "abc", "def", "ghi", "jkl"))

exit()

def cheeseshop(kind, *arguments, **keywords):
    print("-- Do you have any", kind, "?")
    print("-- I'm sorry, we're all out of", kind)
    for arg in arguments:
        print(arg)
    print("-" * 40)
    keys = sorted(keywords.keys())
    for kw in keys:
        print(kw, ":", keywords[kw])

#cheeseshop("Limburger", "It's very runny, sir.",
#           "It's really very, VERY runny, sir.",
#           shopkeeper="Michael Palin",
#           client="John Cleese",
#           sketch="Cheese Shop Sketch")

cheeseshop("Limburger", "It's very runny, sir.",
           "It's really very, VERY runny, sir.",
           shopkeeper="Michael Palin",
           client="John Cleese",
           sketch="Cheese Shop Sketch")

exit()

def parrot(voltage, state='a stiff', action='voom', type='Norwegian Blue'):
    print("-- This parrot wouldn't", action, end=' ')
    print("if you put", voltage, "volts through it.")
    print("-- Lovely plumage, the", type)
    print("-- It's", state, "!")

# parrot(1000)
# parrot(action = 'VOOOOOM', voltage = 1000000)
# parrot('a thousand', state = 'pushing up the daisies')
parrot(action = 'VOOOOOM')

exit()


def f(a, L=None):
    if L == None:
        L = []
    L.append(a)
    return L

print(f(1))
print(f(2))
print(f(3))

exit()

# 默认参数是可变参数的情况下，如列表、字典、或大多数对象时，函式在随后的调用中会累计参数值：
def f(a, L=[]):
    L.append(a)
    print(L)

f(1)
f(2)
f(3)

exit()

# 默认参数的值等于函式定义域中的值。默认参数的值只会被计算一次。
i = 5
def f(arg=i):
    print(arg)

i = 6
f()

exit()

def fib(n):
    """返回一个列表，包含直到 n 的 Fibonacci 序列."""
    result = []
    a, b = 0, 1
    while a < n:
        result.append(a)
        a, b = b, a + b
    return result

f100 = fib(100)
print(f100)

exit()

while True:
    pass # 忙等待键盘中断(Ctrl + C)

exit()

for num in range(2, 10):
    if num % 2 == 0:
        print("Found an even number", num)
        continue
    print("Found a number", num)

exit()

# else 在列表中无元素可迭代时(for) 或 循环条件为false时(while)执行；注意，
# 执行break语句后，else语句不会被执行


# 循环语句可以有一个else子句；当循环因耗尽整个列表而终止时(for)或者当条件变为假
# 时(使用while)，它会被执行。但是，如果循环因break语句终止的话，它不会被执行。
for n in range(2, 10):
    for x in range(2, n):
        if n % x == 0:
            print(n, "equals", x, "*", n//x)
            break
    else:
        print(n, " is a prime number")

exit()

# 在很多时候，由range()返回的对象表现得就像一个列表，但实际上它不是。如果你对其
# 进行迭代时，它能返回所需要的连续项，但实际上为了节省空间并没有真正真正生成制造
# 一个列表。
# 我们称这种对象叫做iterable，也就是说，某些函式和构造器期望能从对象连续接收元
# 素直至终结，我们称这种对象叫做iterable(可迭代的)。


print(range(10))
print(list(range(10)))

print(range(10))
print(list(range(10)))

exit()

a = ['Mary', 'had', 'a', 'little', 'lamb']
for i in range(len(a)):
    print(i, a[i])

exit()



for i in range(0, 10, 2):
    print(i)

exit()


for i in range(5):
    print(i)

exit()


a = ['cat', 'window', 'defenestrate']
for x in a[:]:
    print(x, len(x))
    if len(x) > 6: a.insert(0, 'linux')

print(a)




exit()

def sdf():
    print("Hello world")


a = ["I", "am", "your", "father"]
print(a)

sdf()

# 第一行包括一次*多重赋值*：变量a和b同时地的到新值0和1. 在最后一行又使用了一次，演示了右边的表达式在任何赋值之前就已经被计算了。右边的表达式从左至右地计算。
# 当条件保持真时，while循环一直执行。在Python里，就像C里一样，任何非整数都为真；零为假，

x = int(input("Please enter an integer: "))

if x < 0:
    x = 0
    print("Negative changed to zero")
elif x == 0:
    print("Zero")
elif x == 1:
    print("Single")
else:
    print("More")


