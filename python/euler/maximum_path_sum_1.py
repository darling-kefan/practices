#! /usr/bin/env python3
# -*- coding:utf-8 -*-

'''
问题：
https://projecteuler.net/problem=18
二叉树基本概念：
http://ccc013.github.io/2016/08/18/%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%9F%BA%E6%9C%AC%E6%A6%82%E5%BF%B5%E5%92%8C%E5%AE%9E%E7%8E%B0/
'''

import math

def quadratic(a, b, c):
    '''
    计算平方跟
    '''

    if a == 0:
        return -c/b
    elif b*b - 4*a*c < 0:
        return None
    elif b*b - 4*a*c == 0:
        return -b/(2*a)
    else:
        return (-b - math.sqrt(b*b - 4*a*c))/(2*a), (-b + math.sqrt(b*b - 4*a*c))/(2*a)

def getLine(n, maxLen):
    '''
    获取n在三角形上所处的行号，该行用于计算n的子树
    '''

    curLen = 0
    for l in range(1, maxLen+1):
        if n < ((l+1)/2)*l:
            curLen = l
            break
    return curLen

class treeNode():
    '''
    描述树节点
    '''

    def __init__(self, v=None, l=None, r=None):
        self.data = v
        self.left = l
        self.right = r

def create(l):
    '''
    三角形列表生成二叉树
    '''

    x, y = quadratic(1, 1, -2*len(l))
    if x > 0:
        maxLen = int(math.ceil(x))
    elif y > 0:
        maxLen = int(math.ceil(y))

    # 将三角形里所有数字生成树节点
    for k in range(len(l)):
        globals()['tn_'+str(k)] = treeNode(v=l[k])

    # 设置树节点的左右子树
    for k in range(len(l)):
        # 计算节点所处行号
        line = getLine(k, maxLen)

        if line == maxLen: # 叶子节点
            exec('tn_'+str(k)+'.left=None')
            exec('tn_'+str(k)+'.right=None')
        else: # 非叶子节点
            if eval('"tn_' + str(k+line) + '" in globals()'):
                exec('tn_'+str(k)+'.left=tn_'+str(k+line))
            else:
                exec('tn_'+str(k)+'.left=None')

            if eval('"tn_' + str(k+line+1) + '" in globals()'):
                exec('tn_'+str(k)+'.right=tn_'+str(k+line+1))
            else:
                exec('tn_'+str(k)+'.right=None')

def show(tn):
    '''
    显示二叉树
    '''

    print(tn.data)
    if tn.left != None:
        show(tn.left)
    if tn.right != None:
        show(tn.right)

# 路径集合
paths = []

# list传递是传引用
def getPaths(node, path = []):
    '''
    获取三角形从根节点到叶子节点的所有路径
    '''

    path.append(node)

    # 若是叶子节点
    if node.left == None and node.right == None:
        paths.append(path[:])
        path.pop()
        return

    if node.left != None:
        getPaths(node.left, path)
    if node.right != None:
        getPaths(node.right, path)

    path.pop()



if __name__ == '__main__':

    # 生成二叉树
    originList = [75,95,64,17,47,82,18,35,87,10,20,4,82,47,65,19,1,23,75,3,34,88,2,77,73,7,63,67,99,65,4,28,6,16,70,92,41,41,26,56,83,40,80,70,33,41,48,72,33,47,32,37,16,94,29,53,71,44,65,25,43,91,52,97,51,14,70,11,33,28,77,73,17,78,39,68,17,57,91,71,52,38,17,14,91,43,58,50,27,29,48,63,66,4,68,89,53,67,30,73,16,69,87,40,31,4,62,98,27,23,9,70,98,73,93,38,53,60,4,23]
    create(originList)

    # 获取所有路径
    getPaths(tn_0)

    maxSum  = 0
    maxPath = 0
    for k in range(len(paths)):
        sum = 0
        for kk in range(len(paths[k])):
            sum += paths[k][kk].data
        if sum > maxSum:
            maxSum = sum
            maxPath = k

    print('-------------------------------------')
    print(maxSum)
    print('-------------------------------------')

    # 打印和最大的路径
    maxPath = paths[maxPath]
    for k in range(len(maxPath)):
        print(maxPath[k].data)

