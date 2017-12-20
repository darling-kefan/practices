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
        '''
        1. 生成器函数Sentence的定义体执行完毕后，生成器对象Sentence()会抛出StopIteration异常。
        2. 生成器函数定义体中的return语句会触发生成器对象抛出StopIteration异常。
        3. 调用方可以从StopIteration异常中获取迭代器函数中的return值。
        '''
        for w in self.words:
            yield w
        return 'stop'


if __name__ == '__main__':
    s = Sentence('I Love You, Kefan.')
    print(list(s), list(iter(s)))

    for w in s:
        print(w)

    it = iter(s)
    while True:
        try:
            val = next(it)
            print(type(val), val)
        except Exception as e:
            print(type(e), e)
            del it
            break
