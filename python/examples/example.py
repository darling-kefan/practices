#! /usr/bin/env python3.5
# -*- coding: utf-8 -*-

# 遍历技巧
knights = {'gallahad': 'the pure', 'robin': 'the brave'}
for k, v in knights.items():
    print(k, v)

for i, v in enumerate(['tic', 'tac', 'toe']):
    print(i, v)

questions = ['name', 'quest', 'favorite color']
answers = ['lancelot', 'the holy grail', 'blue']
for q, a in zip(questions, answers):
    print('What is you {0}? It is {1}.'.format(q, a))

print('What is your {0}? It is {1}.'.format("name", "tangshouqiang"))

for i in reversed(range(1, 10, 2)):
    print(i)


basket = ['apple', 'orange', 'apple', 'pear', 'orange', 'banana']
for f in sorted(set(basket)):
    print(f)

exit()

# 字典map
tel = {'jack': 4098, 'sape': 4139}
print(tel)
tel['guido'] = 4127
print(tel)
del tel['sape']
print(tel)
tel['irv'] = 4127
print(tel)

print(list(tel.keys()))
print(sorted(tel.keys()))

print('jack' in tel)
print('jack' not in tel)

print(dict([('sape', 4139), ('guido', 4127), ('jack', 4098)]))
print(dict([(x, x**2) for x in (2, 4, 6)]))
print(dict(sape=4139, guido=4127, jack=4098))

exit()



# 集合set
basket = {'apple', 'orange', 'apple', 'pear', 'orange', 'banana'}
print(basket) # 重复的被移除了
print('orange' in basket)
print('crabgrass' in basket)

a = set('abracadabra')
b = set('alacazam')

print(a, b)

print(a - b) # 差集，a中有而b中没有的集合
print(a | b) # 并集
print(a & b) # 交集
print(a ^ b) # a或b中只有一个有的集合

a = {x for x in 'abracadabra' if x not in 'abc'}
print(a)

exit()

# 元组
t = 12345, 54321, 'hello!' #元组打包
print(t[0], t)
u = t, (1, 2, 3, 4, 5)
print(u)

# 元组打包
t = 12345, 54321, 'hello!'
# 序列解包
x, y, z = t
print(x, y, z)


empty = ()
singleton = "hello",
print(len(empty), len(singleton), singleton)

exit()

# del语句
# 这有一种通过给定索引而不是值，来删除列表中项的方法：用del语句。它与返回一个值的pop不同。
# del语句也可以移除列表中的切片，或者清除整个列表。

a = [-1, 1, 66.25, 333, 333, 1234.5]
print(a)
del a[0]
print(a)
del a[2:4]
print(a)
del a[:]
print(a)

# del用于删除变量实体，删除后a的引用将产生错误。
del a
print(a)

exit()


# 列表推导式
vec = [2, 4, 6]
vec1 = [3*x for x in vec]
print(vec1)

vec2 = [[x, x**2] for x in vec]
print(vec2)

freshfruit = ['  banana', '  loganberry ', 'passion fruit  ']
freshfruit2 = [weapon.strip() for weapon in freshfruit]
print(freshfruit2)

vec3 = [3*x for x in vec if x > 3]
print(vec3)

vec4 = [3*x for x in vec if x < 2]
print(vec4)

# 元组经常能不用圆括号而创建，但这里不行
vec5 = [(x, x**2) for x in vec]
print(vec5)


vec1 = [2, 4, 6]
vec2 = [4, 3, -9]
print(vec1, vec2)

vec3 = [x*y for x in vec1 for y in vec2]
print(vec3)

vec4 = [x+y for x in vec1 for y in vec2]
print(vec4)

vec5 = [vec1[i]*vec2[i] for i in range(len(vec1))]
print(vec5)

vec6 = [str(round(355/113.0, i)) for i in range(1, 6)]
print(vec6)

# 使用嵌套列表推导式时，要注意：
# 从右向左阅读嵌套推导式更容易理解。
mat = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
]
print([[row[i] for row in mat] for i in [0, 1, 2]])

for i in [0, 1, 2]:
    for row in mat:
        print(row[i], end=" ")
    print()

print(list(zip(*mat)))

exit()

# 把列表当成队列使用
# 也可以把列表当成队列使用，队列的特性是第一个添加的元素就是第一个取回的元素(即"先入先出")；
# 然而，这时列表是低效的。从列表的尾部添加和弹出是很快的，而在列表的开头插入或弹出是慢的,
# 因为所有的元素都得移动一个位置

# 要实现一个队列，使用collection.deque，它被设计成两端添加和弹出都很快。

from collections import deque
queue = deque(["Eric", "John", "Michael"])
queue.append("Terry")
queue.append("Graham")
print(queue)
queue.popleft()
queue.popleft()
print(queue)

exit()

# sdfsdfsdf列表作为堆栈使用
stack = [3, 4, 5]
stack.append(1)
stack.append(2)
print(stack)
stack.pop()
print(stack)
stack.pop()
print(stack)
stack.pop()
print(stack)


exit()

a = [66.25, 333, 333, 1, 1234.5]
print(a.count(333), a.count(66.25), a.count('x'))

a.insert(2, -1)
a.append(333)
print(a)

print(a.index(333))

a.remove(333)
print(a)

a.reverse()
print(a)

a.sort()
print(a)
