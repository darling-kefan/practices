import model_v5 as model

class LineItem:
    description = model.NonBlank()
    weight = model.Quantity()
    price = model.Quantity()

    def __init__(self, description, weight, price):
        self.description = description
        self.weight = weight
        self.price = price

    def subtotal(self):
        return self.weight * self.price

if __name__ == '__main__':
    li = LineItem('hahaha', 10, 1.23)
    print(li.weight, li.price, li.subtotal())
    print(getattr(li, '_NonBlank#0'), getattr(li, '_Quantity#1'))

    '''
    本章所举的几个LineItem示例演示了描述符的典型用途---管理数据属性。这种描述符也叫覆盖型描述符，因为描述符的__set__方法使用托管实例中的同名属性覆盖了要设置的属性。
    '''
