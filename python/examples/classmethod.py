class A:
    name = 'polarsnow'

    def __init__(self, name):
        self.name = name

    def f(self):
        return self.name

    @classmethod
    def f2(cls):
        return cls.name

    @staticmethod
    def f3():
        return A.name

if __name__ == '__main__':
    a = A('lvrui')
    print(a.f())
    print(a.f2())
    print(a.f3())
    print(A.f3())
    print(A.f2())
