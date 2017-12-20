import copy
import random

class BingoCage:

    def __init__(self, items):
        self._items = list(items)
        random.shuffle(self._items)

    def pick(self):
        return self._items.pop()

    def leftover(self):
        return self._items

    def __call__(self):
        '''使类实例可以像函数那样被调用'''
        return self.pick()

if __name__ == '__main__':
    bingo = BingoCage(range(0,5))
    raw = copy.deepcopy(bingo.leftover())
    item = bingo()
    leftover = list(bingo.leftover())
    print(callable(bingo), raw, item, leftover)
    print(callable(bingo), bingo.leftover()[:], bingo(), bingo.leftover())
