#! /usr/bin/env python3.5
# -*- coding:utf-8 -*-

import math

def isprime(n):
    for x in range(2, int(math.sqrt(n))+1):
        if n % x == 0:
            return False
    return True

x = int(input('Please input an integer: '))
l = [i for i in range(2, int(math.sqrt(x))+1) if x % i == 0 and isprime(i)]
print(l, max(l))

