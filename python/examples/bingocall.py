#! /usr/bin/env python3
# -*- coding:utf-8 -*-

import random

class BingoCage:
    def __init__(self, items):
        self._items = list(items)
        random.shuffle(self._items)

    def items(self):
        return self._items

    def pick(self):
        'return item'
        try:
            return self._items.pop()
        except IndexError:
            raise LookupError('pick from empty BingoCage')

    def __call__(self):
        return self.pick()

bingo = BingoCage(range(5))
print(bingo.pick(), bingo.items(), bingo(), bingo.items(), callable(bingo))


