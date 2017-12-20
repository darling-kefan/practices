class Foo(object):
    def __init__(self, a, b):
        print(type(self).__name__)
        self.__x = a
        self.__y = b

    @property
    def x(self):
        return self.__x

    @property
    def y(self):
        return self.__y

class SonFoo(Foo):
    # 构造方法不兼容，就尽量不要使用继承
    def __init__(self, l=None):
        if l is None:
            self.l = []
        else:
            self.l = list(l)

    def __iter__(self):
        return iter([self.x,self.y])

    def __repr__(self):
        return '<{!r}, {!r}>'.format(*self)
