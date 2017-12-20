'''
解释器需要迭代对象x时, 会自动调用iter(x).
内置的iter函数有以下作用，
1. 检查对象是否实现了__iter__方法， 如果实现了就调用它，获取一个迭代器。
2. 如果没有实现__iter__方法， 但是实现了__getitem__方法，python解释器会创建一个迭代器，尝试按顺序(从索引0开始)获取元素。
'''

import re
import reprlib
import collections

RE_WORD = re.compile('\w+')

'''
Python从可迭代对象中获取迭代器。
'''

class Sentence:

    def __init__(self, text):
        self.text = text
        self.words = RE_WORD.findall(text)

    def __getitem__(self, index):
        return self.words[index]

    def __len__(self):
        return len(self.words)

    def __repr__(self):
        return 'Sentence({})'.format(reprlib.repr(self.text))

if __name__ == '__main__':
    s = Sentence('"The time has come," the Walrus said,')
    print(s)
    print(len(s), list(s))
    print(isinstance(s, collections.Iterable), issubclass(Sentence, collections.Iterable))
