def quantity():
    '''特性工厂函数
    '''

    # 函数属性会被保留，无论函数被调用多少次
    try:
        quantity.counter += 1
    except AttributeError:
        quantity.counter = 0

    storage_name = '_{}:{}'.format('quantity', quantity.counter)
    print(storage_name)

    def qty_getter(instance):
        return getattr(instance, storage_name)

    def qty_setter(instance, value):
        if value > 0:
            setattr(instance, storage_name, value)
        else:
            raise ValueError('value must be > 0')

    return property(qty_getter, qty_setter)


if __name__ == '__main__':
    a = quantity()
    b = quantity()
    c = quantity()
