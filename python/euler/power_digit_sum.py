#! /usr/bin/env python3
# -*- coding: utf-8 -*-

n = 2**1000
print(len(str(n)))

m = 0
for v in str(n):
    m += int(v)

print(m)
