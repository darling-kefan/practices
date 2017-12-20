#def aritprog_gen(begin, step, end=None):
#    result = type(begin + step)(begin)
#    index = 0
#    forever = end is None
#    while forever or result < end:
#        yield result
#        index += 1
#        result = begin + index*step

import itertools

def aritprog_gen2(begin, step, end=None):
    begin = type(begin+step)(begin)
    gen = itertools.count(begin, step)
    if end is not None:
        gen = itertools.takewhile(lambda n: n < end if step > 0 else n > end, gen)
    return gen
