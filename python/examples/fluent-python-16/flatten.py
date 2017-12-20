from collections import Iterable

def flatten(items, ignore_types=(str, bytes)):
    for item in items:
        if isinstance(item, Iterable) and not isinstance(item, ignore_types):
            yield from flatten(item)
        else:
            yield item



items = [1, 2, [3, 4, [5, 6], 7], 8]

for i in flatten(items):
    print(i)

items = ['Dave', 'Paula', ['Thomas', 'Lewis']]
for x in flatten(items):
    print(x)

print('---------------------------------------------------------------')

def flatten2(items, ignore_types=(str, bytes)):
    for item in items:
        if isinstance(item, Iterable) and not isinstance(item, ignore_types):
            flatten2(item, ignore_types)
        else:
            print(item)

flatten2(items)

