import shelve

'''
shelve类似一个key-value数据库，可以很方便的用来保存Python的内存对象，其内部使用pickle来序列化数据，
'''

def test_shelve():
    # open 返回一个Shelf类的实例
    #
    # 参数flag的取值范围：
    # 'r': 只读打开
    # 'w': 读写访问
    # 'c': 读写访问，如果不存在则创建
    # 'n': 读写访问，总是创建新的、空的数据库文件
    #
    # protocol: 与pickle库一致
    # writeback: 为True时，当数据发生变化会回写，不过会导致内存开销比较大
    d = shelve.open('shelve.db', flag='c', protocol=2, writeback=False)
    assert isinstance(d, shelve.Shelf)

    # 在数据库中插入一条记录
    d['abc'] = {'name': ['a', 'b']}
    d.sync()

    print(d['abc'])

    # writeback是False,因此对value进行修改是不起作用的
    d['abc']['x'] = 'x'
    print(d['abc'])

    d['abc'] = 'xxx'
    print(d['abc'])

    d['abc'] = {'name': ['a', 'b']}
    d.close()

    print('----------------------------------------------')

    d = shelve.open('shelve.db', writeback=True)

    print(d['abc'])

    d['abc']['xxx'] = 'xxxx'
    print(d['abc'])
    d.close()

    print('----------------------------------------------')

    d = shelve.open('shelve.db')
    print(d['abc'])
    d.close()

if __name__ == '__main__':
    test_shelve()
