import itertools

class ArithmetricProgression:

    def __init__(self, begin, step, end=None):
        self.begin = begin
        self.step = step
        self.end = end

    def __iter__(self):
        result = type(self.begin + self.step)(self.begin)
        forever = self.end is None
        index = 0
        while forever or result < self.end:
            yield result
            index += 1
            result = self.begin + index * self.step

l1 = ArithmetricProgression(0, 1, 5)
print(list(l1))

l2 = ArithmetricProgression(0, 2, 10)
print(list(l2))

def arithmetricGen(begin, step, end=None):
    result = type(begin + step)(begin)
    index = 0
    while end is None or result < end:
        yield result
        index += 1
        result = begin + index * step

l3 = arithmetricGen(0, 1, 5)
print(list(l3))

l4 = arithmetricGen(0, 2, 10)
print(list(l4))


def aritprog_gen(begin, step, end=None):
    first = type(begin + step)(begin)
    ap_gen = itertools.count(begin, step)
    if end is not None:
        ap_gen = itertools.takewhile(lambda n: n < end, ap_gen)
    return ap_gen

l5 = aritprog_gen(0, 1, 5)
print(list(l5))

l6 = aritprog_gen(0, 2, 10)
print(list(l6))


