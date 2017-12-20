def averager():
    total = 0.0
    count = 0
    average = None
    while True:
        item = yield average
        total += item
        count += 1
        average = total / count

# 预激协程
avgr = averager()
v = next(avgr)
print(type(v), v)
av1 = avgr.send(10)
print(av1)
av2 = avgr.send(12)
print(av2)


from functools import wraps

def corotine(func):
    
