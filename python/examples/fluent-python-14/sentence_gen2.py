import re
import reprlib
from collections import abc

RE_WORD = re.compile('\w+')

class Sentence:

    def __init__(self, text):
        self.text = text

    def __repr__(self):
        return 'Sentence({})'.format(reprlib.repr(self.text))

    #def __iter__(self):
    #    for match in RE_WORD.finditer(self.text):
    #        yield match.group()

    def __iter__(self):
        return (match.group() for match in RE_WORD.finditer(self.text))

if __name__ == '__main__':
    '''
    生成器表达式可以理解为列表推导式的惰性实现：不会急迫地构建列表，而是返回一个生成器，按需惰性生成元素。
    
    生成器表达式和生成器函数一样都用于产出生成器对象。生成器表达式返回一个生成器对象，而调用生成器函数时候返回一个生成器对象。
    '''
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
