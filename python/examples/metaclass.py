#! /usr/bin/env python3.5
# -*- coding:utf-8 -*-


# -----------------------------------------
# Creating classes dynamically
# -----------------------------------------

def choose_class(name):
    if name == 'foo':
        class Foo(object):
            pass
        return Foo # return the class, not an instance
    else:
        class Bar(object):
            pass
        return Bar

MyClass = choose_class('foo')
print(MyClass)
print(MyClass())

class ObjectCreator(object):
    pass

print(type(1))
print(type("1"))
print(type(ObjectCreator))
print(type(ObjectCreator()))

exit()

# -----------------------------------------
# Classes as objects
# -----------------------------------------

class ObjectCreator(object):
    pass

my_object = ObjectCreator()
print(my_object)


# You can print a class because it's an object
print(ObjectCreator)

def echo(o):
    print(o)

# You can pass a class as a parameter
echo(ObjectCreator)

print(hasattr(ObjectCreator, 'new_attribute'))

# You can add attributes to a class
ObjectCreator.new_attribute = 'foo'

print(hasattr(ObjectCreator, 'new_attribute'))

# you can assign a class to a variable
ObjectCreatorMirror = ObjectCreator

print(ObjectCreatorMirror.new_attribute)

print(ObjectCreatorMirror())
