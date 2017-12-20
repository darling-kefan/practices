import re
import reprlib

RE_WORD = re.compile('\w+')

class Sentence:
    def __init__(self, text):
        self.text = text

    def __repr__(self):
        return 'Sentence(%s)' % reprlib.repr(self.text)

    def __iter__(self):
        #for match in RE_WORD.finditer(self.text):
        #    yield match.group()

        # 生成器表达式和上面生成器函数等价。生成器表达式是语法糖：完全可以替换成生成器函数
        return (match.group() for match in RE_WORD.finditer(self.text))

# 生成器函数会创建一个生成器对象，包装生成器函数的定义体。把生成器传给 next(...) 函数时，生成器函数会向前，执行函数体中的下一个yield语句，返回产出的值
