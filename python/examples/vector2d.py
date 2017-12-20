from array import array
import math

class Vector2d(object):
    typecode = 'd'

    def __init__(self, x, y):
        # 定义私有属性
        self.__x = float(x)
        self.__y = float(y)

    # @property装饰器把读值方法标记为特性
    @property
    def x(self):
        return self.__x

    @property
    def y(self):
        return self.__y

    # 把Vector2d实例变成可迭代的对象，这样才能拆包
    def __iter__(self):
        return (i for i in (self.__x, self.__y))
        # return (self.x, self.y)

    def __repr__(self):
        # return 'Vector2d({0}, {1})'.format(self.__x, self.__y)
        class_name = type(self).__name__
        return '{}({!r}, {!r})'.format(class_name, *self)

    def __str__(self):
        # return '({0}, {1})'.format(self.__x, self.__y)
        return str(tuple(self))

    def __bytes__(self):
        return (bytes([ord(self.typecode)]) +
                bytes(array(self.typecode, self)))

    def __eq__(self, other):
        return tuple(self) == tuple(other)

    def __hash__(self):
        return hash(self.x) ^ hash(self.y)

    def __abs__(self):
        return math.hypot(self.x, self.y)

    def __bool__(self):
        return bool(abs(self))

    def angle(self):
        return math.atan2(self.y, self.x)

    def __format__(self, fmt_spec=''):
        if fmt_spec.endswith('p'):
            fmt_spec = fmt_spec[:-1]
            coords = (abs(self), self.angle())
            outer_fmt = '<{}, {}>'
        else:
            coords = self
            outer_fmt = '({}, {})'
        # 创建一个可迭代的格式化字符串
        components = (format(c, fmt_spec) for c in coords)
        return outer_fmt.format(*components)

    @classmethod
    def frombytes(cls, octets):
        print(cls, octets)
        typecode = chr(octets[0])
        memv = memoryview(octets[1:]).cast(typecode)
        return cls(*memv)
