from functools import wraps

def averager():
    total = 0
    count = 0
    avg   = None
    while True:
        item = yield avg
        total += item
        count += 1
        avg = total / count

def corotine(func):
    '''协程装饰器，将生成器执行到第一个yield处，预激协程'''
    @wraps(func)
    def primer(*args, **kwargs):
        avger = func(*args, **kwargs)
        next(avger)
        return avger
    return primer

@corotine
def averager2():
    total = 0
    count = 0
    avg   = None
    while True:
        item = yield avg
        total += item
        count += 1
        avg = total / count
