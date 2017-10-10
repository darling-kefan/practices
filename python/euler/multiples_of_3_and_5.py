#! /usr/bin/env python3.5
# -*- coding: utf-8 -*-

from functools import reduce

#sum = 0
#for k in range(1, 1000):
#    if k % 3 == 0 or k % 5 == 0:
#        sum += k
#
#print(sum)

r = reduce(lambda x, y: x + y, [x for x in range(1, 1000) if x % 3 == 0 or x % 5 == 0])
print(r)
