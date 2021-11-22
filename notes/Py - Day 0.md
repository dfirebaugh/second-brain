# Day 0
## Learning
Some notes about [[Learning]]

## Functions
Functions are a grouping of instructions.

> tip: functions can be as large as you want, but it's generally a good idea that  functions should do one thing.

> tip: Give functions descriptive names.  Usually verbs are good for function names.

```python
def hello(name):
	return 'Hello {}'.format(name)
```


### Anatomy of a Function
Let's analyze our `hello` function.
```python
def hello(name) -> str:
	return 'Hello {}'.format(name)

hello('Dustin')
```

* `def`  - keyword to define a function
* `hello` - the functions name
* `(name)` - the arguments (i.e. what we pass into the function)
* `-> str` - return type declaration (optional)
* function body - the set of instructions inside the function
* return statement (optional)
* `hello('Dustin')` - function call

## Operators
[python operators](https://www.w3schools.com/python/python_operators.asp)

## Branching
In programming [branching](https://magpi.raspberrypi.com/articles/branching-if-else-python) refers to when the processing flow of a program can go one way or another.

Basically, if conditions.

```python
if False:
    print("The first block of code ran")
elif True:
    print("The second block of code ran")
else:
    print("The third block of code ran")
```

## Recursion
Recursion is when a function calls itself until it doesn't.

To achieve this, you will have a base case that must eventually get called.

for example:
```python
def recurse(num):
	if num == 10:
		return
	else:
		print(num)
		return recurse(num+1)
```

The base case is `if num == 10`.  If the base case is not met, we call the function again, except we increase the passed in number by 1.

## Exercises

#### Hello, World
Write a program that prints ‘Hello World’ to the screen.
#### Guessing Game
Write a guessing game where the user has to guess a secret number. After every guess the program tells the user whether their number was too large or too small. At the end the number of tries needed should be printed. It counts only as one try if they input the same number multiple times consecutively.
#### FizzBuzz
Write a program that prints out all numbers from 1 to 100.
If the number is divisible by 3, print out Fizz.
If the number if divisible by 5, print out Buzz.


## Tips

### Try to avoid nesting
You won't be able to avoid nesting entirely...

It may help you think about how to write the function if handle your edge cases at the top.  Then the `thing` that the function actually does would be at the end of the function.

This can reduce the amount of nesting and make it easier for someone to figure out what the function actually does.

If you're finding yourself wanted to do a bunch of nesting in the function, it may be an indication that you need to write additional functions.

e.g.

```python
def is_odd(num):
  return num % 2 != 0

def print_numbers(current_number):
  if current_number == 100:
    return

  if is_odd(current_number):
    print('\033[92m', end='')
  else:
    print('\033[94m', end='')

  print(current_number)
  print_numbers(current_number+1)


print_numbers(1)
```

### Try to make it human readable
> “... Clean code reads like well-written prose..." -- Grady Booch author of Object  
Oriented Analysis and Design with  
Applications”

It's important to try to communicate what the code does in a human readable way.

Rather than writing something like

```python
if num % 2 != 0:
	// do something
```

it would read better as:
```python
if is_odd(num):
	// do something
```
