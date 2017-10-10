#! /usr/bin/env python3
# -*- coding:utf-8 -*-

class Animal(object):
    def run(self):
        print('Animal is running...')

class Dog(Animal):
    def run(self):
        print('Dog is running...')

    def eat(self):
        print('Eating meat...')

class Cat(Animal):
    def run(self):
        print('Cat is running...')

def run_twice(animal):
    animal.run()
    animal.run()

if __name__ == '__main__':
    a = Animal()
    d = Dog()
    c = Cat()

    a.run()
    d.run()
    d.eat()
    c.run()

    print(isinstance(a, Animal))
    print(isinstance(d, Dog))
    print(isinstance(d, Animal))

    run_twice(a)
    run_twice(d)
    run_twice(c)
