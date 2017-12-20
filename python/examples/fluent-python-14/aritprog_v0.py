class ArithmeticProgression:

    def __init__(self, begin, step, end=None):
        self.begin = begin
        self.step = step
        self.end = end

    def __repr__(self):
        return 'begin:{}, end:{}, step:{}'.format(self.begin,self.end,self.step)

    #def __iter__(self):
    #    result = type(self.begin+self.step)(self.begin)
    #    forever = self.end is None
    #    while forever or result < self.end:
    #        yield result
    #        result = result + self.step

    def __iter__(self):
        result = type(self.begin+self.step)(self.begin)
        forever = self.end is None
        index = 0
        while forever or result < self.end:
            yield result
            index += 1
            result = self.begin + index*self.step

if __name__ == '__main__':
    ap = ArithmeticProgression(1, .5, 3)
    print(ap)
    print(list(ap))

