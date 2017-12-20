class Test(object):
    def __init__(self, x):
        self.x = x

    def printX(self):
        print(self.x)


def __init__(self, x):
    self.x = x

def printX(self):
    print(self.x)

Test1 = type('Test1', (object,), {"__init__": __init__, "printX": printX})

if __name__ == '__main__':
    t = Test(1)
    t.printX()

    t2 = Test1(2)
    t2.printX()
