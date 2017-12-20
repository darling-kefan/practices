#! /usr/bin/env python3
# -*- coding:utf-8 -*-

import time
from clockdeco2 import clock

@clock
def snooze(seconds, hello='world'):
    time.sleep(seconds)

if __name__ == '__main__':
    print('*' * 40, 'Calling snooze(.123)')
    snooze(.123)
