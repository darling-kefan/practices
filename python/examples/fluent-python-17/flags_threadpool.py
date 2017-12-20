from concurrent import futures

from flags import save_flag, get_flag, show, main

MAX_WORKERS = 20

def download_one(cc):
    image = get_flag(cc)
    show(cc)
    save_flag(image, cc.lower() + '.gif')
    return cc

def download_many(cc_list):
    workers = min(MAX_WORKERS, len(cc_list))
    # 使用工作的线程数实例化 ThreadPoolExecutor 类；executor.__exit__方法会调用executor.shutdown(wait=True)方法，它会在所有线程都执行完毕前阻塞线程。
    with futures.ThreadPoolExecutor(workers) as executor:
        res = executor.map(download_one, sorted(cc_list))

    return len(list(res))

def download_many2(cc_list):
    cc_list = cc_list[:5]
    # 把max_workers硬编码为3,以便在输出中观察待完成的期物。
    with futures.ThreadPoolExecutor(max_workers=5) as executor:
        to_do = []
        # 按照字母表顺序迭代国家代码，明确表明输出的顺序与输入一致
        for cc in sorted(cc_list):
            # 排定可调用对象的执行时间，然后返回一个期物，表示这个待执行的操作
            future = executor.submit(download_one, cc)
            # 存储各个期物，后面传给as_completed函数
            to_do.append(future)
            msg = 'Scheduled for {}: {}'
            print(msg.format(cc, future))

        results = []
        # as_completed函数在期物运行结束后产出期物
        for future in futures.as_completed(to_do):
            res = future.result()
            msg = '{} result: {!r}'
            print(msg.format(future, res))
            results.append(res)

    return len(results)

# 如果使用Python处理CPU密集型工作，应该试试PyPy解释器。性能明显提醒3～5倍。
def download_many3(cc_list):
    # I/O密集型作业使用ProcessPoolExecutor类得不到任何好处。
    with futures.ProcessPoolExecutor() as executor:
        res = executor.map(download_one, sorted(cc_list))

    return len(list(res))

if __name__ == '__main__':
    main(download_many)

# futures(期贷|期物)表示封装待完成的操作，可以放入队列，完成的状态可以查询，得到的结果(或抛出异常)后可以获取结果(或异常)


# 阻塞型I/O和GIL

# CPython解释器本身就不是线程安全的，因此有全局解释器锁(GIL)，一次只允许使用一个线程执行Python字节码。因此，一个Python进程通常不能同时使用多个CPU核心。

# 这是CPython解释器的局限性，与Python语言本身无关。Jython和IronPython没有这种限制。不过目前最快的Python解释器PyPy也有GIL。

# Python标准库中的所有阻塞型I/O函数都会释放GIL，允许其它线程运行。（即标准库中所有执行阻塞型I/O操作的函数，在等待操作系统返回结果时都会释放GIL。）time.sleep()函数也会释放GIL。因此，尽管有GIL，Python线程还是能在I/O密集型应用中发挥作用。

