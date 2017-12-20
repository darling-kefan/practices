'''
迭代器对象是迭代器对象，迭代器是迭代器，是不同的东西。迭代器实现collections.abc.Iterator, 包含__iter__和__next__方法；迭代器对象实现collections.abc.Iterable，只包含__iter__方法。

'''

'''
方式一：
>>> import collections
>>> issubclass(Sentence, collections.abc.Iterator)

>>> from collections import abc
>>> issubclass(Sentence, abc.Iterator)

第一种方式报错：
AttributeError: module 'collections' has no attribute 'abc'
第二种方式正确。

'''

import re
import reprlib
from collections import abc

RE_WORD = re.compile('\w+')

class Sentence:

    def __init__(self, text):
        self.text = text
        self.words = RE_WORD.findall(text)

    def __repr__(self):
        return 'Sentence({})'.format(reprlib.repr(self.text))

    def __iter__(self):
        return SentenceIteror(self.words)

class SentenceIteror:

    def __init__(self, words):
        self.words = words
        self.index = 0

    def __next__(self):
        try:
            word = self.words[self.index]
        except IndexError:
            raise StopIteration()
        self.index += 1
        return word

    def __iter__(self):
        return self

if __name__ == '__main__':

    print(issubclass(SentenceIteror, abc.Iterator))
    print(issubclass(SentenceIteror, abc.Iterable))

    print(issubclass(Sentence, abc.Iterator))
    print(issubclass(Sentence, abc.Iterable))

    s = Sentence('I love you, kefan.')
    print(list(s), list(iter(s)))

    for w in s:
        print(w)

    it = iter(s)
    while True:
        try:
            print(next(it))
        except StopIteration:
            del it
            break
