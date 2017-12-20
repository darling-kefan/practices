### 辅助函数，仅用于显示 ###

def cls_name(obj_or_cls):
    cls = type(obj_or_cls)
    if cls is type:
        cls = obj_or_cls
    return cls.__name__.split('.')[-1]

def display(obj):
    cls = type(obj)
    if cls is type:
        return '<class {}>'.format(obj.__name__)
    elif cls in [type(None), int]:
        return repr(obj)
    else:
        return '<{} object>'.format(cls_name(obj))

def print_args(name, *args):
    pseudo_args = ', '.join(display(x) for x in args)
    print('-> {}.__{}__({})'.format(cls_name(args[0]), name, pseudo_args))


### 对这个示例重要的类 ###


class Overriding:
    '''也称为描述符或强制描述符'''

    def __get__(self, instance, owner):
        print_args('get', self, instance, owner)

    def __set__(self, instance, value):
        print_args('set', self, instance, value)


class OverridingNoGet:
    '''没有``__get__``方法的覆盖型描述符'''

    def __set__(self, instance, value):
        print_args('set', self, instance, value)


class NonOverriding:
    '''也称为数据描述符或遮盖型描述符'''

    def __get__(self, instance, owner):
        print_args('get', self, instance, owner)


class Managed:
    over = Overriding()
    over_no_get = OverridingNoGet()
    non_over = NonOverriding()

    def spam(self):
        print('-> Managed.spam({})'.format(display(self)))


'''
在类中定义的函数属于绑定方法(bound method), 因为用户定义的函数都有__init__方法，所以
依附到类上时，就相当于描述符。

obj.spam和Managed.spam获取的是不同的对象。与描述符一样，通过托管类访问时，函数的__get__方法会返回自身的引用。但是，通过实例访问时，函数的__get__方法返回的是绑定方法对象：一种可调用的对象，里面包装着函数，并把托管实例(例如obj)绑定给函数的第一个参数，这与functools.partial函数的行为一致。

描述符类的实例能用作托管类的属性。

'''
