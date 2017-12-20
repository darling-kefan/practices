'''
统一时刻，内存中可以能有几千个LineItem实例，不过只会有两个描述符实例： LineItem.weight和LineItem.price。因此，存储在描述符实例中的数据，其实会变成LineItem类的类属性，从而全部LineItem实例共享。
'''

class Quantity:

    def __init__(self, storage_name):
        self.storage_name = storage_name

    def __set__(self, instance, value):
        if value > 0:
            instance.__dict__[self.storage_name] = value
        else:
            raise ValueError('value must be > 0')

class LineItem:
    weight = Quantity('weight')
    price = Quantity('price')

    def __init__(self, description, weight, price):
        self.description = description
        self.weight = weight
        self.price = price

    def subtotal(self):
        return self.weight * self.price

if __name__ == '__main__':
    li = LineItem('hello world', 10, 20)
    print(type(LineItem.weight), LineItem.weight, type(li.weight), li.weight)
    print(li.subtotal())
