from multiprocessing import queues
import multiprocessing

def func(i, q):
    q.put(i)
    print('--->', i, q.qsize())

q = queues.Queue(9, ctx=multiprocessing)
for i in range(5):
    p = multiprocessing.Process(target=func, args=(i, q,))
    p.start()
p.join()
