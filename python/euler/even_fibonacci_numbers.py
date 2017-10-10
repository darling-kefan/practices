#! /usr/bin/env python3.5
# -*- coding: utf-8 -*-

from functools import reduce

def fib(max):
    x, y = 1, 1
    while y < max:
        yield y
        x, y = y, x+y

r = reduce(lambda x, y: x + y, [x for x in list(fib(4000000)) if x % 2 == 0])
print(r)
