#! /usr/bin/env python3.5
# -*- coding:utf-8 -*-

def ispalindromic(n):
    l = list(str(n))
    if l == [l[x] for x in range(len(l)-1, -1, -1)]:
        return True
    return False

l = [m*n for m in range(999, 99, -1) for n in range(999, 99, -1) if ispalindromic(m*n)]
print(max(l), len(l))

#maxNum = 0
#for m in range(999, 99, -1):
#    for n in range(999, 99, -1):
#        multi = m * n
#        if ispalindromic(multi) and multi > maxNum:
#            maxNum = multi
#            print(m, n)
#
#print(maxNum)
