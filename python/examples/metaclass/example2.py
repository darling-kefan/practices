class MetaClass(type):
    def __new__(meta, classname, bases, classDict):
        print('Class Name:', classname)
        print('Bases:', bases)
        print('Class Attributes', classDict)
        return type.__new__(meta, classname, bases, classDict)


class Test(object):

    __metaclass__ = MetaClass

    def __init__(self):
        pass

    def method(self):
        pass

    classAttribute = 'Something'


if __name__ == '__main__':
    t = Test()
    print(t.classAttribute)
