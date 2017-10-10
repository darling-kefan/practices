#! /usr/bin/env python3.5
# -*- coding: utf-8 -*-

import sys

from datetime import date

now = date.today()
print(now)

nowStr = now.strftime("%m-%d-%y. %d %b %Y is a %A on the %d day of %B.")
print(nowStr)

sys.exit()

import math

print(math.cos(math.pi / 4))
print(math.log(1024, 2))

import random

print(random.choice(['apple', 'pear', 'banana']))
print(random.sample(range(100), 10))
print(random.random())
print(random.randrange(6))

sys.exit()

import re

re1 = re.findall(r'\bf[a-z]*', 'which foot or hand fell fastest')
re2 = re.sub(r'(\b[a-z]+) \1', r'\1', 'cat in the the hat')

print(re1, re2)

print('tea for too'.replace('too', 'two'))

sys.exit()

import glob
print(glob.glob("*.py"))

import sys
print(sys.argv)

sys.stderr.write("Warning, log file not found starting a new one\n")

sys.exit() #结束整个程序

print("hello world")

