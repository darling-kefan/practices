class Quantity:
    __counter = 0

    def __init__(self):
        cls = self.__class__
        name = cls.__name__
        self._storagename = '__{}#{}'.format(name, cls.__counter)
        cls.__counter += 1

    def __set__(self, instance, value):
        if value > 0:
            setattr(instance, self._storagename, value)
        else:
            raise ValueError('value must be > 0')

    def __get__(self, instance, owner):
        return getattr(instance, self._storagename)


class LineItem:
    weight = Quantity()
    price = Quantity()

    def __init__(self, description, weight, price):
        self.description = description
        self.weight = weight
        self.price = price

    def subtotal(self):
        return self.weight * self.price

if __name__ == '__main__':
    print(type(LineItem.weight), LineItem.weight, type(LineItem.price), LineItem.price)
