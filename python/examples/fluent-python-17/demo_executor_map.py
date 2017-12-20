from time import sleep, strftime
from concurrent import futures

def display(*args): # ➊
    print(strftime('[%H:%M:%S]'), end=' ')
    print(*args)

def loiter(n): # ➋
    msg = '{}loiter({}): doing nothing for {}s...'
    display(msg.format('\t'*n, n, n))
    sleep(n)
    msg = '{}loiter({}): done.'
    display(msg.format('\t'*n, n))
    return n * 10 # ➌

def main():
    display('Script starting.')
    executor = futures.ThreadPoolExecutor(max_workers=3) # ➍
    results = executor.map(loiter, range(6)) # ➎
    display('results:', results) # ➏
    display('Waiting for individual results:')
    # for 循环中的enumerate函数会隐式调用next(results)，这个函数又会在（内部）表示第一个
    # 任务（loiter(0)）的_f期物上调用_f.result()方法。result方法会堵塞，直到期物运行结束
    # ，因此这个循环每次迭代时都要等待下一个结果做好准备。
    for i, result in enumerate(results): # ➐
        display('result {}: {}'.format(i, result))

if __name__ == '__main__':
    main()



'''
import collections
from concurrent import futures

import requests
import tqdm  ➊

from flags2_common import main, HTTPStatus  ➋
from flags2_sequential import download_one  ➌

DEFAULT_CONCUR_REQ = 30  ➍
MAX_CONCUR_REQ = 1000  ➎


def download_many(cc_list, base_url, verbose, concur_req):
    counter = collections.Counter()
    with futures.ThreadPoolExecutor(max_workers=concur_req) as executor:  ➏
        to_do_map = {}  ➐
        for cc in sorted(cc_list):  ➑
            future = executor.submit(download_one,
                            cc, base_url, verbose)  ➒
            to_do_map[future] = cc  ➓
        done_iter = futures.as_completed(to_do_map)  ⓫
        if not verbose:
            done_iter = tqdm.tqdm(done_iter, total=len(cc_list))  ⓬
        for future in done_iter:  ⓭
            try:
                res = future.result()  ⓮
            except requests.exceptions.HTTPError as exc:  ⓯
                error_msg = 'HTTP {res.status_code} - {res.reason}'
                error_msg = error_msg.format(res=exc.response)
            except requests.exceptions.ConnectionError as exc:
                error_msg = 'Connection error'
            else:
                error_msg = ''
                status = res.status

            if error_msg:
                status = HTTPStatus.error
            counter[status] += 1
            if verbose and error_msg:
                cc = to_do_map[future]  ⓰
                print('*** Error for {}: {}'.format(cc, error_msg))
    return counter


if __name__ == '__main__':
    main(download_many, DEFAULT_CONCUR_REQ, MAX_CONCUR_REQ)
'''



'''
使用 Executor.submit(...) 方法创建期物，以及如何使用 concurrent.futures.as_completed(...) 函数迭代运行结束的期物。

为什么尽管有 GIL，Python 线程仍然适合 I/O 密集型应用：标准库中每个使用 C 语言编写的 I/O 函数都会释放 GIL，因此，当某个线程在等待 I/O 时， Python 调度程序会切换到另一个线程。

多线程和多进程并发的低层实现（但却更灵活）——threading 和 multiprocessing 模块。这两个模块代表在 Python 中使用线程和进程的传统方式。
'''
