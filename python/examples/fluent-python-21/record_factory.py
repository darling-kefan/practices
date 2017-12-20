def record_factory(cls_name, field_names):
    try:
        field_names = field_names.replace(',', ' ').split()
    except AttributeError:
        pass
    field_names = tuple(field_names)

    def __init__(self, *args, **kwargs):
        attrs = dict(zip(self.__slots__, args))
        attrs.update(kwargs)
        for name, value in attrs.items():
            setattr(self, name, value)

    def __iter__(self):
        for name in self.__slots__:
            yield getattr(self, name)

    def __repr__(self):
        value = ', '.join('{}={!r}'.format(*i) for i
                          in zip(self.__slots__, self))
        return '{}({})'.format(self.__class__.__name__, values)

    cls_attrs = dict(__slots__ = field_names,
                     __init__ = __init__,
                     __iter__ = __iter__,
                     __repr__ = __repr__)

    return type(cls_name, (object,), cls_attrs)





'''
#--------- 解释器会编译函数，但是不会执行定义体 ---------
#--------- 解释器会编译类，并会执行类定义体，包括嵌套的类 ---------
<[100]> evalsupport module start
<[400]> MetaAleph body
<[700]> evalsupport module end
<[1]> evaltime module start
<[2]> ClassOne body
<[6]> ClassTwo body
<[7]> ClassThree body
<[200]> deco_alpha
<[9]> ClassFour body

<[11]> ClassOne tests ....................................
<[3]> ClassOne.__init__
<[5]> ClassOne.method_x
<[12]> ClassThree tests ...................................
<[300]> deco_alpha:inner_1
<[13]> ClassFour tests ...................................
<[10]> ClassFour.method_y

<[14]> evaltime module end

<[4]> ClassOne.__del__

导入时：
> 解释器会执行所导入模块及其依赖中的每个类定义体；
> 解释器先计算类的定义体，然后调用依附在类上的装饰器函数，这是合理的行为，因为必须先构建类对象，装饰器才会有类对象可处理。

根据Python对象模型，类是对象，因此类肯定是另外某个类的实例。默认情况下，Python中的类是type类的实例。也就是说，type是大部分内置类和用户定义的类的元类



'''
