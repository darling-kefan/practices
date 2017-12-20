#import model_v4c as model
import bulkfood_v4prop as model

class LineItem:
    # 使用model.Quantity描述符
    weight = model.quantity()
    price = model.quantity()

    def __init__(self, description, weight, price):
        self.description = description
        self.weight = weight
        self.price = price

    def subtotal(self):
        return self.weight * self.price

if __name__ == '__main__':
    li = LineItem('hahaha', 10, 1.23)
    print(li.weight, li.price, li.subtotal())
    print(getattr(li, '_quantity:0'), getattr(li, '_quantity:1'))
