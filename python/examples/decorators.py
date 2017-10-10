#! /usr/bin/env python3.5
# -*- coding:utf-8 -*-

import functools

def benchmark(func):
    """
    A decorator that prints the time a function takes
    to execute.
    """
    import time
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        t = time.clock()
        res = func(*args, **kwargs)
        print("{0} {1}".format(func.__name__, time.clock()-t))
        return res
    return wrapper

def logging(func):
    """
    A decorator that logs the activity of the script.
    (it actually just prints it, but it could be logging!)
    """
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        res = func(*args, **kwargs)
        print("{0} {1} {2}".format(func.__name__, args, kwargs))
        return res
    return wrapper

def counter(func):
    """
    A decorator that counts and prints the number of the times a function has been executed.
    """
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        wrapper.count = wrapper.count + 1
        res = func(*args, **kwargs)
        print("{0} has been used: {1}x".format(func.__name__, wrapper.count))
        return res
    wrapper.count = 0
    return wrapper

@counter
@benchmark
@logging
def reverse_string(string):
    #return str(reversed(string))
    return ''.join(reversed(string))

print(reverse_string("Able was I are I Saw Elba"))
print(reverse_string("A man, a plan, a canoe, pasta, heros, rajahs."))

@counter
@benchmark
@logging
def get_random_futurama_quote():
    from urllib import urlopen
    result = urlopen("http://subfusion.net/cgi-bin/quote.pl?quote=futurama").read()
    try:
        value = result.split("<br><b><hr><br>")[1].split("<br><br><hr>")[0]
        return values.strip()
    except:
        return "No, I'm ... doesn't!"

print(get_random_futurama_quote())
print(get_random_futurama_quote())

exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

# For debugging, the stacktrace prints you the function __name__
def foo():
    print("foo")

print(foo.__name__)
# outputs: foo

# With a decorator, it gets messy
def bar(func):
    def wrapper():
        print("bar")
        return func()
    return wrapper

@bar
def foo():
    print("foo")

print(foo.__name__)
# outputs: wrapper

import functools

def bar(func):
    # We say that "wrapper", is wrapping "func"
    # and the magic begins
    @functools.wraps(func)
    def wrapper():
        print("bar")
        return func()
    return wrapper

@bar
def foo():
    print("foo")

print(foo.__name__)
# outputs: foo

exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

def decorator_with_args(decorator_to_enhance):
    """
    This function is supposed to be used as a decorator.
    It must decorate an other function, that is intended to be used as a decorator.
    Take a cup of coffee.
    It will allow any decorator to accept an arbitrary number of arguments,
    saving you the headache to remember how to do that every time.
    """

    # We use the same trick we did to pass arguments
    def decorator_maker(*args, **kwargs):

        # We create on the fly a decorator that accepts only a function
        # but keeps the passed arguments from the maker.
        def decorator_wrapper(func):

            # We return the result of the original decorator, which, after all,
            # IS JUST AN ORDINARY FUNCTION (which returns a function).
            # Only pitfall: the decorator must have this specific signature of it won't work:
            return decorator_to_enhance(func, *args, **kwargs)

        return decorator_wrapper

    return decorator_maker

# You create the function you will uses as a decorator. Add stick a decorator on it
# Don't forget, the signature is "decorator(func, *args, **kwargs)"
@decorator_with_args
def decorated_decorator(func, *args, **kwargs):
    def wrapper(function_arg1, function_arg2):
        print("Decorated with {0} {1}".format(args, kwargs))
        return func(function_arg1, function_arg2)
    return wrapper

# Then you decorate the functions you wish with your brand new decoreted decorator.

@decorated_decorator(42, 404, 1024)
def decorated_function(function_arg1, function_arg2):
    print("Hello {0} {1}".format(function_arg1, function_arg2))

decorated_function("Universe and", "everything")


exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

def decorator_maker_with_arguments(decorator_arg1, decorator_arg2):

    print("I make decorators! And I accept arguments: {0}, {1}".format(decorator_arg1, decorator_arg2))

    def my_decorator(func):
        # The ability to pass arguments here is a grift from closures.
        # Objects are data with methods attached, closures are functions with data attached.
        print("I am the decorator. Somehow you passed me arguments: {0}, {1}".format(decorator_arg1, decorator_arg2))

        # Don't confuse arguments and function arguments
        def wrapped(function_arg1, function_arg2):
            print("I am the wrapper around the decorated function.\n"
                  "I can access all the variables\n"
                  "\t- from the decorator: {0} {1}\n"
                  "\t- from the function call: {2} {3}\n"
                  "Then I can pass them to the decorated function"
                  .format(decorator_arg1, decorator_arg2, function_arg1, function_arg2))
            return func(function_arg1, function_arg2)

        return wrapped

    return my_decorator

# @decorator_maker_with_arguments("Leonard", "Sheldon")
# def decorated_function_with_arguments(function_arg1, function_arg2):
#     print("I am the decorated function and only knows about my arguments: {0}"
#           " {1}".format(function_arg1, function_arg2))
# 
# decorated_function_with_arguments("Rajesh", "Howard")

c1 = "Penny"
c2 = "Leslie"

@decorator_maker_with_arguments("Leonard", c1)
def decorated_function_with_arguments(function_arg1, function_arg2):
    print("I am the decorated function and only knows about my arguments:"
          " {0} {1}".format(function_arg1, function_arg2))

decorated_function_with_arguments(c2, "Howard")



exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

def decorator_maker():

    print("I make decorators! I am executed only once: "
          "when you make me create a decorator.")

    def my_decorator(func):
        print("I am a decorator! I am executed only when you decorate a function.")
        def wrapped():
            print("I am the wrapper around the decorated function. "
                  "I am called when you call the the decorated function."
                  "As the wrapper, I return the result of the decorated function.")
            return func()

        print("As the decorator, I return the wrapped function.")

        return wrapped

    print("As a decorator maker, I return a decorator")
    return my_decorator

##### The First Way
# # Let's create a decorator. It's just a new function after all.
# new_decorator = decorator_maker()
# # outputs:
# 
# # Then we decorate the function
# def decorated_function():
#     print("I am the decorated function.")
# 
# decorated_function = new_decorator(decorated_function)
# 
# # Let's call the decorated function
# decorated_function()

##### The Second Way
# # Then we decorate the function
# def decorated_function():
#     print("I am the decorated function.")
# 
# decorated_function = decorator_maker()(decorator_maker)
# 
# decorated_function()

##### The Third Way
@decorator_maker()
def decorated_function():
    print("I am the decorated function.")

decorated_function()


exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

# Decorators are ordinary functions
def my_decorator(func):
    print("I am an ordinary function")
    def wrapper():
        print("I am function returned")
        func()
    return wrapper

# Therefor, you can call it without any "@"
def lazy_function():
    print("zzzzzzzzzzzzzzzzzz")

decorated_function = my_decorator(lazy_function)
# outputs: I am an ordinary function

# I outputs "I am an ordinary function", because that's just what you do:
# calling a function. Nothing magic.

@my_decorator
def lazy_function():
    print('zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz')

exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

def a_decorator_passing_arbitrary_arguments(function_to_decorate):
    # The wrapper accepts any arguments
    def a_wrapper_accepting_arbitrary_arguments(*args, **kw):
        print("Do I have args?:")
        print(args)
        print(kw)
        # Then you unpack the arguments, here *args, **kw
        function_to_decorate(*args, **kw)
    return a_wrapper_accepting_arbitrary_arguments

@a_decorator_passing_arbitrary_arguments
def function_with_no_argument():
    print('Python is cool, no argument here.')

function_with_no_argument()
# outputs:
# Do I have args?:
# ()
# {}
# Python is cool, no argument here.

@a_decorator_passing_arbitrary_arguments
def function_with_arguments(a, b, c):
    print(a, b, c)

function_with_arguments(1, 2, 3)
# outputs:
# Do I have args?:
# (1, 2, 3)
# {}
# 1 2 3

@a_decorator_passing_arbitrary_arguments
def function_with_named_arguments(a, b, c, platypus="Why not ?"):
    print("Do {0}, {1} and {2} like platypus? {3}".format(a, b, c, platypus))

function_with_named_arguments("Bill", "Linus", "Steve")
# outputs:
# Do I have args?:
# ("Bill", "Linus", "Steve")
# {"platypus": "Indeed!"}
# Do Bill, Linus and Steve like platypus? Indeed!

class Mary(object):
    def __init__(self):
        self.age = 31

    @a_decorator_passing_arbitrary_arguments
    def sayYourAge(self, lie=-3):
        print("I am {0}, what did you think?".format(self.age + lie))

m = Mary()
m.sayYourAge()

exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

def method_friendly_decorator(method_to_decorate):
    def wrapper(self, lie):
        lie = lie - 3
        return method_to_decorate(self, lie)
    return wrapper

class Lucy(object):
    def __init__(self):
        self.age = 32

    @method_friendly_decorator
    def sayYourAge(self, lie):
        print("I am {0}, what did you think?".format(self.age + lie))

l = Lucy()
l.sayYourAge(-3)
# outputs:
# I am 26, what did you think?

exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

def a_decorator_passing_arguments(function_to_decorator):
    def a_wrapper_accepting_arguments(arg1, arg2):
        print("I got args! Look: {0}, {1}".format(arg1, arg2))
        function_to_decorator(arg1, arg2)
    return a_wrapper_accepting_arguments

@a_decorator_passing_arguments
def print_full_name(first_name, last_name):
    print("My name is {0} {1}".format(first_name, last_name))

print_full_name("Peter", "Venkman")


exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

def bread(func):
    def wrapper():
        print("</''''''\>")
        func()
        print("<\______/>")
    return wrapper

def ingredients(func):
    def wrapper():
        print("#tomatoes#")
        func()
        print("~salad~")
    return wrapper


# def sandwich(food="--ham--"):
#     print(food)
# 
# sandwich()
# # outputs: --ham--
# 
# sandwich = bread(ingredients(sandwich))
# sandwich()
# # outputs:
# # </''''''\>
# # #tomatoes
# # --ham--
# # ~salad~
# # <\______/>

# Using the Python descorator syntax:
@bread
@ingredients
def sandwich(food="--ham--"):
    print(food)

sandwich()
# outputs:
# </''''''\>
# #tomatoes
# --ham--
# ~salad~
# <\______/>

@ingredients
@bread
def strange_sandwich(food="--ham--"):
    print(food)

strange_sandwich()
# outputs:
# #tomatoes
# </'''''''\>
# --ham--
# <\______/>
# ~salad~

exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

# A decorator is a function that expects ANOTHOR function as a parameter
def my_shiny_new_decorator(a_function_to_decorate):
    # Inside, the decorator defines a function on the fly: the wrapper.
    # This function is going to be wrapped around the original function
    # so it can execute code before and after it.
    def the_wrapper_around_the_original_function():
        # Put here the code you want to be executed BEFORE the original function is called
        print('Before the function runs')

        # Call the function here(using parenthese)
        a_function_to_decorate()

        # Put here the code you want to be executed AFTER the original function is called
        print('After the function runs')

    # At this point, "a_function_to_decorate" HAS NEVER BEEN EXECUTED.
    # We return the wrapper function we have just created.
    # The wrapper contains the function and the code to execute before and after.
    return the_wrapper_around_the_original_function

def a_stand_alone_function():
    print("I am a stand alone function, don't you dare modify me")

a_stand_alone_function()

a_stand_alone_function_decorated = my_shiny_new_decorator(a_stand_alone_function)
a_stand_alone_function_decorated()


@my_shiny_new_decorator
def another_stand_alone_function():
    print("Leave me alone")
# @my_shiny_new_decorator 等价于 another_stand_alone_function = my_shiny_new_decorator(another_stand_alone_function)
# @my_shiny_new_decorator 表示用my_shiny_new_decorator函数包裹another_stand_alone_function
another_stand_alone_function()

exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

def getTalk(kind='shout'):
    def shout(word='yes'):
        return word.capitalize() + '!'

    def whisper(word='yes'):
        return word.lower()+'...';

    # Then return one of them
    if kind == 'shout':
        # We don't use "()", we are not calling the function,
        # we are returning the function object.
        return shout
    else:
        return whisper

# How do you use this strange beast?

# Get the function object and assign it to a variable
talk = getTalk()

# You can see that "talk" is here a function object:
print(talk)
# outputs: <function shout at 0xb7ea817c>

# The object is the one returned by the function:
print(talk())
# outputs: Yes!

# And you can even use it directly if you feel wild:
print(getTalk('whisper')('no'))
# outputs: no...

exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

def talk():
    # You can define a function on the fly in "talk" ...
    def whisper(word="yes"):
        return word.lower()+"..."

    # ... and use it right away!
    print(whisper())

# You call "talk", that defines "whisper" EVERY TIME you call it, then
# "whisper" is called in "talk"
talk()
# outputs:
# "yes..."

# But "whisper" DOES NOT EXIST outside "talk":

try:
    print(whisper())
except NameError as e:
    print(e)

exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

def shout(word = "yes"):
    return word.capitalize() + "!"

print(shout())
# outputs: 'Yes!'

# As an object, you can assign the function to a variable like any other
# object.
scream = shout

print(scream())
# outputs: 'Yes!'

# More than that, it means you can remove the old name 'shout',
# and the function will still be accessible from 'scream'
del shout
try:
    print(shout())
except NameError as e:
    print(e)

print(scream())
# outputs: 'Yes!'


exit()

# ---------------------------------------------------------------------
# ---------------------------------------------------------------------

def makebold(fn):
    def wrapped():
        return "<b>" + fn() + "</b>"
    return wrapped

def makeitalic(fn):
    def wrapped():
        return "<i>" + fn() + "</i>"
    return wrapped

@makebold
@makeitalic
def hello():
    return "hello world"

# hello = makebold(makeitalic(hello))

print(hello())


