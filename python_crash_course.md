# 코딩 테스트를 위한 자료구조 & 알고리즘 공부


* [1. Getting Started](#getting-started)
* [2. Variables and Simple Data Types](#variables-and-simple-data-types)
* [3. Introducing Lists](#introducing-lists)
* [4. Working with Lists](#working-with-lists)
* [5. if Statements](#if-statements)
* [6. Dictionaries](#dictionaries)
* [7. User Input and while loops](#user-input-and-while-loops)
* [8. Functions](#functions)
* [9. Classes](#classes)
* [10. Files and Exceptions](#files-and-exceptions)
* [11. Testing Your Code](#testing-your-code)
* [Project 1. Alien Invasion](#project-1-alien-invasion)
* [Project 2. Data Visualization](#project-2-data-visualization)
* [Project 3. Web Applications](#project-3-web-applications) 

# Getting Started


```bash
%%bash
# check installation
python3 --version

# compile and execute
python3 -m py_compile "chap02/max.py"
```

    Python 3.7.3


# Variables and Simple Data Types


```python
first_name = "ada"
last_name = "lovelace"
full_name = f"{first_name} {last_name}"

print(full_name.upper())
print(full_name.lower())
print(full_name.title())
print(f"Hello, {full_name.title()}!")
```

    ADA LOVELACE
    ada lovelace
    Ada Lovelace
    Hello, Ada Lovelace!


# White spaces


```python
greeting = '\tHello\n'

word = '  python  '
word.rstrip()
word.lstrip()
word.strip()
```




    'python'



# Quote inside quote


```python
msg_valid = "monty's best friend"
msg2_valid = 'I said, \'Hi Monty.\''
# msg_invalid = 'monty's best friend' # Error!
```

## Operator


```python
# Arithmetic operator + - * / and Exponential operator **
bool_0 = (10 + 2 * 3 / 3 - 5 == 7)
bool_1 = (1024 == 2 ** 10)

# Floating number
.1 + .1 # .2000000000001
```




    0.2



## Avoid type errors


```python
age = 23
message = "Happy " + str(age) + "rd Birthday!"
```

- Integers, Floats


```python
3 ** 2 # 9
2 + 3*4 # 14
0.1 + 0.1 # .2
0.1 + 0.2 # .30000...004
5 * 0.1 # .50000...004
2 * 0.1 # .2

# When you divide any two numbers, even if they are integers that result in a whole number, you’ll always get a float:
3/2 # 1.5
3.0/2 # 1.5
3/2.0 # 1.5
3.0/2.0 # 1.5
```




    1.5



- Underscores in Numbers


```python
# Make more readable
universe_age = 14_000_000_000
# 14000000000
print(universe_age)

```

    14000000000


- Multiple Assignment


```python
x, y, z = 0, 0, 0
```

- Constants


```python
MAX_CONNECTIONS = 5000
```

- Comments


```python
# python line comment
"""
python
block comment
"""
```




    '\npython\nblock comment\n'



- The Zen of Python


```python
import this
```

    The Zen of Python, by Tim Peters
    
    Beautiful is better than ugly.
    Explicit is better than implicit.
    Simple is better than complex.
    Complex is better than complicated.
    Flat is better than nested.
    Sparse is better than dense.
    Readability counts.
    Special cases aren't special enough to break the rules.
    Although practicality beats purity.
    Errors should never pass silently.
    Unless explicitly silenced.
    In the face of ambiguity, refuse the temptation to guess.
    There should be one-- and preferably only one --obvious way to do it.
    Although that way may not be obvious at first unless you're Dutch.
    Now is better than never.
    Although never is often better than *right* now.
    If the implementation is hard to explain, it's a bad idea.
    If the implementation is easy to explain, it may be a good idea.
    Namespaces are one honking great idea -- let's do more of those!


# Lists


```python
bicycles = ['trek', 'cannondale', 'redline', 'specialized']
print(bicycles) # print entire list
print(bicycles[0].title()) # 'Trek'
print(bicycles[-1]) # 'specialized'
bicycles[2] = 'electric bike' # change an element
```

    ['trek', 'cannondale', 'redline', 'specialized']
    Trek
    specialized


## List methods


```python
fruit = ['apple', 'orange', 'grapes', 'blueberry', 'watermelon']

# append(str)
fruit.append('strawberry') # append at the end

# insert(i, str)
fruit.insert(0, 'banana') # add at the beginning
fruit.insert(len(fruit), 'melon') # add (append) at the end

# del list[i]
del fruit[0] # remove first element

# pop()
motorcycles = ['honda', 'yamaha', 'suzuki']
last_owned = motorcycles.pop()
print("The last motorcycle I owned was a " + last_owned.title() + ".")

# pop(i) removes and returns the i th element
first_owned = motorcycles.pop(0)
print('The first motorcycle I owned was a ' + first_owned.title() + '.')

# remove(str) remove item (first occurrence) by specifying the value
# Have to use a loop to remove all occurrences
try:
    motorcycles.remove('honda')
except:
    print('honda DOES NOT EXIST')
else:
    print("No Error Occurred")
finally:
    print("remove() operation")

try:
    too_expensive = 'ducati'
    motorcycles.remove(too_expensive)
    print("\nA " + too_expensive.title() + " is too expensive for me.")
    print(motorcycles)
except:
    print('ducati DOES NOT EXIST')
else:
    print("No Error Occurred")
finally:
    print("remove() operation")


# PERMANENT SORT
# list.sort()
cars = ['bmw', 'audi', 'toyota', 'subaru']
cars.sort()
print(cars)  # audi bmw subaru toyota

# sort(reverse = True)
cars.sort(reverse = True) # sort and reverse

# TEMPORARY SORT
# sorted(list)
print('Sorted List: ')
print(sorted(cars))
print('Original List: ')
print(cars) # original list

# sorted(list, reverse = True)
location = ['tokyo', 'los angeles', 'new york', 'seoul']
print(sorted(location, reverse = True)) # in reverse order
print(location) # original list

# ALPHABETICAL SORT when some values are not in lowercase
location = ['tokyo', 'los angeles', 'New York', 'Seoul'] # reassign
# disregard lower/uppercase preferences
print(sorted(location, key=str.lower))
# ['los angeles', 'New York', 'Seoul', 'tokyo']

# PERMANENT REVERSE
# reverse()
cars.reverse() # reverse only; it does not sort
print(cars)

# len()
cars = ['bmw', 'audi', 'toyota', 'subaru']
len(cars)
```

    The last motorcycle I owned was a Suzuki.
    The first motorcycle I owned was a Honda.
    honda DOES NOT EXIST
    remove() operation
    ducati DOES NOT EXIST
    remove() operation
    ['audi', 'bmw', 'subaru', 'toyota']
    Sorted List: 
    ['audi', 'bmw', 'subaru', 'toyota']
    Original List: 
    ['toyota', 'subaru', 'bmw', 'audi']
    ['tokyo', 'seoul', 'new york', 'los angeles']
    ['tokyo', 'los angeles', 'new york', 'seoul']
    ['los angeles', 'New York', 'Seoul', 'tokyo']
    ['audi', 'bmw', 'subaru', 'toyota']





    4



- Avoiding index errors


```python
cars = ['honda', 'audi']
print(cars[2])
cars = []
print(cars[0])
```


    ---------------------------------------------------------------------------

    IndexError                                Traceback (most recent call last)

    <ipython-input-15-1ae47671d41a> in <module>
          1 cars = ['honda', 'audi']
    ----> 2 print(cars[2])
          3 cars = []
          4 print(cars[0])


    IndexError: list index out of range


# Working with Lists


```python
magicians = ['alice', 'david', 'carolina']
for magician in magicians:
    print(f"{magician.title()}, that was a great trick!")
    print(f"I can't wait to see your next trick, {magician.title()}.\n")

print("Thank you, everyone. That was a great magic show!")
```

    Alice, that was a great trick!
    I can't wait to see your next trick, Alice.
    
    David, that was a great trick!
    I can't wait to see your next trick, David.
    
    Carolina, that was a great trick!
    I can't wait to see your next trick, Carolina.
    
    Thank you, everyone. That was a great magic show!


- Avoiding Indentation Errors

missing or unnecessary indentations or colons

## Numerical Lists


```python
for value in range(1, 6):
    print(value) # 1 ~ 5

# [1, 2, 3, 4, 5]
numbers = list(range(1,6))
# [2, 4, 6, 8, 10]
even_numbers = list(range(2, 11, 2))

# square numbers from 1 to 11 as a list [1, 4, ... , 100] 
squares = []
for value in range(1, 11):
    square = value ** 2
    squares.append(square)

# min(list) max(list) sum(list)
digits = [1, 2, 3, 4, 5, 6, 7, 8, 9, 0]
min_val = min(digits)
max_val = max(digits)
sum_val = sum(digits)
```

    1
    2
    3
    4
    5


## List Comprehensions


```python
# allows you to generate the same list in just one line of code.
squares = [value ** 2 for value in range(1, 11)]
print(squares) # [1, 4, 9, 16, 25, 36, 49, 64, 81, 100]

cubes = [value ** 3 for value in range(1, 11)]
print(cubes) # [1, 8, 27, 64, 125, 216, 343, 512, 729, 1000]
```

    [1, 4, 9, 16, 25, 36, 49, 64, 81, 100]
    [1, 8, 27, 64, 125, 216, 343, 512, 729, 1000]


## Slicing a list


```python
players = ['charles', 'martina', 'michael', 'florence', 'eli']
print(players[0:3]) # ['charles', 'martina', 'michael']
print(players[:4]) # ['charles', 'martina', 'michael', 'florence']
print(players[2:]) # ['michael', 'florence', 'eli']

for player in players[:3]:
    print(player.title()) # Charles Martina Michael
```

    ['charles', 'martina', 'michael']
    ['charles', 'martina', 'michael', 'florence']
    ['michael', 'florence', 'eli']
    Charles
    Martina
    Michael


## Copying a list


```python
my_foods = ['pizza', 'falafel', 'carrot cake']
# copied list (distinct from original)
friend_foods = my_foods[:]

my_foods.append('cannoli')
friend_foods.append('ice cream')

# ['pizza', 'falafel', 'carrot cake', 'cannoli']
print(my_foods)
# ['pizza', 'falafel', 'carrot cake', 'ice cream']
print(friend_foods)

# Copying a list using assignment (not working as expected)
my_foods = ['pizza', 'carrot cake']
# both variables will point to the same list
friend_foods = my_foods

my_foods.append('cannoli')
friend_foods.append('ice cream')

# ['pizza', 'carrot cake', 'cannoli', 'ice cream']
print(my_foods)

# same result as above!
print(friend_foods)
```

    ['pizza', 'falafel', 'carrot cake', 'cannoli']
    ['pizza', 'falafel', 'carrot cake', 'ice cream']
    ['pizza', 'carrot cake', 'cannoli', 'ice cream']
    ['pizza', 'carrot cake', 'cannoli', 'ice cream']


## Tuples
tuples are simple data structures to store an immutable set of values 


```python
# sometimes you’ll want to create a list of items that cannot change.
# tuple is an immutable(values that cannot change) list;
dimensions = (200, 50)
for dimension in dimensions:
    print(dimension)

# Error! immutable!
# dimensions[0] = 250

# Assign new tuples
dimensions = (400, 100)
for dimension in dimensions:
    print(dimension) # 400 100
```

    200
    50
    400
    100


- Styling Your Code
1. Indent 4 spaces with 'tab' key to avoid unpredictable errors
2. Line length less than 72 or 80 
3. Blank line between the two sections (no more than 2 blank lines)
4. [PEP8 style guide](https://python.org/dev/peps/pep-0008/)

# if Statements


```python
cars = ['audi', 'bmw', 'subaru', 'toyota']
for car in cars:
    if car == 'bmw':
        print(car.upper())
    else:
        print(car.title())
```

    Audi
    BMW
    Subaru
    Toyota


- Conditional Tests


```python
age_0 = 22
age_1 = 18
if age_0 >= 21 and age_1 >= 21:
    print('both of you are 21 years old or older')
if age_0 >= 21 or age_1 >= 21:
    print('either of you are 21 years old or older')

requested_toppings = ['mushrooms', 'onions', 'pineapple']
if 'mushrooms' in requested_toppings:
    print('mushrooms are in the toppings')
if 'mushrooms' not in requested_toppings:
    print('mushrooms are not in the toppings')
```

    either of you are 21 years old or older
    mushrooms are in the toppings


- Boolean Expressions


```python
game_active = True
can_edit = False
```

- if elif else block


```python
age = 12
if age < 4:
    price = 0
elif age < 18:
    price = 5
elif age < 65:
    price = 10
else:
    price = 5

print("Your admission cost is $" + str(price) + ".")
# Sometimes it is clearer to use an additional elif, omitting else block
# elif age >= 65
```

    Your admission cost is $5.


- Testing Multiple conditions


```python
# check all conditions of interest, rather than checking one specific
requested_toppings = ['mushrooms', 'extra cheese']

if 'mushrooms' in requested_toppings:
    print("Adding mushrooms.")
if 'pepperoni' in requested_toppings:
    print("Adding pepperoni.")
if 'extra cheese' in requested_toppings:
    print("Adding extra cheese.")

print("\nFinished making your pizza!")
```

    Adding mushrooms.
    Adding extra cheese.
    
    Finished making your pizza!


- Using if statement with Lists


```python
requested_toppings = ['mushrooms', 'green peppers', 'extra cheese']
for topping in requested_toppings:
    if topping == 'green peppers':
        print("Sorry, we are out of green peppers right now.")
    else:
        print(f"Adding {topping}.")

print("\nFinished making your pizza!")
```

    Adding mushrooms.
    Sorry, we are out of green peppers right now.
    Adding extra cheese.
    
    Finished making your pizza!


- Checking that a list is not Empty


```python
requested_toppings = [] 
if requested_toppings: # returns True if non-empty list
    for requested_topping in requested_toppings:
        print("Adding " + requested_topping + ".")
    print("\nFinished making your pizza!")
else:
    print("Are you sure you want a plain pizza?")
```

    Are you sure you want a plain pizza?


- Using multiple lists


```python
available_toppings = ['mushrooms', 'olives', 'pepperoni', 'pineapple']
requested_toppings = ['mushrooms', 'french fries', 'extra cheese']
for topping in requested_toppings:
    if topping in available_toppings:
        print(f"Adding {topping}.")
    else:
        print(f"Sorry, we don't have {topping}.")

print("\nFinished making your pizza!")
```

    Adding mushrooms.
    Sorry, we don't have french fries.
    Sorry, we don't have extra cheese.
    
    Finished making your pizza!


- Try it yourself


```python
current_users = ['foo', 'bar']
new_users = ['FOo', 'baR', 'AAa', 'Bbb', 'ccc']
for user in new_users:
    if user.upper() in [user.upper() for user in current_users]:
        print(f"Hey! Your username, {user} is already in use!")
    else:
        print(f"Your username, {user} is available to use!")
```

    Hey! Your username, FOo is already in use!
    Hey! Your username, baR is already in use!
    Your username, AAa is available to use!
    Your username, Bbb is available to use!
    Your username, ccc is available to use!


# Dictionaries
collection of key-value pairs. Each key is connected to a value, and you can use a key to access the value associated with that key. In fact, you can use any object that you can create in Python as a value in a dictionary.


```python
alien = {'color': 'green', 'points': 5}
print(alien) # {'color': 'green', 'points': 5}
print(alien['color'])
print(alien['points'])

new_points = alien['points']
print("You just earned " + str(new_points) + " points!")

# Starting with an empty dictionary
alien = {}

# Adding New Key-Value pairs
alien['color'] = 'green'
alien['points'] = 5
alien['x_position'] = 0
alien['y_position'] = 25

alien['speed'] = 0

# Modifying Values in a dictionary
alien['color'] = 'yellow' # modify color

if alien['speed'] == 'slow':
    x_increment = 1
elif alien['speed'] == 'medium':
    x_increment = 2
elif alien['speed'] == 'fast':
    x_increment = 3
else:
    x_increment = 0

alien['x_position'] += x_increment # modify position
print("New x-position: " + str(alien['x_position']))
# Python variables are scoped to the innermost function or module; 
# control blocks like if and while blocks don't count. 
# (IIUC, this is also how JavaScript's var-declared variables work.)

# Removing Key-Value pairs
del alien['points'] # remove key-value pair

# A Dictionary of Similar Objects
favorite_languages = {
    'jen': 'python',
    'sarah': 'c',
    'edward': 'ruby',
    'phil': 'python',
}
print("Sarah's favorite language is "
      + f"{favorite_languages['sarah'].title()}.")
```

    {'color': 'green', 'points': 5}
    green
    5
    You just earned 5 points!
    New x-position: 0
    Sarah's favorite language is C.


## Looping Through a Dictionary


```python
# 1. Looping through all Key-Value pairs

# key-value pairs are not returned in the order they were stored 
# Python tracks only the connections between individual keys and their values
user_0 = {
    'username': 'foobar12341',
    'first': 'foo',
    'last': 'bar',
}
for key, value in user_0.items():
    print(f'Key: {key}')
    print(f'Value: {value}')

favorite_languages = {
    'jen': 'python',
    'sarah': 'c',
    'edward': 'ruby',
    'phil': 'python',
}
for name, language in favorite_languages.items():
    print(f"{name.title()}'s favorite language is {language.title()}.")

# 2-1. Looping Through All the Keys in a Dictionary
for name in favorite_languages.keys():
    print(name.title())

for name in favorite_languages: # is actually the default behavior when looping through a dictionary
    print(name.title())

friends = ['phil', 'sarah']
for name in favorite_languages.keys():
    print(name.title())
    if name in friends:
        print(" Hi " + name.title() 
        + ", I see your favorite language is " 
        + favorite_languages[name].title() + "!")

if 'erin' not in favorite_languages.keys():
    print("Erin, please take our poll!")

# 2-2. Looping Through a Dictionary’s Keys in Order
for name in sorted(favorite_languages.keys()):
    print(f"{name.title()}, thank you for taking the poll.")

# 3-1. Looping Through All Values in a Dictionary
# This approach pulls all the values from the dictionary without checking for repeats.
print("The following languages have been mentioned:")
for language in favorite_languages.values():
    print(language.title())

# 3-2. A set has unique items 
print("The following languages have been mentioned:")
for language in set(favorite_languages.values()):
    print(language.title())
```

    Key: username
    Value: foobar12341
    Key: first
    Value: foo
    Key: last
    Value: bar
    Jen's favorite language is Python.
    Sarah's favorite language is C.
    Edward's favorite language is Ruby.
    Phil's favorite language is Python.
    Jen
    Sarah
    Edward
    Phil
    Jen
    Sarah
    Edward
    Phil
    Jen
    Sarah
     Hi Sarah, I see your favorite language is C!
    Edward
    Phil
     Hi Phil, I see your favorite language is Python!
    Erin, please take our poll!
    Edward, thank you for taking the poll.
    Jen, thank you for taking the poll.
    Phil, thank you for taking the poll.
    Sarah, thank you for taking the poll.
    The following languages have been mentioned:
    Python
    C
    Ruby
    Python
    The following languages have been mentioned:
    Python
    Ruby
    C


## Nesting


```python
# 1. A List of Dictionaries
alien_0 = {'color': 'green', 'points': 5}
alien_1 = {'color': 'yellow', 'points': 10}
alien_2 = {'color': 'red', 'points': 15}
aliens = [alien_0, alien_1, alien_2]
for alien in aliens:
    print(alien)

# Make an empty list for storing aliens.
aliens = []
# Make 30 green aliens.
for alien_number in range(30):
    new_alien = {'color': 'green', 'points': 5, 'speed': 'slow'}
    aliens.append(new_alien)

print("Total number of aliens: " + str(len(aliens)))

# aliens are changing color and moving faster as the game progresses. 
# we can use a for loop and an if statement to change the color of aliens.
for alien in aliens[:3]:
    if alien['color'] == 'green':
        alien['color'] = 'yellow'
        alien['speed'] = 'medium'
        alien['point'] = 10
    elif alien['color'] == 'yellow':
        alien['color'] = 'red'
        alien['speed'] = 'fast'
        alien['point'] = 15
for alien in aliens[:5]:
    print(alien)

# 2. A List in a Dictionary
pizza = {
    'crust': 'thick',
    'toppings': ['mushroom', 'extra cheese'],
}
print('You ordered a ' + pizza['crust'] + '-crust pizza with the following toppings: ')
for topping in pizza['toppings']:
    print('\t' + topping)

favorite_languages = {
    'jen': ['python', 'ruby'],
    'sarah': ['c'],
    'edward': ['ruby', 'go'],
    'phil': ['python', 'haskell'],
}
for name, languages in favorite_languages.items():
    print('\n' + name.title() + "'s favorite languages are: ")
    for lang in languages:
        print('\t' + lang)
        
# You should not nest lists and dictionaries too deeply.

# 3. Dictionary in a Dictionary
users = {
    'aeinstein': {
        'first' : 'albert',
        'last': 'einstein',
        'location': 'princeton',
    },
    'mcurie': {
        'first': 'marie',
        'last': 'curie',
        'location': 'paris',
    }
}

for username, user_info in users.items():
    print('\nUsername: ' + username)
    full_name = user_info['first'] + user_info['last']
    
    print('\tFull name: ' + full_name.title())
    print('\tLocation: ' + user_info['location'].title())
```

    {'color': 'green', 'points': 5}
    {'color': 'yellow', 'points': 10}
    {'color': 'red', 'points': 15}
    Total number of aliens: 30
    {'color': 'yellow', 'points': 5, 'speed': 'medium', 'point': 10}
    {'color': 'yellow', 'points': 5, 'speed': 'medium', 'point': 10}
    {'color': 'yellow', 'points': 5, 'speed': 'medium', 'point': 10}
    {'color': 'green', 'points': 5, 'speed': 'slow'}
    {'color': 'green', 'points': 5, 'speed': 'slow'}
    You ordered a thick-crust pizza with the following toppings: 
    	mushroom
    	extra cheese
    
    Jen's favorite languages are: 
    	python
    	ruby
    
    Sarah's favorite languages are: 
    	c
    
    Edward's favorite languages are: 
    	ruby
    	go
    
    Phil's favorite languages are: 
    	python
    	haskell
    
    Username: aeinstein
    	Full name: Alberteinstein
    	Location: Princeton
    
    Username: mcurie
    	Full name: Mariecurie
    	Location: Paris


- Try it yourself


```python
# 6-9
favorite_places = {
    'foo': ['los angeles', 'house', 'home'],
    'bar': ['new york', 'mama', 'lalala'],
    'fugu': ['san francisco', 'dadada', 'wiwiwi'],
}
for key in favorite_places.keys():
    print(key + "'s favorite place: ")
    for place in favorite_places[key]:
        print('\t' + place)
```

    foo's favorite place: 
    	los angeles
    	house
    	home
    bar's favorite place: 
    	new york
    	mama
    	lalala
    fugu's favorite place: 
    	san francisco
    	dadada
    	wiwiwi


# User Input and while loops


```python
prompt = "If you tell us who you are, we can personalize the messages you see."
prompt += "\nWhat is your first name? "
name = input(prompt)
print("\nHello, " + name + "!")
```

    If you tell us who you are, we can personalize the messages you see.
    What is your first name? Lulu
    
    Hello, Lulu!


- int()


```python
height = input("How tall are you, in inches? ") # store string value '71' 
height = int(height) # converted to integer 71 
if height >= 36:
    print("\nYou're tall enough to ride!")
else:
    print("\nYou'll be able to ride when you're a little older.")
```

    How tall are you, in inches? 200
    
    You're tall enough to ride!


- The Modulo Operator


```python
number = input("Enter a number, and I'll tell you if it's even or odd: ")
number = int(number)
if number % 2 == 0:
    print(f"\nThe number {number} is even.")
else:
    print(f"\nThe number {number} is odd.")
```

    Enter a number, and I'll tell you if it's even or odd: 200
    
    The number 200 is even.


- Accepting Input in Python 2.7
  - you should use the raw_input() function when prompting for user input.

## Introducing while loops


```python
current_number = 1
while current_number <= 5:
    print(current_number)
    current_number += 1

prompt = "\nTell me something ('quit' or 'q' to end the program.)"
message = ""
while message != 'quit' and message != 'q':
    message = input(prompt)
    if message != 'quit' and message != 'q':
        print(message)

# Using a Flag
flag = True
while flag:
    message = input("\nEnter anything. 'q' to quit: ")
    if  message == 'q':
        flag = False
    else:
        print(message)
```

    1
    2
    3
    4
    5
    
    Tell me something ('quit' or 'q' to end the program.)Abc
    Abc
    
    Tell me something ('quit' or 'q' to end the program.)q
    
    Enter anything. 'q' to quit: aa
    aa
    
    Enter anything. 'q' to quit: q


## Use break to Exit a Loop


```python
while True:
    city = input("\nEnter a city you would like to visit. 'q' to quit: ")
    if city == 'q':
        break
    else:
        print('I would like to visit ' + city.title() + '!')
```

    
    Enter a city you would like to visit. 'q' to quit: Seoul
    I would like to visit Seoul!
    
    Enter a city you would like to visit. 'q' to quit: q


## Use continue in a Loop


```python
# return to the beginning of the loop based on the result of a conditional test 
current_number = 0
while current_number < 10:
    current_number += 1
    if current_number % 2 == 0:
        continue
    print(current_number)
```

    1
    3
    5
    7
    9


- Avoiding Infinite Loops
  - make sure a loop does not run forever!

## Using a while Loop with Lists and Dictionaries


```python
""" Moving Items from One List to Another """
unconfirmed_users = ['alice', 'brian', 'candace']
confirmed_users = []
while unconfirmed_users:
    current_user = unconfirmed_users.pop()
    print("Verifying user: " + current_user.title())
    confirmed_users.append(current_user)

print("\nThe following users have been confirmed:")
for confirmed_user in confirmed_users:
    print(confirmed_user.title())
```

    Verifying user: Candace
    Verifying user: Brian
    Verifying user: Alice
    
    The following users have been confirmed:
    Candace
    Brian
    Alice


- Removing All Instances of Specific Values from a List


```python
pets = ['dog', 'cat', 'dog', 'goldfish', 'cat', 'rabbit', 'cat']
while 'cat' in pets:
    pets.remove('cat')
print(pets)
```

    ['dog', 'dog', 'goldfish', 'rabbit']


- Filling a Dictionary with User Input


```python
responses = {}
active = True
while active:
    name = input("\nWhat's your name? ")
    mountain = input('Which mountain do you want to climb? ')
    responses[name] = mountain
    answer = input('do you want us to add another person? (yes/no) ')
    if answer.lower() == 'no':
        active = False
        
print('\nPolling result >>>>>>')
for name, mountain in responses.items():
    print(name + ' would like to climb ' + mountain + '.')
```

    
    What's your name? Lalal
    Which mountain do you want to climb? Fuji
    do you want us to add another person? (yes/no) no
    
    Polling result >>>>>>
    Lalal would like to climb Fuji.


- Try it yourself


```python
sandwich_orders = ['peanut', 'pastrami', 'cheese', 'pastrami',
                   'burrito', 'pastrami']

print('Sorry guys but, we ran out of pastrami sandwiches.')
while 'pastrami' in sandwich_orders:
    sandwich_orders.remove('pastrami')

finished_sandwiches = []
while sandwich_orders:
    order = sandwich_orders.pop()
    print("we're making your " + order + ' sandwich. Please Wait.')
    finished_sandwiches.append(order)

print('\nFinished dishes : ')
print('\t' + str(finished_sandwiches))
```

    Sorry guys but, we ran out of pastrami sandwiches.
    we're making your burrito sandwich. Please Wait.
    we're making your cheese sandwich. Please Wait.
    we're making your peanut sandwich. Please Wait.
    
    Finished dishes : 
    	['burrito', 'cheese', 'peanut']


# Functions

- pass information to functions.
- write functions whose job is to display information or process data and return a value or set of values. 
- store functions in separate files called modules to help organize your main program files.


```python
def greet_user():
    """Display a simple greeting"""
    print("Hello!")
greet_user()

# Docstrings are enclosed in triple quotes, which Python looks for 
# when it generates documentation for the functions in your programs
```

    Hello!


## Arguments and Parameters


```python
# the argument value 'jesse' is stored in the parameter username
def greet_user(username):
    print(f"Hello, {username.title()}!")
greet_user('jesse')
```

    Hello, Jesse!


- Positional Arguments


```python
# order matters in Positional Arguments
def describe_pet(animal_type, pet_name):
    print(f"\nMy {animal_type}'s name is {pet_name.title()}.")

describe_pet('hamster', 'harry')
describe_pet('dog', 'willie')
```

    
    My hamster's name is Harry.
    
    My dog's name is Willie.


- Keyword arguments


```python
# order does Not matter in Keyword Arguments 
# use the exact names of the parameters in the function’s definition
def describe_pet(animal_type, pet_name):
    print(f"\nI have a {animal_type}.")
    print(f"My {animal_type}'s name is {pet_name.title()}.")

describe_pet(animal_type='hamster', pet_name='harry')
describe_pet(pet_name='harry', animal_type='hamster') # same result
```

    
    I have a hamster.
    My hamster's name is Harry.
    
    I have a hamster.
    My hamster's name is Harry.


- Default Values


```python
# When writing a function, you can define a default value for each parameter.  
# When you use default values, any parameter with a default value needs to be listed 
# after all the parameters that don’t have default values. 
# This allows Python to continue interpreting positional arguments correctly.
def describe_pet(pet_name, animal_type='dog'):
    print(f"\nI have a {animal_type}.")
    print(f"My {animal_type}'s name is {pet_name.title()}.")

describe_pet(pet_name = 'willie')
describe_pet('willie') # simpler way 
describe_pet(pet_name='harry', animal_type='hamster') # an animal other than a dog
```

    
    I have a dog.
    My dog's name is Willie.
    
    I have a dog.
    My dog's name is Willie.
    
    I have a hamster.
    My hamster's name is Harry.


- Equivalent Function Calls


```python
def describe_pet(pet_name, animal_type='dog'):
    print("\nI have a " + animal_type + ".")
    print("My " + animal_type + "'s name is " + pet_name.title() + ".")

# A dog named Willie.
describe_pet('willie')
describe_pet(pet_name='willie')

# Keyword argument : order does not matter
# python knows where each value should go
describe_pet('harry', 'hamster')
describe_pet(pet_name='harry', animal_type='hamster')
describe_pet(animal_type='hamster', pet_name='harry')
```

    
    I have a dog.
    My dog's name is Willie.
    
    I have a dog.
    My dog's name is Willie.
    
    I have a hamster.
    My hamster's name is Harry.
    
    I have a hamster.
    My hamster's name is Harry.
    
    I have a hamster.
    My hamster's name is Harry.


- Avoiding Argument Errors


```python
def describe_pet(pet_name, animal_type='dog'):
    print("\nI have a " + animal_type + ".")
    print("My " + animal_type + "'s name is " + pet_name.title() + ".")
    
# ERROR! pet_name argument value was not provided
describe_pet()
describe_pet(animal_type='cat')
```


    ---------------------------------------------------------------------------

    TypeError                                 Traceback (most recent call last)

    <ipython-input-52-647671008042> in <module>
          4 
          5 # ERROR! pet_name argument value was not provided
    ----> 6 describe_pet()
          7 describe_pet(animal_type='cat')


    TypeError: describe_pet() missing 1 required positional argument: 'pet_name'


## Return Values


```python
def get_formatted_name(first_name, last_name):	
    full_name = first_name + ' ' + last_name
    return full_name.title()

musician = get_formatted_name('jimi', 'hendrix') # Jimi Hendrix
```

- Making an Argument Optional


```python
def get_formatted_name(first_name, last_name, middle_name=''):
    """Return a full name, neatly formatted."""
    if middle_name:
        full_name = first_name + ' ' + middle_name + ' ' + last_name
    else:
        full_name = first_name + ' ' + last_name
    return full_name.title()

musician = get_formatted_name('jimi', 'hendrix')
print(musician)
musician = get_formatted_name('john', 'hooker', 'lee')
print(musician)
```

    Jimi Hendrix
    John Lee Hooker


- Returning a Dictionary


```python
# takes in textual information and returns a meaningful data structure
def build_person(first_name, last_name, age=''):
    """Return a dictionary of information about a person."""
    person = {'first': first_name, 'last': last_name}
    if age:
        person['age'] = age
    return person

musician = build_person('jimi', 'hendrix', age=27)
# {'first': 'jimi', 'last': 'hendrix', 'age': 27}
print(musician)
```

    {'first': 'jimi', 'last': 'hendrix', 'age': 27}


## Using a Function with a while Loop


```python
def get_formatted_name(first_name, last_name):
    """Return a full name, neatly formatted."""
    full_name = first_name + ' ' + last_name
    return full_name.title()

while True:
    print("\nPlease tell me your name:")
    print("(enter 'q' at any time to quit)")
    f_name = input("First name: ")
    if f_name == 'q':
        break
    l_name = input("Last name: ")
    if l_name == 'q':
        break
    formatted_name = get_formatted_name(f_name, l_name)
    print("\nHello, " + formatted_name + "!")
```

    
    Please tell me your name:
    (enter 'q' at any time to quit)
    First name: Lulu
    Last name: JJ
    
    Hello, Lulu Jj!
    
    Please tell me your name:
    (enter 'q' at any time to quit)
    First name: q


## Passing a List


```python
def greet_users(names):
    for name in names:
        msg = "Hello, " + name.title() + "!"
        print(msg)

user_names = ['hannah', 'ty', 'margot']
greet_users(user_names)
```

    Hello, Hannah!
    Hello, Ty!
    Hello, Margot!


## Modifying a List in a Function


```python
unprinted_designs = ['iphone case', 'robot pendant', 'dodecahedron']
completed_models = []

while unprinted_designs:
    current_design = unprinted_designs.pop()
    print("Printing model: " + current_design)
    completed_models.append(current_design)

print("\nThe following models have been printed:")
for completed_model in completed_models:
    print(completed_model)

    
# reorganize the above code
def print_models(unprinted_designs, completed_models):
    while unprinted_designs:
        current_design = unprinted_designs.pop()
        print("Printing model: " + current_design)
        completed_models.append(current_design)

def show_completed_models(completed_models):
    print("\nThe following models have been printed:")
    for completed_model in completed_models:
        print(completed_model)

unprinted_designs = ['iphone case', 'robot pendant', 'dodecahedron']
completed_models = []
print_models(unprinted_designs, completed_models)
show_completed_models(completed_models)
```

    Printing model: dodecahedron
    Printing model: robot pendant
    Printing model: iphone case
    
    The following models have been printed:
    dodecahedron
    robot pendant
    iphone case
    Printing model: dodecahedron
    Printing model: robot pendant
    Printing model: iphone case
    
    The following models have been printed:
    dodecahedron
    robot pendant
    iphone case


## Preventing a Function from Modifying a List


```python
def print_models(unprinted_designs, completed_models):
    while unprinted_designs:
        current_design = unprinted_designs.pop()
        print("Printing model: " + current_design)
        completed_models.append(current_design)

def show_completed_models(completed_models):
    print("\nCompleted printing models :")
    for completed_model in completed_models:
        print(f"\t{completed_model}")
        
unprinted_designs = ['iphone case', 'robot pendant', 'dodecahedron']
completed_models = []

# pass the function a copy of the list, not the original. 
# The copied list will be modified, leaving the original intact.
print_models(unprinted_designs[:], completed_models)
print(f"\nUnprinted designs (unchanged) :\n\t{unprinted_designs}")
show_completed_models(completed_models)
```

    Printing model: dodecahedron
    Printing model: robot pendant
    Printing model: iphone case
    
    Unprinted designs (unchanged) :
    	['iphone case', 'robot pendant', 'dodecahedron']
    
    Completed printing models :
    	dodecahedron
    	robot pendant
    	iphone case


## Passing an Arbitrary Number of Arguments


```python
# make an empty tuple called toppings 
# and pack whatever values it receives into this tuple.
def make_pizza(*toppings):
    print(toppings)
make_pizza('pepperoni')
make_pizza('mushrooms', 'green peppers', 'extra cheese')
```

    ('pepperoni',)
    ('mushrooms', 'green peppers', 'extra cheese')


### Mixing Positional and Arbitrary Arguments


```python
def make_pizza(size, *toppings):
    """Summarize the pizza we are about to make."""
    print("\nMaking a " + str(size) +
          "-inch pizza with the following toppings:")
    for topping in toppings:
        print("- " + topping)

make_pizza(16, 'pepperoni')
make_pizza(12, 'mushrooms', 'green peppers', 'extra cheese')
```

    
    Making a 16-inch pizza with the following toppings:
    - pepperoni
    
    Making a 12-inch pizza with the following toppings:
    - mushrooms
    - green peppers
    - extra cheese


### Using Arbitrary Keyword Arguments


```python
# to accept an arbitrary number of 'keyword' arguments 
# as many key-value pairs as the calling statement provides
# **user_info cause Python to create an empty dictionary, user_info 
# and pack whatever name-value pairs it receives into this dictionary.
def build_profile(first, last, **user_info):
    """Build a dictionary containing everything we know about a user."""
    profile = {}
    profile['first_name'] = first
    profile['last_name'] = last
    for key, value in user_info.items():
        profile[key] = value
    return profile

user_profile = build_profile('albert', 'einstein', 
    location='princeton', 
    field='physics')
print(user_profile)
```

    {'first_name': 'albert', 'last_name': 'einstein', 'location': 'princeton', 'field': 'physics'}


## Storing Your Functions in Modules


```python
# store your functions in a separate file called a module(.py) 
# and then import that module into your main program. 
# pizza.py
def make_pizza(size, *toppings):
    print("\nMaking a " + str(size) + 
        "-inch pizza with the following toppings:")
    for topping in toppings:
        print("- " + topping)
```

### Importing an Entire Module


```python
# making_pizzas.py
import pizza

# module_name.function_name()
pizza.make_pizza(16, 'pepperoni')
pizza.make_pizza(12, 'mushrooms', 'green peppers', 'extra cheese')

# import module makes every function from the module available in your program 
```


    ---------------------------------------------------------------------------

    ModuleNotFoundError                       Traceback (most recent call last)

    <ipython-input-64-3a4bbe027e3d> in <module>
          1 # making_pizzas.py
    ----> 2 import pizza
          3 
          4 # module_name.function_name()
          5 pizza.make_pizza(16, 'pepperoni')


    ModuleNotFoundError: No module named 'pizza'


### Importing Specific Functions


```python
# With this syntax, you don’t need to use the dot notation when you call a function.
# from module_name import function_name
# from module_name import function_0, function_1, function_2
from pizza import make_pizza
make_pizza(16, 'pepperoni')
make_pizza(12, 'mushrooms', 'green peppers', 'extra cheese')
```


    ---------------------------------------------------------------------------

    ModuleNotFoundError                       Traceback (most recent call last)

    <ipython-input-65-e474112fee9a> in <module>
          2 # from module_name import function_name
          3 # from module_name import function_0, function_1, function_2
    ----> 4 from pizza import make_pizza
          5 make_pizza(16, 'pepperoni')
          6 make_pizza(12, 'mushrooms', 'green peppers', 'extra cheese')


    ModuleNotFoundError: No module named 'pizza'


### Using as to Give a Function an Alias


```python
# use a short, unique alias—an alternate name for the function.
from pizza import make_pizza as mp
mp(16, 'pepperoni')
mp(12, 'mushrooms', 'green peppers', 'extra cheese')
```

### Using as to Give a Module an Alias


```python
# alias for a module name
# import module_name as mn
import pizza as p
p.make_pizza(16, 'pepperoni')
p.make_pizza(12, 'mushrooms', 'green peppers', 'extra cheese')
```


    ---------------------------------------------------------------------------

    ModuleNotFoundError                       Traceback (most recent call last)

    <ipython-input-66-b2baca6bd01d> in <module>
          1 # alias for a module name
          2 # import module_name as mn
    ----> 3 import pizza as p
          4 p.make_pizza(16, 'pepperoni')
          5 p.make_pizza(12, 'mushrooms', 'green peppers', 'extra cheese')


    ModuleNotFoundError: No module named 'pizza'


### Importing All Functions in a Module


```python
# this approach is Not recommended
# if the module has a function name that matches an existing name 
# in your project, you can get some unexpected results.
# from module_name import *
from pizza import *
make_pizza(16, 'pepperoni')
make_pizza(12, 'mushrooms', 'green peppers', 'extra cheese')
```


    ---------------------------------------------------------------------------

    ModuleNotFoundError                       Traceback (most recent call last)

    <ipython-input-68-f8454aea797d> in <module>
          3 # in your project, you can get some unexpected results.
          4 # from module_name import *
    ----> 5 from pizza import *
          6 make_pizza(16, 'pepperoni')
          7 make_pizza(12, 'mushrooms', 'green peppers', 'extra cheese')


    ModuleNotFoundError: No module named 'pizza'


### Styling Functions

- descriptive function name with only lowercase and underscore
- comment should appear immediately after the function definition and use the docstring format.
- default value for a parameter: no spaces should be used on either side of the equal sign. e.g. param='default value'
- also same convention for arguments in function call. e.g. func(val_0, param='value')


```python
def function_name(
        parameter_0, parameter_1, parameter_2,
        parameter_3, parameter_4, parameter_5):
    function_body
```

- in a module, separate functions by 2 blank lines
- import statements should be at the beginning of a file
- exceptions are comments at the beginning to describe the program

# Classes
- create objects with unique traits
- instantiate; make an object from a class
- actions and information for an object
- write classes that extends the functionality of existing classes
- store classes in modules and import classes written by other programmers into your own program files

## Creating and Using a Class


```python
# __init__() runs automatically whenever we create a new instance
# 'self' parameter must come first before the other parameters
# it gives the individual instance access to the attributes and methods
# variables accessible through 'self' are called attributes (name, age)
class Dog():
    """A simple attempt to model a dog."""
    
    def __init__(self, name, age):
        """Initialize name and age attributes."""
        self.name = name
        self.age = age
    
    def sit(self):
        """Simulate a dog sitting in response to a command."""
        print(self.name.title() + " is now sitting.")
    
    def roll_over(self):
        """Simulate rolling over in response to a command."""
        print(self.name.title() + " rolled over!")
```

### python 2.7 class definition


```python
# put object inside parentheses
class ClassName(object):
```


      File "<ipython-input-70-f4f8fa6d9474>", line 2
        class ClassName(object):
                                ^
    SyntaxError: unexpected EOF while parsing



### Making an Instance from a Class


```python
class Dog():
    def __init__(self, name, age):
        self.name = name
        self.age = age
    
    def sit(self):
        print(self.name.title() + " is now sitting.")
    
    def roll_over(self):
        print(self.name.title() + " rolled over!")

my_dog = Dog('willie', 6)

# Accessing attributes
print("My dog's name is " + my_dog.name.title() + ".")
print("My dog is " + str(my_dog.age) + " years old.")

# Calling Methods
my_dog.sit()
my_dog.roll_over()

# Creating Multiple Instances
your_dog = Dog('lucy', 3)
print("\nYour dog's name is " + your_dog.name.title() + ".")
print("Your dog is " + str(your_dog.age) + " years old.")
your_dog.sit()
```

    My dog's name is Willie.
    My dog is 6 years old.
    Willie is now sitting.
    Willie rolled over!
    
    Your dog's name is Lucy.
    Your dog is 3 years old.
    Lucy is now sitting.


## Working with Classes and Instances


```python
# car.py
class Car():
    """A simple attempt to represent a car."""

    def __init__(self, make, model, year):
        """Initialize attributes to describe a car."""
        self.make = make
        self.model = model
        self.year = year
        # Setting a Default Value for an Attribute
        self.odometer_reading = 0

    def get_descriptive_name(self):
        """Return a neatly formatted descriptive name."""
        long_name = str(self.year) + ' ' + self.make + ' ' + self.model
        return long_name.title()

    def read_odometer(self):
        """Print a statement showing the car's mileage."""
        print("This car has " + str(self.odometer_reading) + " miles.")

my_new_car = Car('audi', 'a4', 2016)
print(my_new_car.get_descriptive_name())
my_new_car.read_odometer()
```

    2016 Audi A4
    This car has 0 miles.


### Modifying Attribute Values


```python
# 1. Modifying an attribute’s Value Directly
from car import Car

my_new_car = Car('audi', 'a4', 2017)
my_new_car.odometer_reading = 23 # modify 1

# 2. Modifying an attribute’s Value through a Method
class Car():
    def __init__(self, make, model, year):
        """Initialize attributes to describe a car."""
        self.make = make
        self.model = model
        self.year = year
        # Setting a Default Value for an Attribute
        self.odometer_reading = 0
    
    def update_odometer(self, mileage):
        """
        Set the odometer reading to the given value.
        Reject the change if it attempts to roll the odometer back.
        """
        if mileage >= self.odometer_reading:
            self.odometer_reading = mileage
        else:
            print("You can't roll back an odometer!")

my_new_car = Car('audi', 'a4', 2017)
my_new_car.update_odometer(23) # modify 2

# 3. Incrementing an attribute’s Value through a Method
class Car():
    def __init__(self, make, model, year):
        """Initialize attributes to describe a car."""
        self.make = make
        self.model = model
        self.year = year
        # Setting a Default Value for an Attribute
        self.odometer_reading = 0
    
    def increment_odometer(self, miles):
        """Add the given amount to the odometer reading."""
        self.odometer_reading += miles

my_used_car = Car('subaru', 'outback', 2013)
my_used_car.increment_odometer(100) # modify 3
```


    ---------------------------------------------------------------------------

    ModuleNotFoundError                       Traceback (most recent call last)

    <ipython-input-84-aab97358b19f> in <module>
          1 # 1. Modifying an attribute’s Value Directly
    ----> 2 from car import Car
          3 
          4 my_new_car = Car('audi', 'a4', 2017)
          5 my_new_car.odometer_reading = 23 # modify 1


    ModuleNotFoundError: No module named 'car'


## Inheritance
When one class inherits from another, it automatically takes on all the attributes and methods of the first class. The child class inherits every attribute and method from its parent class but is also free to define new attributes and methods of its own.


```python
from car import Car

class ElectricCar(Car):
    """Represent aspects of a car, specific to electric vehicles."""

    def __init__(self, make, model, year):
        """Initialize attributes of the parent class."""
        super().__init__(make, model, year)

my_tesla = ElectricCar('tesla', 'model s', 2016)
print(my_tesla.get_descriptive_name())
```


    ---------------------------------------------------------------------------

    ModuleNotFoundError                       Traceback (most recent call last)

    <ipython-input-85-0825b9fb3176> in <module>
    ----> 1 from car import Car
          2 
          3 class ElectricCar(Car):
          4     """Represent aspects of a car, specific to electric vehicles."""
          5 


    ModuleNotFoundError: No module named 'car'


### Inheritance in Python 2.7
super() requires a reference to the child class


```python
# super() needs two arguments: a reference to the child class and the self object
from car import Car

class ElectricCar(Car):
    def __init__(self, make, model, year):
        """Initialize attributes of the parent class."""
        super(ElectricCar, self).__init__(make, model, year)
```


    ---------------------------------------------------------------------------

    ModuleNotFoundError                       Traceback (most recent call last)

    <ipython-input-86-140dca9af6d5> in <module>
          1 # super() needs two arguments: a reference to the child class and the self object
    ----> 2 from car import Car
          3 
          4 class ElectricCar(Car):
          5     def __init__(self, make, model, year):


    ModuleNotFoundError: No module named 'car'


### Defining Attributes and Methods for the Child Class


```python
from car import Car

class ElectricCar(Car):
    def __init__(self, make, model, year):
        super().__init__(make, model, year)
        self.battery_size = 70
    
    def describe_battery(self):
        print('battery: ' + str(self.battery_size) + '-kWh')

my_tesla = ElectricCar('tesla', 'model s', 2016)
print(my_tesla.get_descriptive_name())
my_tesla.describe_battery()
```


    ---------------------------------------------------------------------------

    ModuleNotFoundError                       Traceback (most recent call last)

    <ipython-input-87-beefc43b7777> in <module>
    ----> 1 from car import Car
          2 
          3 class ElectricCar(Car):
          4     def __init__(self, make, model, year):
          5         super().__init__(make, model, year)


    ModuleNotFoundError: No module named 'car'


### Overriding Methods from the Parent Class


```python
from car import Car

class ElectricCar(Car):
    """say fill_gas_tank() is defined both in Car and ElectricCar classes 
    Python will ignore fill_gas_tank() in Car and run this code instead"""
    def fill_gas_tank(self):
        print('Electric Car does not need Gas!')
```

### Instances as Attributes


```python
from car import Car

class Battery():
    def __init__(self, battery_size=70):
        self.battery_size = battery_size
    
    def describe_battery(self):
        print('battery: ' + str(self.battery_size) + '-kWh')

    def get_range(self):
        """Print a statement about the range this battery provides."""
        if self.battery_size == 70:
            range = 240
        elif self.battery_size == 85:
            range = 270
        
        message = "This car can go approximately " + str(range)
        message += " miles on a full charge."
        print(message)

class ElectricCar(Car):
    def __init__(self, make, model, year):
        super().__init__(make, model, year)
        # battery instance as an attribute!
        self.battery = Battery()

my_tesla = ElectricCar('tesla', 'model s', 2016)
print(my_tesla.get_descriptive_name())

my_tesla.battery.describe_battery()
my_tesla.battery.get_range()

my_tesla.battery.upgrade_battery()
my_tesla.battery.get_range()
```

### Modeling Real-World Objects
- As you begin to model more complicated items like electric cars, you’ll wrestle with interesting questions. Is the range of an electric car a property of the battery or of the car?

## Importing Classes


```python
# 1. Importing a Single Class
# store classes (e.g. Car, User) in modules (e.g car.py, user.py) 
# and then import the classes you need into your main program.
# any program that uses this module will need a more specific filename,
# my_car.py
from car import Car

my_new_car = Car('audi', 'a4', 2016)
print(my_new_car.get_descriptive_name())


# 2. Storing Multiple Classes in a Module
# car.py
class Car(): # each class in a module should be related somehow.
class Battery():
class ElectricCar():

# my_electric_car.py
from car import ElectricCar

my_tesla = ElectricCar('tesla', 'model s', 2016)
print(my_tesla.get_descriptive_name())

my_tesla.battery.describe_battery()
my_tesla.battery.get_range()


# 3. Importing Multiple Classes from a Module
# my_cars.py
from car import Car, ElectricCar

my_beetle = Car('volkswagen', 'beetle', 2016)
print(my_beetle.get_descriptive_name())

my_tesla = ElectricCar('tesla', 'roadster', 2016)
print(my_tesla.get_descriptive_name())


# 4. Importing an Entire Module
import car

my_beetle = car.Car('volkswagen', 'beetle', 2016)
print(my_beetle.get_descriptive_name())
my_tesla = car.ElectricCar('tesla', 'roadster', 2016)
print(my_tesla.get_descriptive_name())


# 5. Importing All Classes from a Module
# This method is not recommended for two reasons. First, it’s helpful
# to be able to read the import statements at the top of a file and get a clear
# sense of which classes a program uses. Secondly, this approach can lead to confusion 
# with names in the file. If you accidentally import a class with the same name as 
# something else in your program file, you can create errors that are hard to diagnose. 
# from module_name import *
from car import *


# 6. Importing a Module into a Module
# When you store your classes in several modules, you may find that a class 
# in one module depends on a class in another module. When this happens, 
# you can import the required class into the first module.

# car.py
class Car():

    
# electric_car.py
from car import Car

class Battery():
class ElectricCar(Car):


# my_cars.py
from car import Car
from electric_car import ElectricCar

my_beetle = Car('volkswagen', 'beetle', 2016)
print(my_beetle.get_descriptive_name())

my_tesla = ElectricCar('tesla', 'roadster', 2016)
print(my_tesla.get_descriptive_name())
```

### Finding Your Own Workflow


```python
# When you’re starting out, keep your code structure simple.
# Try doing everything in one file and moving your classes 
# to separate modules once everything is working.
```

### The Python Standard Library


```python
# dictionary that keeps track of the order in which key-value pairs are 
# added, you can use the OrderedDict class from the collections module.
from collections import OrderedDict

# OrderedDict() creates an empty ordered dictionary
# OrderedDict objects behave like dictionaries except it is ordered
favorite_languages = OrderedDict()

favorite_languages['jen'] = 'python'
favorite_languages['sarah'] = 'c'
favorite_languages['edward'] = 'ruby'
favorite_languages['phil'] = 'python'

# prints in order!
for name, language in favorite_languages.items():
    print(name.title() + "'s favorite language : " + language.title())
```


```python
# die.py
from random import randint

class Die():
    def __init__(self, sides):
        self.sides = sides

    def roll_die(self):
        # random number between 1 and the number of sides
        x = randint(1, self.sides)
        print(str(self.sides) + '-side die result = ' + str(x))

die_6 = Die(6)
die_10 = Die(10)
die_20 = Die(20)

for i in range(10):
    die_6.roll_die()
    die_10.roll_die()
    die_20.roll_die()
    print()
```


```python
from random import choice
players = ['charlse', 'martina', 'michael', 'florence', 'eli']
first_up = choice(players)
print(first_up)

print('---10 random choices---')
for i in range(10):
    print(choice(players))
```

Go to [python module of the week](http://pymotw.com/) and look at the table of contents . Find a module that looks interesting to you and read about it, or explore the documentation of the collections and random module

Styling Classes
- Class names should be written in CamelCaps with no underscores
- Instance and module names should be written in lowercase and underscores
- class/function should have a docstring immediately following the definition 
- module should also have a docstring describing the classes in a module 
- Within a class, one blank line between methods
- Within a module, two blank lines to separate classes
- Place the import statement for the standard library module first. 
- Then add a blank line and the import statement for the module you wrote. 

# Files and Exceptions

## Reading from a File

Reading an Entire File


```python
# 'with' lets Python open and close the file properly
# 'with' keyword closes the file once access to it is no longer needed
# read() reads the entire contents of the file
# read() returns an empty string when it reaches the end of the file
# open(): returns a file object and its contents is stored in the variable
with open('pi_digits.txt') as file_object:
    contents = file_object.read() # contains an empty string at the end
    print(contents) # additional empty line at the end


# relative path
with open('text_files/filename.txt') as file_object:
    contents = file_object.read()
    print(contents)

# absolute path
with open('/home/user/workspace/python/text_files/filename.txt') as file_object:
    contents = file_object.read()
    print(contents)
```


```python
# Reading Line by Line
filename = 'pi_digits.txt'
with open(filename) as file_object:
    for line in file_object:
        print(line) # two newlines: from end of each line and 'print'
        # print(line.rstrip()) # avoid two new lines 

# Making a List of Lines from a File
filename = "pi_digits.txt"
with open(filename) as file_object:
    lines = file_object.readlines()

for line in lines:
    print(line.rstrip())


# Working with a File’s Contents
filename = 'pi_digits.txt'
with open(filename) as file_object:
    lines = file_object.readlines()

pi_string = ''
for line in lines:
    pi_string += line.strip()

print(pi_string)
print(len(pi_string))


# Large Files: One Million Digits
filename = 'pi_million_digits.txt'
with open(filename) as file_object:
    lines = file_object.readlines()

pi_string = ''
for line in lines:
    pi_string += line.strip()

print(pi_string[:52] + "...") # 3.1415926535897932384626433832795028...
print(len(pi_string)) # 1000002

birthday = input('Enter your birthday, in the form mmddyy: ')
if birthday in pi_string:
    print('Your birthday appears in the first million digits of pi!')
else:
    print('Your birthday does not appear in the first million digits.')
```

## Writing to a File


```python
# Writing to an Empty File

filename = 'programming.txt'

# read('r'), write('w'), append('a') mode
with open('filename', 'w') as file_object:
    file_object.write('I love programming.\n')
    file_object.write('I love python language')

# only strings can be written to a text file.
# use str() for storing numerical data

# Appending to a File
filename = 'programming.txt'
with open(filename, 'a') as file_object: #append
    file_object.write("I love finding meaning in large data sets.\n")
    file_object.write("I love creating apps.\n")


# Try it yourself
filename = 'guest_book.txt'
flag = True
while flag:
    name = input("Type your name ('q' to quit) >> ")
    if not name:
        print("Empty name is not allowed, please type again...\n")
        continue
    elif name != 'q':
        print('\tHi, ' + name + '. Welcome back!')
        with open(filename, 'a') as f_obj:
            f_obj.write(name + '\n')
    else:
        print("Exiting the program...")
        break

with open(filename) as f_obj:
    for line in f_obj:
        print(line.strip())
```

## Exceptions
Whenever an error occurs during a program's execution, exceptions objects are created to manage errors. If you don’t handle the exception, the program will halt and show a traceback, which includes a report of the exception that was raised. Exceptions are handled with try-except blocks.


```python
# ZeroDivisionError
try:
    print(5/0)
except ZeroDivisionError:
    print('You cannot divide by zero!')


# if try block is executed successfully, else block is run
print("Give me two numbers, and I'll divide them.")
print("Enter 'q' to quit.")
while True:
    first_number = input("\nFirst number: ")
    if first_number == 'q':
        break
    second_number = input("Second number: ")
    if second_number == 'q':
        break
    try:
        answer = int(first_number) / int(second_number)
    except ZeroDivisionError:
        print("You can't divide by 0!")
    else:
        print(answer)


# FileNotFoundError
def count_words(filename):
    """Count the approximate number of words in a file."""
    try:
        with open(filename) as f_obj:
            contents = f_obj.read()
    except FileNotFoundError:
        msg = "Sorry, the file '" + filename + "' does not exist."
        print(msg)
    else:
        # count the approximate number of words in the file
        words = contents.split()
        num_words = len(words)
        print("'" + filename + "' has " + str(num_words) + " words")

filenames = ['alice.txt', 'moby_dick.txt', 'little_women.txt']
for filename in filenames:
    count_words(filename)


# Failing Silently
# Deciding Which Errors to Report
# you don’t need to report every exception you catch.
def count_words(filename):
    """Count the approximate number of words in a file."""
    try:
        # --snip--
    except FileNotFoundError:
        pass
    else:
        # --snip--

# Try it yourself
def add_numbers():
    while True:
        num1 = input("integer num1 ('q' to quit) > ")
        if num1 == 'q':
            break
        num2 = input("integer num2 ('q' to quit) > ")
        if num2 == 'q':
            break
        try:
            division = int(num1) / int(num2)
        except ValueError:
            print('One of your inputs are not divisible; ValueError!\n')
        except ZeroDivisionError:
            print('Cannot be divided by 0\n')
        else:
            print("num1 / num2 = " + str(division) + '\n')

add_numbers()

# Count words
line = "Row, row, row your boat."
print(line.count('row')) # 2
print(line.lower().count('row')) # 3
```

## Storing Data
The json module allows you to dump simple Python data structures into a file and load the data from that file the next time the program runs. You can also use json to share data between different Python programs. JSON(JavaScript Object Notation) data format is not specific to Python, so you can share data you store in the JSON format with people who work in many other programming languages.


```python
# Using json.dump() and json.load()
# number_writer.py
import json

numbers = [2, 3, 5, 7, 11, 13]
filename = 'numbers.json'
with open(filename, 'w') as f_obj:
    json.dump(numbers, f_obj) # stores a piece of data to a file object
    

# number_reader.py
import json

filename = 'numbers.json'
with open(filename) as f_obj:
    numbers = json.load(f_obj) # reads numbers back into memory

print(numbers) # [2, 3, 5, 7, 11, 13]


# Saving and Reading User-Generated Data
import json
# Load the username, if it has been stored previously.
# Otherwise, prompt for the username and store it.
filename = 'username.json'
try:
    with open(filename) as f_obj:
        username = json.load(f_obj)
except FileNotFoundError:
    username = input("What is your name? ")
    with open(filename, 'w') as f_obj:
        json.dump(username, f_obj)
        print("We'll remember you when you come back, " + username + "!")
else:
    print("Welcome back, " + username + "!")
    

# Refactoring
""" Often, you’ll come to a point where your code will work, but you’ll recog-
nize that you could improve the code by breaking it up into a series of func-
tions that have specific jobs. This process is called refactoring. Refactoring
makes your code cleaner, easier to understand, and easier to extend.
We can refactor remember_me.py by moving the bulk of its logic into one
or more functions."""
import json

def get_stored_username():
    """Get stored username if available."""
    filename = 'username.json'
    try:
        with open(filename) as f_obj:
            username = json.load(f_obj)
    except FileNotFoundError:
        return None
    else:
        return username

def get_new_username():
    """Prompt for a new username."""
    filename = 'username.json'
    username = input("What is your name? ")
    with open(filename, 'w') as f_obj:
        json.dump(username, f_obj)
    return username

def greet_user():
    """ Greet the user by name """
    username = get_stored_username()
    if verify_user(username):
        print("Welcome back, " + username + "!")
    else:
        username = get_new_username()
        print("We'll remember you, " + username + "!")

def verify_user(username):
    if username:
        answer = input("Is username, '" + username + "' correct?(y/n)")
        if answer.lower() == 'y':
            return True
        else:
            return False
    else:
        return False

greet_user()
```

# Testing Your Code

## Testing a Function


```python
# name_function.py
def get_formatted_name(first, last):
    """ Generate a neatly formatted full name. """
    full_name = first + ' ' + last
    return full_name
```


```python
# names.py
from name_function import get_formatted_name

while True:
    print("Enter 'q' at any time to quit.")
    first = input("Your first name >> ")
    if first == 'q':
        break
    last = input("Your last name >> ")
    if last == 'q':
        break
    formatted_name = get_formatted_name(first, last)
    print("\tNeatly formatted name : " + formatted_name + ".")
```

### Unit Tests and Test Cases


```python
""" A unit test verifies that one specific aspect of a function’s
behavior is correct. A test case is a collection of unit tests that together prove
that a function behaves as it’s supposed to, within the full range of situations 
you expect it to handle """
# test_name_function.py
# Assert methods verify that a result you received matches the expected one
# The method name must start with test_ so the method runs automatically when we run test_name_function.py.
 
import unittest
from name_function import get_formatted_name

class NamesTestCase(unittest.TestCase):
    """Tests for 'name_function.py'"""
    
    def test_first_last_name(self):
        """Do names like 'Janis Joplin' work?"""
        formatted_name = get_formatted_name('janis', 'joplin')
        self.assertEqual(formatted_name, 'Janis Joplin')

unittest.main()
```

Check error messages after changing get_formatted_name implementation, or fix the code to avoid the error


```python
def get_formatted_name(first, last, middle=''):
    if middle:
        full_name = first + ' ' + middle + ' ' + last
    else:
        full_name = first + ' ' + last
    return full_name
```

Adding New Tests


```python
import unittest
from name_function import get_formatted_name

class NamesTestCase(unittest.TestCase):
    """Test for 'name_function.py'"""
    
    def test_first_last_name(self):
        """Do names like 'Janis Joplin' work?"""
        formatted_name = get_formatted_name('janis', 'joplin')
        self.assertEqual(formatted_name, 'Janis Joplin')
        
    def test_first_last_middle_name(self):
        """Do names like 'Wolfgang Amadeus Mozart' work?"""
        formatted_name = get_formatted_name('wolfgang', 'mozart', 'amadeus')
        self.assertEqual(formatted_name, 'Wolfgang Amadeus Mozart')

unittest.main()
```

Try it yourself


```python
# city_functions.py
def get_city_country(city, country, population=''):
    if population:
        location_info = city.title() + ', ' + country.title() \
                    + ' - population ' + str(population)
    else:
        location_info = city.title() + ', ' + country.title()
    return location_info


# test_city_functions.py
import unittest
from city_functions import get_city_country


class CityCountryTestCase(unittest.TestCase):
    def test_city_country(self):
        location = get_city_country('aaa bbb', 'ccc')
        self.assertEqual(location, "Aaa Bbb, Ccc")

    def test_city_country_population(self):
        location = get_city_country('aaa bbb', 'ccc', 1000)
        self.assertEqual(location, "Aaa Bbb, Ccc - population 1000")

unittest.main()
```

## Testing a Class
Python provides a number of assert methods in the 'unittest.TestCase' class. You can use these methods only in a class that inherits from 'unittest.TestCase'
```html
<table>
    <tr>
        <th>Method</th>
        <th>Use</th>
    </tr>
    <tr>
        <td>assertEqual(a, b)</td>
        <td>Verify that a == b</td>
    </tr>
    <tr>
        <td>assertNotEqual(a, b)</td>
        <td>Verify that a != b</td>
    </tr>
    <tr>
        <td>assertTrue(x)</td>
        <td>Verify that x is True</td>
    </tr>
    <tr>
        <td>assertFalse(x)</td>
        <td>Verify that x is False</td>
    </tr>
    <tr>
        <td>assertIn(item, list)</td>
        <td>Verify that item is in list</td>
    </tr>
    <tr>
        <td>assertNotIn(item, list)</td>
        <td>Verify that item is not in list</td>
    </tr>
</table>
```

A Class to Test


```python
# survey.py
class AnonymousSurvey():
    """Collect anonymous answers to a survey question."""

    def __init__(self, question):
        """Store a question, and prepare to store responses."""
        self.question = question
        self.responses = []

    def show_question(self):
        """Show the survey question."""
        print(self.question)

    def store_response(self, new_response):
        """Store a single response to the survey."""
        self.responses.append(new_response)

    def show_results(self):
        """Show all the responses that have been given."""
        print("Survey results:")
        for response in self.responses:
            print('- ' + response)
```


```python
# language_survey.py
from survey import AnonymousSurvey

# Define a survey, and make a survey
question = "What language did you first learn to speak?"
my_survey = AnonymousSurvey(question)

# Show the question, and store responses to the question
my_survey.show_question()
print("Enter 'q' at any time to quit.\n")
while True:
    response = input("Lanugage: ")
    if response == 'q':
        break
    my_survey.store_response(response)

# Show the survey results.
print("\nThank you to everyone who participated in the survey!")
my_survey.show_results()
```


```python
# test_survey.py
import unittest
from survey import AnonymousSurvey


class TestAnonymousSurvey(unittest.TestCase):
    """
    Tests for the class AnonymousSurvey
    The method test_store_single_response() verifies that the first response in
    self.responses—self.responses[0]—can be stored correctly, and test_store_
    single_response() verifies that all three responses in self.responses can be
    stored correctly.
    """

    def setUp(self):
        """
        Create a survey and a set of responses for use in all test methods.
        """
        question = "What language did you first learn to speak?"
        self.my_survey = AnonymousSurvey(question)
        self.responses = ['English', 'Spanish', 'Mandarin']

    def test_store_single_response(self):
        """Test that a single response is stored properly"""
        self.my_survey.store_response(self.responses[0])
        self.assertIn(self.responses[0], self.my_survey.responses)

    def test_store_three_responses(self):
        """Test that three individual responses are stored properly."""

        for response in self.responses:
            self.my_survey.store_response(response)

        for response in self.responses:
            self.assertIn(response, self.my_survey.responses)

unittest.main()
```

Try it yourself


```python
# employee.py
class AnonymousEmployee():
    def __init__(self, first, last, annual_salary):
        self.first = first
        self.last = last
        self.annual_salary = int(annual_salary)

    def give_raise(self, annual_raise=5000):
        if annual_raise:
            self.annual_salary += annual_raise
        else:
            self.annual_salary += 5000
```


```python
# test_employee.py
import unittest
from employee import AnonymousEmployee


class TestAnonymousEmployee(unittest.TestCase):
    def setUp(self):
        self.annual_salary = 80000
        self.my_employee = AnonymousEmployee('aaa', 'bbb',
                                             self.annual_salary)

    def test_give_default_raise(self):
        self.my_employee.give_raise()
        self.assertEqual(85000, self.my_employee.annual_salary)

    def test_give_custom_raise(self):
        self.my_employee.give_raise(8000)
        self.assertEqual(88000, self.my_employee.annual_salary)

unittest.main()
```

# Project 1. Alien Invasion

## install pip in Linux
```
sudo apt update
sudo apt upgrade
sudo apt dist-upgrade

sudo apt -y install python3-pip

# check version
pip3 -V
```

## download and install [pip](https://bootstrap.pypa.io/get-pip.py) in Windows
```
python -m pip --version
python get-pip.py
```

## install pygame Linux
If you’re running Python 3, two steps are required: installing the libraries Pygame depends on, and downloading and installing Pygame.

The following is for python 3.5
```
# libraries for pygame
sudo apt-get install python3.5-dev mercurial
sudo apt-get install libsdl-image1.2-dev libsdl2-dev libsdl-ttf2.0-dev

# more libraries
sudo apt-get install libsdl-mixer1.2-dev libportmidi-dev
sudo apt-get install libswscale-dev libsmpeg-dev libavformat-dev libavcodec-dev
sudo apt-get install python-numpy

pip3 install --user hg+http://bitbucket.org/pygame/pygame

# if previous commands give "freetype-config" error try the following
sudo apt-get install libfreetype6-dev
```

To install pygame on python 3.6, check [here](https://askubuntu.com/questions/401342/how-to-download-pygame-in-python3-3)
```
# 1. Change to your home directory.
cd ~

# 2. Get Pygame source code.
sudo apt-get install mercurial
hg clone https://bitbucket.org/pygame/pygame
cd pygame

# 3. Install dependencies.
sudo apt-get install python3-dev python3-numpy libsdl-dev libsdl-image1.2-dev \
  libsdl-mixer1.2-dev libsdl-ttf2.0-dev libsmpeg-dev libportmidi-dev \
  libavformat-dev libswscale-dev libjpeg-dev libfreetype6-dev

# 4. Build and install Pygame.
python3.6 setup.py build
sudo python3.6 setup.py install
```

## download and install [pygame](http://www.lfd.uci.edu/~gohlke/pythonlibs/#pygame) in Windows
```
python -m pip install --user pygame-1.9.3-cp36-cp36m-win32.whl
```


# Project 2. Data Visualization

## install matplotlib in Linux
```
sudo apt update
sudo apt install python3-matplotlib
sudo apt install python3.6-dev python3.6-tk tk-dev
sudo apt install libfreetype6-dev g++
pip3 install --user matplotlib
```

## download and install [matplotlib](http://www.lfd.uci.edu/~gohlke/pythonlibs/#matplotlib) in Windows
```
python -m pip install --user matplotlib-2.1.0-cp36-cp36m-win32.whl
```

mpl_squares.py
```python
import matplotlib.pyplot as plt

input_values = [1, 2, 3, 4, 5]
squares = [1, 4, 9, 16, 25]
plt.plot(input_values, squares, linewidth=5)

# Set chart title and label axes.
plt.title("Square Numbers", fontsize=24)
plt.xlabel("Value", fontsize=14)
plt.ylabel("Square of Value", fontsize=14)

# Set size of tick labels.
plt.tick_params(axis='both', labelsize=14)

plt.show()
```

scatter_squares.py
```python
import matplotlib.pyplot as plt

x_values = list(range(1, 1001))
y_values = [x**2 for x in x_values]

# plt.scatter(x_values, y_values, edgecolor='none', s=40)
# plt.scatter(x_values, y_values, c='red', edgecolor='none', s=40)
# plt.scatter(x_values, y_values, c=(0, 0, .8), edgecolor='none', s=40)
plt.scatter(x_values, y_values, c=y_values, cmap=plt.cm.Blues,
            edgecolor='none', s=40)

# Set chart title, and label axes.
plt.title("Square Numbers", fontsize=24)
plt.xlabel("Value", fontsize=14)
plt.ylabel("Square of Value", fontsize=14)

# Set size of tick labels.
plt.tick_params(axis='both', which='major', labelsize=14)

# Set the range for each axis.
plt.axis([0, 1100, 0, 1100000])

plt.savefig('squares_plot.png', bbox_inches='tight')

plt.show()
```

Random Walks
```python
# random_walk.py
from random import choice

class RandomWalk():
    """A class to generate random walks."""
    
    def __init__(self, num_points=5000):
        """Initialize attributes of a walk."""
        self.num_points = num_points
        
        # All walks start at (0, 0).
        self.x_values = [0]
        self.y_values = [0]

    def fill_walk(self):
        """Calculate all the points in the walk."""
        
        # Keep taking steps until the walk reaches the desired length.
        while len(self.x_values) < self.num_points:
            
            # Decide which direction to go, and how far to go in that direction.
            x_step = self.get_step()
            y_step = self.get_step()
            
            # Reject moves that go nowhere.
            if x_step == 0 and y_step == 0:
                continue
            
            # Calculate the next x and y values.
            next_x = self.x_values[-1] + x_step
            next_y = self.y_values[-1] + y_step
            
            self.x_values.append(next_x)
            self.y_values.append(next_y)
            
    def get_step(self):
        direction = choice([1, -1])
        distance = choice([0, 1, 2, 3, 4])
        return direction * distance
```

```python
# rw_visual.py
import matplotlib.pyplot as plt

from random_walk import RandomWalk

# Keep making new walks, as long as the program is active.
while True:
    # Make a random walk, and plot the points.
    rw = RandomWalk(50000)
    rw.fill_walk()
    
    # Set the size of the plotting window.
    plt.figure(dpi=128, figsize=(10, 6))
    
    point_numbers = list(range(rw.num_points))
    plt.scatter(rw.x_values, rw.y_values, c=point_numbers,
                cmap=plt.cm.Blues, edgecolor='none', s=1)
    # Emphasize the first and last points.
    plt.scatter(0, 0, c='green', edgecolors='none', s=100)
    plt.scatter(rw.x_values[-1], rw.y_values[-1], c='red',
                edgecolors='none', s=100)
    
    # Remove the axes.
    plt.axes().get_xaxis().set_visible(False)
    plt.axes().get_yaxis().set_visible(False)
    
    plt.show()
    
    keep_running = input("Make another walk? (y/n): ")
    if keep_running == 'n':
        break
```

## install pygal in Linux
```
pip3 install --user pygal
```

## install pygal in Windows
```
python -m pip install --user pygal
```

```python
# die.py
from random import randint

class Die():
    def __init__(self, num_sizes=6):
        self.num_sizes = num_sizes

    def roll(self):
        return randint(1, self.num_sizes)
```

```python
# different_dice.py
import pygal

from die import Die

die_1 = Die()
die_2 = Die(10)

results = []
for i in range(50000):
    result = die_1.roll() + die_2.roll()
    results.append(result)

frequencies = []
max_result = die_1.num_sides + die_2.num_sides
for value in range(2, max_result + 1):
    frequency = results.count(value)
    frequencies.append(frequency)

print(frequencies)

hist = pygal.Bar()

hist.title = "Results of rolling D6 and D10 50000 times."
hist.x_labels = ['2', '3', '4', '5', '6', '7', '8', '9', '10', '11',
                 '12', '13', '14', '15', '16']
hist.x_title = "Result"
hist.y_title = "Frequency of Result"

hist.add('D6 + D10', frequencies)
hist.render_to_file('dice_visual.svg')
```

## Downloading Data
You’ll download data sets from online sources and create working visualizations of that data.

We’ll use Python’s csv module to process data stored in the CSV format and analyze. Then use matplotlib to generate a chart based on downloaded data to display variations in temperature in two very different environments. Later, we’ll use the json module to access data stored in the JSON format and use Pygal to draw a population map by country.

At the end, you’ll be prepared to work with different types and formats of data sets, and you’ll have a deeper understanding of how to build complex visualizations. The ability to access and visualize online data of different types and formats is essential to working with a wide variety of real-world data sets.

### The CSV File Format
Python’s csv module in the standard library parses the lines in a CSV file and allows us to quickly extract the values we’re interested in. 
```python
# highs_lows.py

import csv
from datetime import datetime

from matplotlib import pyplot as plt

# filename = 'sitka_weather_2014.csv'
filename = 'death_valley_2014.csv'

try:
    with open(filename) as f:
        reader = csv.reader(f)
        header_row = next(reader)

        # for index, column_header in enumerate(header_row):
        #     print(index, column_header)

        # Get dates and high temperatures from file.
        dates, highs, lows = [], [], []
        for row in reader:
            try:
                current_date = datetime.strptime(row[0], "%Y-%m-%d")
                high = int(row[1])
                low = int(row[3])
            except ValueError:
                print(current_date, 'missing data')
            else:
                dates.append(current_date)
                highs.append(high)
                lows.append(low)

        highs = [int((x - 32) / 1.8) for x in highs]
        lows = [int((x - 32) / 1.8) for x in lows]

except FileNotFoundError:
    msg = "\nSorry, the file '" + filename + "' does not exist."
    print(msg)
else:
    print("\nFile '" + filename + "' was successfully processed...")

# Plot data using matplotlib.
fig = plt.figure(dpi=128, figsize=(10, 6))
plt.plot(dates, highs, c='red', alpha=.5)
plt.plot(dates, lows, c='blue', alpha=.5)
plt.fill_between(dates, highs, lows, facecolor='blue', alpha=.1)

# Format plot.
title = "Daily high and low temperatures - 2014\nDeath Valley, CA"
plt.title(title, fontsize=20)
plt.xlabel('', fontsize=12)
fig.autofmt_xdate()
plt.ylabel("Temperature (F)", fontsize=12)
plt.tick_params(axis='both', which='major', labelsize=12)

plt.show()
```


## The JSON File Format
```python
# i18n was removed recently
from pygal.i18n import COUNTRIES
```

install pygal_maps_world module
```
# linux
pip3 install --user pygal_maps_world
# windows
python -m pip install --user pygal_maps_world
```

```python
# world_population.py
import json
from country_codes import get_country_code

filename = 'population_data.json'
with open(filename) as f:
    pop_data = json.load(f)

for pop_dict in pop_data:
    if pop_dict['Year'] == '2010':
        country_name = pop_dict['Country Name']
        population = int(float(pop_dict['Value']))
        code = get_country_code(country_name)
        if code:
            print(code + ": " + str(population))
        else:
            print('ERROR - ' + country_name)
```

```python
# countries.py
# from pygal_maps_world import i18n
from pygal.maps.world import COUNTRIES

for country_code in sorted(COUNTRIES.keys()):
    print(country_code, COUNTRIES[country_code])
```

```python
# country_codes.py
# from pygal_maps_world import i18n
from pygal.maps.world import COUNTRIES


def get_country_code(country_name):
    """Return the Pygal 2-digit country code"""
    for code, name in COUNTRIES.items():
        if name == country_name:
            return code
    return None
```

Building a World Map
```python
# americas.py
from pygal.maps.world import World

wm = World()
wm.title = 'North, Central, and South America'

wm.add('North America', ['ca', 'mx', 'us'])
wm.add('Central America', ['bz', 'cr', 'gt', 'hn', 'ni', 'pa', 'sv'])
wm.add('South America', ['ar', 'bo', 'br', 'cl', 'co', 'ec', 'gf',
                         'gy', 'pe', 'py', 'sr', 'uy', 've'])
wm.render_to_file('americas.svg')
```

```python
# na_populations.py
from pygal.maps.world import World

wm = World()
wm.title = 'Populations of Countries in North America'
wm.add('North America',
       {'ca': 34126000,
        'us': 309349000,
        'mx': 113423000})
wm.render_to_file('na_populations.svg')
```

```python
# world_population.py
import json
from country_codes import get_country_code
from pygal.maps.world import World
from pygal.style import LightColorizedStyle, RotateStyle


filename = 'population_data.json'
with open(filename) as f:
    pop_data = json.load(f)

# Build a dictionary of population data.
cc_populations = {}

for pop_dict in pop_data:
    if pop_dict['Year'] == '2010':
        country_name = pop_dict['Country Name']
        population = int(float(pop_dict['Value']))
        code = get_country_code(country_name)
        if code:
            cc_populations[code] = population

# Group the countries into 3 population levels.
cc_pops_1, cc_pops_2, cc_pops_3 = {}, {}, {}
for cc, pop in cc_populations.items():
    if pop < 10000000:
        cc_pops_1[cc] = pop
    elif pop < 1000000000:
        cc_pops_2[cc] = pop
    else:
        cc_pops_3[cc] = pop

# See how many countries are in each level
print(len(cc_pops_1), len(cc_pops_2), len(cc_pops_3))


# wm_style = RotateStyle('#336699')
wm_style = RotateStyle('#336699', base_style=LightColorizedStyle)
wm = World(style=wm_style)
wm.title = 'World Population in 2010, by Country'
wm.add('0-10m', cc_pops_1)
wm.add('10m-1bn', cc_pops_2)
wm.add('>1bn', cc_pops_3)

wm.render_to_file('world_population.svg')
```

## Working with APIs
Using a Web API
```
# API call example : 
# https://api.github.com/ directs the request to the part of GitHub’s 
# website that responds to API calls. 
# search/repositories, tells the API to search through all repositories
# The question mark signals passing an argument. The q stands for query  
# language:python indicates that we want information only on 
# repositories that have Python as the primary language. 
# &sort=stars sorts the projects by the number of stars

https://api.github.com/search/repositories?q=language:python&sort=stars
```

## install requests
The requests package allows a Python program to easily request informa-
tion from a website and examine the response that’s returned.
```
pip3 install --user requests # linux
python -m pip install --user requests # windows
```

Processing an API Response
```python
# python_repos.py
import requests

# Make an API call and store the response.
url = 'https://api.github.com/search/repositories?q=language:python&sort=stars'
r = requests.get(url)
print("Status code:", r.status_code)

# Store API response in a variable.
response_dict = r.json()
print("Total repositories:", response_dict['total_count'])

repo_dicts = response_dict['items']
print("Repositories returned:", len(repo_dicts))

print("\nSelected information about each repository: ")
for repo_dict in repo_dicts:
    print("\nName:", repo_dict['name'])
    print("Owner:", repo_dict['owner']['login'])
    print("Stars:", repo_dict['stargazers_count'])
    print("Repository:", repo_dict['html_url'])
    print("Created:", repo_dict['created_at'])
    print("Updated:", repo_dict['updated_at'])
    print("Description:", repo_dict['description'])
```

## Monitoring API Rate Limits
Most APIs are rate-limited, which means there’s a limit to how many requests you can make in a certain amount of time. To see if you’re approaching GitHub’s limits, enter https://api.github.com/rate_limit into a web browser. You should see a response like this:
```
# The reset value represents the time in Unix or epoch time (the number 
# of seconds since midnight on January 1, 1970) when quota will reset
{
  "resources": {
    "core": {
      "limit": 60,
      "remaining": 60,
      "reset": 1508215491
    },
    "search": {
      "limit": 10,
      "remaining": 8,
      "reset": 1508211923
    },
    "graphql": {
      "limit": 0,
      "remaining": 0,
      "reset": 1508215491
    }
  },
  "rate": {
    "limit": 60,
    "remaining": 60,
    "reset": 1508215491
  }
}
```

## Visualizing Repositories Using Pygal
```python
# python_repos.py
import requests
import pygal
from pygal.style import LightColorizedStyle as LCS, LightenStyle as LS

# Make an API call and store the response.
url = 'https://api.github.com/search/repositories?q=language:python&sort=stars'
r = requests.get(url)
print("Status code:", r.status_code)

# Store API response in a variable.
response_dict = r.json()
print("Total repositories:", response_dict['total_count'])

repo_dicts = response_dict['items']

names, plot_dicts = [], []
for repo_dict in repo_dicts:
    names.append(repo_dict['name'])

    description = repo_dict['description']
    if not description:
        description = "No description provided."
    plot_dict = {
        'value': repo_dict['stargazers_count'],
        'label': description,
        'xlink': repo_dict['html_url'],
    }
    plot_dicts.append(plot_dict)

# Make visualization.
my_style = LS('#333366', base_style=LCS)

my_config = pygal.Config()
my_config.x_label_rotation = 45
my_config.show_legend = False
my_config.title_font_size = 24
my_config.label_font_size = 14
my_config.major_label_font_size = 18
my_config.truncate_label = 15
my_config.show_y_guides = False  # hide horizontal lines in the graph
my_config.width = 1000

# pygal.Bar(style=my_style, x_label_rotation=45, show_legend=False)
chart = pygal.Bar(my_config, style=my_style)
chart.title = 'Most-Starred Python Projects on GitHub'
chart.x_labels = names

chart.add('', plot_dicts)
chart.render_to_file('python_repos.svg')
```

## The Hacker News API
The following call returns information about the current top article on [Hacker's News](http://news.ycombinator.com/) as of this writing:
```
https://hacker-news.firebaseio.com/v0/item/9884165.json
```

The response is a dictionary of information about the article with the ID 9884165:
```
{
    'url': 'http://www.bbc.co.uk/news/science-environment-33524589',
    'type': 'story',
    'title': 'New Horizons: Nasa spacecraft speeds past Pluto',
    'descendants': 141,
    'score': 230,
    'time': 1436875181,
    'text': '',
    'by': 'nns',
    'id': 9884165,
    'kids': [9884723, 9885099, 9884789, 9885604, 9885844]
}
```

Let’s make an API call that returns the IDs of the current top articles on Hacker News, and then examine each of the top articles:
```python
# hn_submissions.py
import requests

from operator import itemgetter

# Make an API call and store the response.
url = 'https://hacker-news.firebaseio.com/v0/topstories.json'
r = requests.get(url)
print("Status code:", r.status_code)

# Process information about each submission.
submission_ids = r.json()
submission_dicts = []

for submission_id in submission_ids[:10]:
    # Make a separate API call for each submission.
    url = 'https://hacker-news.firebaseio.com/v0/item/' + \
          str(submission_id) + '.json'
    submission_r = requests.get(url)
    print(submission_r.status_code)
    response_dict = submission_r.json()

    submission_dict = {
        'title': response_dict['title'],
        'link': 'http://news.ycombinator.com/item?id=' +
                str(submission_id),
        'comments': response_dict.get('descendants', 0)
    }
    submission_dicts.append(submission_dict)

submission_dicts = sorted(submission_dicts, key=itemgetter('comments'),
                          reverse=True)

for submission_dict in submission_dicts:
    print("\nTitle:", submission_dict['title'])
    print("Discussion link:", submission_dict['link'])
    print("Comments:", submission_dict['comments'])
```

When you’re not sure if a key exists in a dictionary, use the dict.get() method, which returns the value associated with the given key if it exists or the value you provide if the key doesn’t exist (0 in this example).


# Project 3. Web Applications

## Getting started with Django
Django is a web framework—a set of tools designed to help you build interactive websites. Django can respond to page requests and make it easier to read and write to a database, manage users, and much more.

### Create virtual environment
A virtual environment is a place on your system where you can install packages and isolate them from all other Python packages.
```
mkdir learning_log && cd learning_log

# 1. Create w/ venv module
python3.6 -m venv ll_env

# 2. Create w/ virtualenv package. 
# python -m pip install --user virtualenv in Windows
# sudo apt install python3-virtualenv or
pip3 install --user virtualenv  # in Linux
virtualenv ll_env --python=python3.6
```

### Activating the Virtual Environment
Packages in ll_env will be available while the environment is activated
```
# ll_env\Scripts\activate in Windows
source ll_env/bin/activate  # Linux
deactivate
```

### Installing Django
```
# install inside 'learning_log' directory
pip3 install Django
```

### Creating a Project in Django
Tells django to set up a new project called 'learning_log'. The dot at the end of the command creates the new project with a directory structure that will make it easy to deploy the app to a server when we’re finished developing it.
```
django-admin.py startproject learning_log .

ls
    learning_log  ll_env  manage.py
ls learning_log
    __init__.py  settings.py  urls.py  wsgi.py
```

### Creating the Database
Because Django stores most of the information related to a project in a database, we need to create a database that Django can work with. Create the database for the Learning Log project.
```
python manage.py migrate

ls
    db.sqlite3 learning_log ll_env manage.py
```

### Viewing the Project
Django starts a server so you can view the project on your system to see how well it works.
```
python manage.py runserver

# the project is listening for requests on port 8000 on localhost
# enter http://localhost:8000/, or http://127.0.0.1:8000/ on a browser
# If you receive the error message That port is already in use,
# tell Django to use a different port; python manage.py runserver 8001 
# and cycle through higher numbers until you find an open port.
```

### Starting an App
A Django project is organized as a group of individual apps that work together to make the project work as a wwhole. For now, we’ll create just one app. The command startapp appname tells Django to create the infrastructure needed to build an app. You should still be running runserver in the terminal window you opened earlier.
```
python manage.py startapp learning_logs
ls
    db.sqlite3  learning_log  learning_logs  ll_env  manage.py
ls learning_logs
    admin.py  apps.py  __init__.py  migrations  models.py  tests.py  views.py
```

#### Defining Models
We’ll need to store data; for example, topics and entries as text and timestamp of each entry so we can show users when they made each entry.
```python
# learning_logs/models.py
# model tells Django how to work with data that'll be stored in the app
from django.db import models

class Topic(models.Model):
    """A topic the user is learning about"""
    text = models.CharField(max_length=200) # store characters or text
    date_added = models.DateTimeField(auto_now_add=True) # store date

    def __str__(self):
        """Return a string representation of the model. We need to tell
        Django which attribute to use by default when it displays
        information about a topic."""
        return self.text 
```

To see the different kinds of fields you can use in a model, see the [Django Model Field Reference](https://docs.djangoproject.com/en/1.8/ref/models/fields/)

#### Activating Models
To use our models, we have to tell Django to include our app in the overall project.
```python
# learning_log/settings.py
# --snip--
INSTALLED_APPS = (
    'django.contrib.admin',
    # --snip--
    # My apps
    'learning_logs',
)
# --snip--
```

Migration
```
# modify the database; store information of 'Topic'
python manage.py makemigrations learning_logs
# apply this migration and have Django modify the database
python manage.py migrate
```

#### The Django Admin Site
```
python3.6 manage.py createsuperuser
```

```python
# Registering a Model with the Admin Site
# Django automatically includes some models; 'User' and 'Group' in the
# admin site, but the models we create need to be registered manually
# learning_logs/admin.py
from django.contrib import admin

from learning_logs.models import Topic

admin.site.register(Topic)
```
Now go to [http://localhost:8000/admin/](http://localhost:8000/admin/) or http://127.0.0.1:8000/admin/ where you can not only add User or Group, but also add Topic.

#### Defining the Entry Model
read about [delete_on](https://stackoverflow.com/questions/38388423/what-does-on-delete-do-on-django-models) argument of ForeignKey()
```python
# model.py
from django.db import models


class Topic(models.Model):
    text = models.CharField(max_length=200)
    date_added = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return self.text

class Entry(models.Model):
    """Something specific learned about a topic"""
    topic = models.ForeignKey(Topic, on_delete=models.CASCADE)
    text = models.TextField()
    date_added = models.DateTimeField(auto_now_add=True)

    class Meta:
        verbose_name_plural = 'entries'

    def __str__(self):
        """Return a string representation of the model."""
        return self.text[:50] + "..."
```

#### Migrating the Entry Model
```
# Since we’ve added a new model, we need to migrate the database again
python manage.py makemigrations learning_logs
python manage.py migrate
```

#### Registering Entry with the Admin Site
```python
# admin.py
from django.contrib import admin

from learning_logs.models import Topic, Entry

admin.site.register(Topic)
admin.site.register(Entry)
```

#### The Django Shell
Now that we’ve entered some data, we can examine that data through an interactive terminal session, Django Shell.
```
python3.6 manage.py shell
>>> from learning_logs.models import Topic
>>> Topic.objects.all()
<QuerySet [<Topic: Rock Climbing>, <Topic: Chess>]>

# We can loop over a queryset just as we’d loop over a list.
>>> topics = Topic.objects.all()
>>> for topic in topics:
...     print(topic.id, topic)
1 Rock Climbing
2 Chess

# If you know the ID of a particular object, you can examine the object
>>> t = Topic.objects.get(id=1)
>>> t.text
'Rock Climbing'
>>> t.date_added
datetime.datetime(2017, 11, 1, 5, 58, 53, 244175, tzinfo=<UTC>)

# We can also look at the entries related to a certain topic. 
# Earlier we defined the 'topic' attribute for the Entry model. 
# This was a ForeignKey, a connection between each entry and a topic. 
# Django can use this connection to get every entry related to a 
# certain topic, like this:
>>> t.entry_set.all()
<QuerySet [<Entry: I love Rock Climbing more than Chess!...>,
            <Entry: I've been doing Rock Climbing for 20 years....>]>
```

### Making Pages: The Learning Log Home Page
Making web pages with Django consists of three stages: defining URLs, writing views, and writing templates. First, you must define patterns for URLs. Each URL then maps to a particular view—the view function retrieves and processes the data needed for that page. The view function often calls a template, which builds a page that a browser can read.

Mapping a URL
```python
# learning_log/urls.py
# map admin url http://localhost:8000/admin/
# map base URL http://localhost:8000/ to project's home page
from django.contrib import admin
from django.urls import path, include
from django.conf.urls import url

urlpatterns = [
    path('admin/', admin.site.urls), # url(r'^admin/', include(...)),
    url(r'', include('learning_logs.urls')),
]
```

```python
# learning_logs/urls.py
"""Defines URL patterns for learning_logs."""
from django.conf.urls import url
from django.urls import path

from . import views

app_name = 'learning_logs'
# The actual URL pattern is a call to url(regex, views.func, name='')
# urlpatterns: list of pages that can be requested from learning_logs
# regexes r'^$' tells python to interpret as a raw string and empty URL
# base project URL is ignored, so an empty regexes matches the base URL
# views.index tells python which view function to call
# name='index' provides the name for this URL patterns so we can refer
# we will later create index.html as a homepage http://localhost:8000/
urlpatterns = [
    # Home page
    url(r'^$', views.index, name='index'),
]
```

Writing a View
```python
# A view function takes in information from a request, prepares the 
# data needed to generate a page, and then sends the data back to the 
# browser, using a template that defines what the page will look like
# learning_logs/views.py
from django.shortcuts import render

def index(request):
    """The home page for Leaning Log"""
    return render(request, 'learning_logs/index.html')
```

Writing a Template
```html
<!--index.html-->
<!--A template sets up the structure for a web page.-->
<p>Leaning Log</p>
<p>Learning Log helps you keep track of your learning, for any topic 
  you're learning about.</p>
```
Now when we request the project’s base URL, http://localhost:8000/, we’ll see the page we just built instead of the default Django page. Django will take the requested URL, and that URL will match the pattern r'^$'; then Django will call the function views.index(), and this will render the page using the template contained in index.html.

Although it may seem a complicated process for creating one page, this separation between URLs, views, and templates actually works well. It allows you to think about each aspect of a project separately, and in larger projects it allows individuals to focus on the areas in which they’re strongest. For example, a database specialist can focus on the models, a programmer can focus on the view code, and a web designer can focus on the templates.

### Building Additional Pages
We'll build two pages. For each of these pages, we’ll specify a URL pattern, write a view function, and write a template. But before we do this, we’ll create a base template that all templates in the project can inherit from.

Template Inheritance
```
the template tag {\% url 'learning_logs:index' \%} generates a URL 
matching the URL pattern defined in learning_logs/urls.py with the
name 'index'. (learning_logs: namespace, index: uniquely named URL 
pattern in that namespace.) We insert a pair of block tags, named 
content. This block is a placeholder; the child template will define 
the kind of info that goes in the content block
```

```html
<!--base.html-->
<p>
  <a hred ="{% url 'learning_logs:index' %}">Learning Log</a>
</p>

{% block content %}{% endblock content %}
```
```html
<!--index.html-->
{% extends "learning_logs/base.html" %}

{% block content %}
  <p>Learning Log helps you keep track of your learning, for any topic
    you're learning about.</p>
{% endblock content %}
```

The Topics Page
```python
# The Topics URL Pattern
# learning_logs/urls.py
"""Define URL patterns for leaning_logs."""
from django.conf.urls import url
from django.urls import path

from . import views

app_name = 'learning_logs'

urlpatterns = [
    # Home page
    url(r'^$', views.index, name='index'),
    
    # Show all topics.
    url(r'^topics/$', views.topics, name='topics'),
]
```

The Topics View
```python
# views.py
from django.shortcuts import render

from .models import Topic
def index(request):
    """The home page for Learning Log"""
    return render(request, 'learning_logs/index.html')

def topics(request):
    topics = Topic.objects.order_by('date_added')  # queryset
    context = {'topics': topics}
    return render(request, 'learning_logs/topics.html', context)
```

The Topics Template
```html
<!--topics.html-->
<!--wrap the variable name in double braces; indicate to Django that 
we’re using a template variable.-->
{% extends 'learning_logs/base.html'  %}

{% block comment %}
  <p>Topics</p>
  <ul>
    {% for topic in topics %}
      <li>{{ topic }}</li>
    {% empty %}
      <li>No topics have been added yet.</li>
    {% endfor %}
  </ul>

{% endblock comment %}
```
```html
<!--base.html-->
<p>
  <a href="{% url 'learning_logs:index' %}">Learning Log</a> -
  <a href="{% url 'learning_logs:topics' %}">Topics</a>
</p>
{% block comment %}{% endblock comment %}
```

Individual Topic Pages
```python
# learning_logs/url.py
# /(?P<topic_id>\d+)/ matches an integer between two forward slashes 
# and stores the integer value in an argument called topic_id.
# ?P<topic_id> stores the matched value in topic_id;
# \d+ matches any number of digits that appear between the slashes.

"""Defines URL patterns for learning_logs."""

from django.conf.urls import url
from django.urls import path

from . import views

app_name = 'learning_logs'

urlpatterns = [
    # Home page
    url(r'^$', views.index, name='index'),
    # Show all topics.
    url(r'^topics$', views.topics, name='topics'),
    # Detail page for a single topic.for
    url(r'^topics/(?P<topic_id>\d+)/$', views.topic, name='topic'),
]
```

The Topic View
```python
# views.py
from django.shortcuts import render

from .models import Topic

def index(request):
    """The home page for Learning Log"""
    return render(request, 'learning_logs/index.html')

def topics(request):
    """Show all topics."""
    topics = Topic.objects.order_by('date_added')
    context = {'topics': topics}
    return render(request, 'learning_logs/topics.html', context)

def topic(request, topic_id):
    topic = Topic.objects.get(id=topic_id)
    entries = topic.entry_set.order_by('date_added')
    context = {'topic': topic, 'entries': entries}
    return render(request, 'learning_logs/topic.html', context)
```

The Topic Template
```html
<!--topic.html-->
<!--a vertical line, |, represents a template filter—a function that 
modifies the value in a template variable.
filter date:'M d, Y H:i' displays timestamps in the format
filter linebreaks ensures that long text entries include line breaks-->
{% extends 'learning_logs/base.html' %}

{% block comment %}
  <p>Topic: {{ topic }}</p>
  <p>Entries:</p>
  <ul>
  {% for entry in entries %}
    <li>
      <p>{{ entry.date_added|date:'M d, Y H:i' }}</p>
      <p>{{ entry.text|linebreaks }}</p>
    </li>
    {% empty %}
    <li>
      There are no entries for this topic yet.
    </li>
    {% endfor %}
  </ul>

{% endblock comment %}
```

Links from the Topics Page
```html
<!--topics.html-->
<!--modify existing topics.html to display links to each topic. URL 
template tag to generate the proper link, based on the URL pattern 
in learning_logs with the name 'topic'. This URL pattern requires a 
topic_id argument, so we add the attribute topic.id to the URL 
template tag-->
{% extends 'learning_logs/base.html'  %}

{% block comment %}
  <p>Topics</p>
  <ul>
    {% for topic in topics %}
      <li>
        <a href="{% url 'learning_logs:topic' topic.id %}">{{ topic }}</a>
      </li>
    {% empty %}
      <li>No topics have been added yet.</li>
    {% endfor %}
  </ul>

{% endblock comment %}
```

## User Accounts

### Allowing Users to Enter Data
Before we build an authentication system for creating accounts, we’ll first add some pages that allow users to enter their own data.

Currently, only a superuser can enter data through the admin site. We don’t want users to interact with the admin site, so we’ll use Django’s form-building tools to build pages that allow users to enter data.

#### Adding New Topics
Let’s start by giving users the ability to add a new topic. Adding a form-based page works in much the same way as the pages we’ve already built: we define a URL, write a view function, and write a template. The one major difference is the addition of a new module called forms.py, which will contain the forms.

The Topic ModelForm
```python
# nested Meta class telling Django which model('Topic') to base
# the form on and which fields('text') to include in the form.
# forms.py
from django import forms

from .models import Topic

class TopicForm(forms.ModelForm):
    class Meta:
        model = Topic
        fields = ['text']
        labels = {'text': ''}  # empty label
```

The new_topic URL
```python
# learning_logs/urls.py
from django.conf.urls import url
from django.urls import path

from . import views

app_name = 'learning_logs'

urlpatterns = [
    # --snip--
    # Page for adding a new topic.
    url(r'^new_topic/$', views.new_topic, name='new_topic'),
]
```

The new_topic() View Function
```python
# handles initial requests for new_topic page(show a blank form)
# and the processing of submitted data in the form
# reverse() to get the URL for the topics page and pass the URL to 
# HttpResponseRedirect(), which redirects the reader to topics page
# after they submit their topic.
# views.py

from django.shortcuts import render
from django.http import HttpResponseRedirect
from django.urls import reverse

from .models import Topic
from .forms import TopicForm

# --snip--

def new_topic(request):
    """Add a new topic."""
    if request.method != 'POST':
        # No data submitted; create a blank form.
        form = TopicForm()
    else:
        # POST data submitted; process data
        form = TopicForm(request.POST)
        if form.is_valid():
            form.save()
            return HttpResponseRedirect(reverse('learning_logs:topics'))

    context = {'form': form}
    return render(request, 'learning_logs/new_topic.html', context)
```

GET and POST Requests

The two main types of request you’ll use when building web apps are GET requests and POST requests. You use GET requests for pages that only read data from the server. You usually use POST requests when the user needs to submit information through a form. We’ll be specifying the POST method for processing all of our forms.

Depending on the request, we’ll know whether the user is requesting a blank form (a GET request) or asking us to process a completed form (a POST request).


The new_topic Template
```html
<!--new_topic.html-->
<!--'action' tells the server where to send the data submitted in the 
form; in this case, we send it back to the view function new_topic()
{% csrf_token %} prevents cross-site request forgery attack. form.as_p
to create all fields necessary to display the form automatically.
as_p modifier to render all form elements in paragraph format-->
{% extends 'learning_logs/base.html' %}

{% block comment %}
  <p>Add a new topic:</p>
  <form action="{% url 'learning_logs:new_topic' %}" method="post">
    {% csrf_token %}
    {{ form.as_p }}
    <button name="submit">add topic</button>
  </form>

{% endblock comment %}
```

Linking to the new_topic Page
```html
<!--topics.html-->
{% extends "learning_logs/base.html" %}

{% block content %}
  <p>Topics</p>
  <ul>
    # --snip--
  </ul>
  <a href="{% url 'learning_logs:new_topic' %}">Add a new topic:</a>

{% endblock content %}
```

#### Adding New Entries

The Entry ModelForm
```python
# forms.py
# widget is an HTML form element such as text box, text area, drop down
# widgets attribute to overdie django's default widget choices
# customize the input widget for the field 'text' to 80 columns wide
from django import forms
from .models import Topic, Entry
# --snip--
class EntryForm(forms.ModelForm):
    class Meta:
        model = Entry
        field = ['text']
        labels = {'text': ''}
        widgets = {'text': forms.Textarea(attrs={'cols': 80})}
```

The new_entry URL
```python
# learning_logs/urls.py
# new_entry = form.save(commit=False) to tell django to create a new 
# entry object and store it in new_entry without saving it to database
# now set entry's topic attribute, then save it to database
from django.shortcuts import render
from django.http import HttpResponseRedirect
from django.urls import reverse

from .models import Topic
from .forms import TopicForm, EntryForm
# --snip--

def new_entry(request, topic_id):
    topic = Topic.objects.get(id=topic_id)

    if request.method != 'POST':
        form = EntryForm()
    else:
        form = EntryForm(data=request.POST)
        if form.is_valid():
            new_entry = form.save(commit=False)
            new_entry.topic = topic
            new_entry.save()
            return HttpResponseRedirect(reverse('learning_logs:topic',
                                                args=[topic_id]))

        context = {'topic': topic, 'form': form}
        return render(request, 'learning_logs/new_entry.html', context)
```

The new_entry Template
```html
<!--new_entry.html-->
{% extends 'learning_logs/base.html' %}

{% block comment %}
  <p><a href="{% url 'learning_logs:topic' topic.id %}">{{ topic }}</a></p>
  
  <p>Add a new entry:</p>
  <form action="{% url 'learning_logs:new_entry' topic.id %}" method="post">
    {% csrf_token %}
    {{ form.as_p }}
    <button name="submit">add entry</button>
  </form>
{% endblock comment %}
```

Linking to the new_entry Page
```html
<!--topic.html-->
{% extends "learning_logs/base.html" %}

{% block comment %}
  <p>Topic: {{ topic }}</p>
  <p>Entries:</p>
  <p>
    <a href="{% url 'learning_logs:new_entry' topic.id %}">add new entry</a>
  </p>
  <ul>
    <!--snip-->
  </ul>

{% endblock comment %}
```

#### Editing Entries

The edit_entry URL
```python
# learning_logs/urls.py
from django.conf.urls import url
from django.urls import path

from . import views

app_name = 'learning_logs'

urlpatterns = [
    # --snip--
    # edit entry
    url(r'^edit_entry/(?P<entry_id>\d+)/$', views.edit_entry,
        name='edit_entry'),
]
```

The edit_entry() View Function
```python
# views.py
from django.shortcuts import render
from django.http import HttpResponseRedirect
from django.urls import reverse

from .models import Topic, Entry
from .forms import TopicForm, EntryForm
# --snip--
def edit_entry(request, entry_id):
    entry = Entry.objects.get(id=entry_id)
    topic = entry.topic

    if request.method != 'POST':
        # Initial request; pre-fill form with the current entry.
        form = EntryForm(instance=entry)
    else:
        form = EntryForm(instance=entry, data=request.POST)
        if form.is_valid():
            form.save()
            return HttpResponseRedirect(reverse('learning_logs:topic',
                                                args=[topic.id]))

    context = {'entry': entry, 'topic': topic, 'form': form}
    return render(request, 'learning_logs/edit_entry.html', context)
```

### Setting Up User Accounts
we’ll set up a user registration and authorization system to allow people to register an account and log in and out.

#### The users App
```
python manage.py startapp users
ls users 
```

Adding users to settings.py
```python
# settings.py
INSTALLED_APPS = [
    # --snip--
    # My apps
    'learning_logs',
    'users',
]
```

Including the URLs from users
```python
# urls.py
from django.contrib import admin
from django.urls import path, include
from django.conf.urls import url

urlpatterns = [
    path('admin/', admin.site.urls),
    url(r'^users/', include('users.urls')),
    url(r'', include('learning_logs.urls')),
]
```

#### The Login Page
```python
# users/urls.py
# pass a dictionary telling where to find the template, login.html
from django.conf.urls import url
from django.contrib.auth.views import login
from . import views

app_name = 'users'

urlpatterns = [
    # Login page
    url(r'^login/$', login, {'template_name': 'users/login.html'},
        name='login'),
]
```

The login template
```html
<!--login.html-->
<!--the value argument tells Django where to redirect the user after 
they’ve logged in successfully; redirect to learning_logs home-->
{% extends 'learning_logs/base.html' %}

{% block content %}
  {% if form.errors %}
    <p>Your username and password didn't match. Please try again.</p>
  {% endif %}
  <form action="{% url 'users:login' %}" method="post">
    {% csrf_token %}
    {{ form.as_p }}
    <button name="submit">log in</button>

    <input type="hidden" name="next"
           value="{% url 'learning_logs:index' %}" />
  </form>
{% endblock content %}
```

Linking to the Login Page
```html
<!--base.html-->
<!--In Django’s authentication system, every template has user variable
available, which always has an is_authenticated attribute set-->
<p>
  <a href="{% url 'learning_logs:index' %}">Learning Log</a> -
  <a href="{% url 'learning_logs:topics' %}">Topics</a> -
  {% if user.is_authenticated %}
    Hello, {{ user.username }}.
  {% else %}
    <a href="{% url 'user:login' %}">log in</a>
  {% endif %}
</p>

{% block content %}{% endblock content %}
```

#### Logging Out
We won’t build a page for logging out; users will just click a link and be sent back to the home page.

The logout URL
```python
# urls.py
from django.conf.urls import url
from django.contrib.auth.views import login
from . import views

urlpatterns = [
    url(r'^login/$', login, {'template_name': 'users/login.html'},
        name='login'),
    url(r'^logout/$', views.logout_view, name='logout'),
]
```

The logout_view() View Function
```python
# logout_view.py
# import and call django's logout(), and redirect back to home page 
from django.http import HttpResponseRedirect
from django.urls import reverse
from django.contrib.auth import logout

def logout_view(request):
    logout(request)
    return HttpResponseRedirect(reverse('learning_logs:index'))
```

Linking to the logout View
```html
<!--base.html-->
<p>
  <a href="{% url 'learning_logs:index' %}">Learning Log</a> -
  <a href="{% url 'learning_logs:topics' %}">Topics</a> -
  {% if user.is_authenticated %}
    Hello, {{ user.username }}.
    <a href="{% url 'users:logout' %}">log out</a>
  {% else %}
    <a href="{% url 'users:login' %}">log in</a>
  {% endif %}ls
  
</p>

{% block content %}{% endblosck content %}
```

#### The Registration Page
We'll use Django’s default UserCreationForm but write our own view function and template.

The register URL
```python
# urls.py
from django.conf.urls import url
from django.contrib.auth.views import login
from . import views

urlpatterns = [
    # --snip--
    url(r'^register/$', views.register, name='register'),
```

The register() View Function
```python
# views.py
from django.shortcuts import render
from django.http import HttpResponseRedirect
from django.urls import reverse
from django.contrib.auth import logout, login, authenticate
from django.contrib.auth.forms import UserCreationForm

def logout_view(request):
    # --snip--
    
def register(request):
    """Register a new user."""
    if request.method != 'POST':
        # Display blank registration form.
        form = UserCreationForm()
    else:
        # Process completed Form.
        form = UserCreationForm(data=request.POST)
        if form.is_valid():
            new_user = form.save()
            # Log the user in and then redirect to home page.
            authenticated_user = authenticate(
                username=new_user.username,
                password=request.POST['password1'])
            login(request, authenticated_user)
            return HttpResponseRedirect(reverse('learning_logs:index'))

    context = {'form': form}
    return render(request, 'users/register.html', context)
```

The register Template
```html
<!--users/register.html-->
{% extends 'learning_logs/base.html' %}

{% block content %}
  <form method="post" action="{% url 'users:register' %}">
    {% csrf_token %}
    {{ form.as_p }}
    <button name="submit">register</button>
    <input type="hidden" name="next"
           value="{% url 'learning_logs:index' %}"/>
  </form>

{% endblock content %}
```

Linking to the Registration Page
```html
<!--base.html-->
<p>
  <a href="{% url 'learning_logs:index' %}">Learning Log</a> -
  <a href="{% url 'learning_logs:topics' %}">Topics</a> -
  {% if user.is_authenticated %}
    Hello, {{ user.username }}.
    <a href="{% url 'users:logout' %}">log out</a>
  {% else %}
    <a href="{% url 'users:register' %}">register</a> -
    <a href="{% url 'users:login' %}">log in</a>
  {% endif %}
</p>

{% block content %}{% endblock content %}
```

### Allowing Users to Own Their Data
Users should be able to enter data exclusive to them, so we’ll create a system to figure out which data belongs to which user, and then we’ll restrict access to certain pages so users can work with only their own data.

#### Restricting Access with @login_required
A decorator is a directive placed just before a function definition that Python applies to the function before it runs to alter how the function code behaves. @login_required decorator:

Restricting Access to the Topics Page
```python
# views.py
# Each topic will be owned by a user. Only registered users should be
# able to request the topics page. login_required() checks to see if
# a user is logged in, and Django will run the code in topics() only if
# they are. If the user is not logged in, redirect to the login page.
# --snip--
from django.contrib.auth.decorators import login_required

from .models import Topic, Entry
from .forms import TopicForm, EntryForm

# --snip--
@login_required
def topics(request):
    # --snip--

@login_required
def topic(request, topic_id):
    # --snip--

@login_required
def new_topic(request):
    # --snip--

@login_required
def new_entry(request, topic_id):
    # --snip--

@login_required
def edit_entry(request, entry_id):
    # --snip
```

```python
# settings.py
# When unauthenticated user requests page protected by @login_required 
# user will receive URL defined in LOGIN_URL in settings.py
# --snip--
LOGIN_URL = '/users/login/'
```

#### Connecting Data to Certain Users
Now we need to connect the data submitted to the user who submitted it. We need to connect only the data highest in the hierarchy to a user, and the lower-level data will follow.

Modifying the Topic Model
```python
# models.py
# add a foreign key relationship to a user
from django.db import models
from django.contrib.auth.models import User

class Topic(models.Model):
    text = models.CharField(max_length=200)
    date_added = models.DateTimeField(auto_now_add=True)
    owner = models.ForeignKey(User)
    def __str__(self):
        return self.text

class Entry(models.Model):
    # --snip--
```

Identifying Existing Users
```
python manage.py shell
>>> from django.contrib.auth.models import User
>>> User.objects.all()
<QuerySet [<User: ll_admin>, <User: lopehih>, <User: abc>, <User: aaa>]>
>>> for user in User.objects.all():
...     print(user.username, user.id)
... 
ll_admin 1
lopehih 2
abc 3
aaa 4
```

Migrating the Database
```
python manage.py makemigrations learning_logs

You are trying to add a non-nullable field 'owner' to topic without a 
default; we can't do that (the database needs something to populate 
existing rows).
Please select a fix:
    1) Provide a one-off default now (will be set on all existing rows)
    2) Quit, and let me add a default in models.py
Select an option: 1
Please enter the default value now, as valid Python
>>> 1

python manage.py migrate

# verify that the migration worked as expected
python manage.py shell
>>> from learning_logs.models import Topic
>>> for topic in Topic.objects.all():
   ...     print(topic, topic.owner)
   ...
   Rock Climbing ll_admin
   Chess ll_admin
```

#### Restricting Topics Access to Appropriate Users
Currently, if you’re logged in, you’ll be able to see all the topics, no matter which user you’re logged in as. We’ll change that by showing users only the topics that belong to them.
```python
# views.py
# To see if this works, log in as the user you connected all existing 
# topics to, and go to the topics page. You should see all the topics. 
# Now log out, and log back in as a different user. The topics page 
# should list no topics.
from django.shortcuts import render
from django.urls import reverse
from django.http import HttpResponseRedirect
from django.contrib.auth.decorators import login_required

from .models import Topic, Entry
from .forms import TopicForm, EntryForm

@login_required
def topics(request):
    topics = Topic.objects.filter(owner=request.user).order_by('date_added')
    context = {'topics': topics}
    return render(request, 'learning_logs/topics.html', context)
# --snip--
```

#### Protecting a User’s Topics
Any registered user could try a bunch of URLs, like http://localhost:8000/topics/1/, and retrieve topic pages that happen to match. We'll fix this by performing a check before retrieving the requested entries in topic() view function
```python
# views.py
# If the current user doesn’t own the requested topic, we raise the 
# Http404 exception. Now if you try to view another user’s topic 
# entries, you’ll see a Page Not Found message from Django. Later, 
# we’ll configure the project so users will see a proper error page.
from django.shortcuts import render
from django.http import HttpResponseRedirect, Http404
from django.urls import reverse
    # --snip--
    @login_required
    def topic(request, topic_id):
        topic = Topic.objects.get(id=topic_id)
        # Make sure the topic belongs to the current user.
        if topic.owner != request.user:
           raise Http404

        entries = topic.entry_set.order_by('-date_added')
        context = {'topic': topic, 'entries': entries}
        return render(request, 'learning_logs/topic.html', context)
    # --snip--
```


#### Protecting the edit_entry Page
Any registered user can access http://localhost:8000/edit_entry/entry_id/. Let’s protect this page so no one can use the URL to gain access to someone else’s entries
```python
# --snip--
@login_required
def edit_entry(request, entry_id):
    """Edit an existing entry."""
    entry = Entry.objects.get(id=entry_id)
    topic = entry.topic
    if topic.owner != request.user:
        raise Http404
    if request.method != 'POST':
        # Initial request; pre-fill form with the current entry.
        # --snip--
```

#### Associating New Topics with the Current User
```python
# --snip--

@login_required
def new_topic(request):
    if request.method != 'POST':
        form = TopicForm()
    else:
        form = TopicForm(request.POST)
        if form.is_valid():
            new_topic = form.save(commit=False)
            new_topic.owner = request.user
            new_topic.save()
            return HttpResponseRedirect(reverse('learning_logs:topics'))

    context = {'form': form}
    return render(request, 'learning_logs/new_topic.html', context)
# --snip--
```

## Styling and Deploying an App

### Styling Learning Log

#### The django-bootstrap3 App
We’ll use django-bootstrap3 to integrate Bootstrap into our project.
```
pip3 install django-bootstrap3
```

```python
# settings.py
INSTALLED_APPS = [
    # --snip--
    # Third party apps: include django-boostrap3 
    'bootstrap3',
    # My apps
    'learning_logs',
    'users',
]

# My settings
LOGIN_URL = '/users/login/'

# We need django-bootstrap3 to include jQuery, a JavaScript library 
# that enables some of the interactive elements that the Bootstrap 
# template provides.
# Settings for django-bootstrap3
BOOTSTRAP3 = {
    'include_jquery': True,
    }
```

#### Using Bootstrap to Style Learning Log
[Bootstrap](http://getbootstrap.com/) is basically a large collection of styling tools. It also has a number of templates you can apply to your project to create a particular overall style.

#### Modifying base.html

```html
<!--base.html-->
{% load bootstrap3 %}

<!DOCTYPE html>
<html lang="en">
  <!--Defining the HTML Headers-->
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>Learning Log</title>
    {% bootstrap_css %}
    {% bootstrap_javascript %}
  </head>
  <body>
    <!--Defining the Navigation Bar-->
    <nav class="navbar navbar-default navbar-static-top">
      <div class="container">

        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed"
                  data-toggle="collapse" data-target="#navbar"
                  aria-expanded="false" aria-controls="navbar">
          </button>
          <a class="navbar-brand" href="{% url 'learning_logs:index' %}">
            Learning Log</a>
        </div>

        <div id="navbar" class="navbar-collapse collapse">
          <ul class="nav navbar-nav">
            <li><a href="{% url 'learning_logs:topics' %}">Topics</a></li>
          </ul>

          <ul class="nav navbar-nav navbar-right">
            {% if user.is_authenticated %}
              <li><a>Hello, {{ user.username }}.</a></li>
              <li><a href="{% url 'users:logout' %}">log out</a></li>
            {% else %}
              <li><a href="{% url 'users:register' %}">register</a></li>
              <li><a href="{% url 'users:login' %}">log in</a></li>
            {% endif %}
          </ul>
        </div><!--/.navbar-collapse-->

      </div>
    </nav>
    
    <!--Defining the Main Part of the Page-->
    <div class="container">
      <div class="page-header">
        {% block header %}{% endblock header %}
      </div>
      <div>
        {% block content %}{% endblock content %}
      </div>
    </div> <!--/container-->
    
  </body>
</html>
```

#### Styling the Home Page Using a Jumbotron
Let’s update the home page using the newly defined header block and another Bootstrap element called a jumbotron—a large box that will stand out from the rest of the page and can contain anything you want. It’s typically used on home pages to hold a brief description of the overall project.
```html
<!--index.html-->
{% extends 'learning_logs/base.html' %}

{% block header %}
  <div class="jumbotron">
    <h1>Track your learning.</h1>
  </div>
{% endblock header %}

{% block content %}
  <h2>
    <a href="{% url 'users:register' %}">Register an account</a> to make
    your own Learning Log, and list the topics you're learning about.
  </h2>
  <h2>
    Whenever you learn something new about a topic, make an entry
    summarizing what you've learned.
  </h2>
{% endblock content %}
```

#### Styling the Login Page
```html
<!--login.html-->
<!--form.errors validation check is handled by django-bootstrap3-->
{% extends 'learning_logs/base.html' %}
{% load bootstrap3 %}

{% block header %}
  <h2>Log in to your account.</h2>
{% endblock header %}

{% block content %}
  <form class="form" action="{% url 'users:login' %}" method="post">
    {% csrf_token %}

    {% bootstrap_form form %}

    {% buttons %}
      <button name="submit" class="btn btn-primary">log in</button>
    {% endbuttons %}

    <input type="hidden" name="next" value="{% url 'learning_logs:index' %}" />
  </form>
  
{% endblock content %}
```

#### Styling the new_topic Page
```html
<!--new_topic.html-->
{% extends 'learning_logs/base.html' %}
{% load bootstrap3 %}

{% block header %}
  <h2>Add a new topic:</h2>
{% endblock header %}

{% block content %}
  <form class='form' action="{% url 'learning_logs:new_topic' %}" method="post">
    {% csrf_token %}
    {% bootstrap_form form %}

    {% buttons %}
      <button name="submit" class="btn btn-primary">add topic</button>
    {% endbuttons %}
  </form>

{% endblock content %}
```

#### Styling the Topics Page
```html
{% extends 'learning_logs/base.html' %}

{% block header %}
  <h1>Topics</h1>
{% endblock header %}

{% block content %}
  <p>Topics:</p>
  <ul>
    {% for topic in topics %}
      <li>
        <h3>
          <a href="{% url 'learning_logs:topic' topic.id %}">{{ topic }}</a>
        </h3>
      </li>
    {% empty %}
      <li>No topics added yet.</li>
    {% endfor %}
  </ul>
  <h3>
    <a href="{% url 'learning_logs:new_topic' %}">Add new topic</a>
  </h3>
  
{% endblock content %}
```

#### Styling the Entries on the Topic Page
```html
<!--topic.html-->
{% extends 'learning_logs/base.html' %}

{% block header %}
  <h2>{{ topic }}</h2>
{% endblock header %}

{% block content %}
  <p>
    <a href="{% url 'learning_logs:new_entry' topic.id %}">Add new entry</a>
  </p>

  {% for entry in entries %}
    <div class="panel panel-default">
      <div class="panel-heading">
        <h3>
          <p>{{ entry.date_added|date:'M d, Y H:i' }}</p>
          <small>
            <a href="{% url 'learning_logs:edit_entry' entry.id %}">Edit this entry</a>
          </small>
        </h3>
      </div>
      <div class="panel-body">
        <p>{{ entry.text|linebreaks }}</p>
      </div>
    </div> <!--panel-->
  {% empty %}
    No entries added yet.
  {% endfor %}

{% endblock content %}
```

### Deploying Learning Log
Now that we have a professional-looking project, let’s deploy it to a live server so anyone with an internet connection can use it. We’ll use Heroku, a web-based platform that allows you to manage the deployment of web applications. We’ll get Learning Log up and running on Heroku.

#### Making a Heroku Account
signup for [Heroku](https://heroku.com/)

#### Installing the Heroku Toolbelt
install [Heroku CLI](https://toolbelt.heroku.com/)
```
sudo snap install --classic heroku

heroku --version
```

#### Installing Required Packages
You’ll also need to install a number of packages that help serve Django projects on a live server. In an active virtual environment, issue the following commands to install packages:
```commandline
# dj-database-url helps Django communicate with the database Heroku uses
# dj-static and static3 help Django manage static files correctly
# WhiteNoise to manage static files instead of dj-static and static3
#   also use DjangoWhiteNoise() instead of Cling() in wsgi.py
#   and add STATICFILES_STORAGE variable in settings.py heroku settings 
# (Static files contain style rules and JavaScript files.)
# gunicorn is a server capable of serving apps in a live environment. 


pip install dj-database-url
pip install whitenoise
pip install gunicorn
```

```commandline
pip uninstall dj-static
pip uninstall static3
```

#### Creating a Packages List with a requirements.txt File
Heroku needs to know which packages our project depends on, so we’ll use pip to generate a file listing them. The freeze command tells pip to write the names of all the packages currently installed in the project into the file requirements.txt.
```
pip3 freeze > requirements.txt
```

Learning Log already depends on 7 different packages (remove the line pkg-resources==0.0.0) with specific version numbers, so it requires a specific environment to run properly. When we deploy Learning Log, Heroku will install all the packages listed in requirements.txt, creating an environment with the same packages we’re using locally. For this reason, we can be confident the deployed project will behave the same as it does on our local system. This is a huge advantage as you start to build and maintain various projects on your system.

Next, we need to add psycopg2, which helps Heroku manage the live database, to the list of packages. Open requirements.txt and add the line psycopg2>=2.6.1. This will install version 2.6.1 of psycopg2, or a newer version if it’s available:
```
# requirements.txt
dj-database-url==0.4.2
Django==2.0
django-bootstrap3==9.1.0
gunicorn==19.7.1
pytz==2017.3
whitenoise==3.3.1
psycopg2>=2.6.1
```

#### Specifying the Python Runtime
Let’s make sure Heroku uses the same version(3.6.3) of Python we’re using. Create file named runtime.txt and put 'python-3.6.3'
```
python --version
    Python 3.6.3

# runtime.txt
python-3.6.3
```

#### Modifying settings.py for Heroku
Now we need to add a section at the end of settings.py to define some settings specifically for the Heroku environment:
```python
# settings.py
# getcwd() gets the current working directory the file is running from. 
# /app in a Heroku deployment, and project folder name in local
# import dj_database_url to help configure the database on Heroku. 
# PostgreSQL is a more advanced database than SQLite
# SECURE_PROXY_SSL_HEADER to support HTTPS requests 
# ['*'] to ensure that Django will serve the project from Heroku’s URL
# set up the project to serve static files correctly on Heroku
# When you push a Django project to Heroku, the project is initially 
# built from a temporary directory. If we only look for the /app 
# directory, the project will fail to build properly.
# --snip--

# Heroku settings
cwd = os.getcwd()
print("--- CWD ---\n", cwd, "\n---\n")
if os.getcwd() == '/app' or cwd[:4] == '/tmp':
    import dj_database_url
    DATABASES = {
        'default': dj_database_url.config(default='postgres://localhost')
    }

    # Honor the 'X-Forwarded-Proto' header for request.is_secure().
    SECURE_PROXY_SSL_HEADER = ('HTTP_X_FORWARDED_PROTO', 'https')

    # Allow all host headers.
    ALLOWED_HOSTS = ['*']

    # Static asset configuration
    BASE_DIR = os.path.dirname(os.path.abspath(__file__))
    STATIC_ROOT = 'staticfiles'
    STATICFILES_DIRS = (
        os.path.join(BASE_DIR, 'static')
    )
    
    STATICFILES_STORAGE = 'whitenoise.django.GzipManifestStaticFilesStorage'
```

#### Making a Procfile to Start Processes
A Procfile tells Heroku which processes to start in order to serve the project properly. This is a one-line file that you should save as 'Procfile', with an uppercase P and no file extension, in the same directory as manage.py. This line tells Heroku to use gunicorn as a server and to use the settings in learning_log/wsgi.py to launch the app. The log-file flag tells Heroku the kinds of events to log.
```
web: gunicorn learning_log.wsgi --log-file -
```

#### Modifying wsgi.py for Heroku
We also need to modify wsgi.py for Heroku, because Heroku needs a slightly different setup than what we’ve been using:
```python
# wsgi.py
# import Cling, which helps serve static files correctly, and use it to 
# launch the application. CHANGE it to using django whitenoise package
import os

from django.core.wsgi import get_wsgi_application
# from dj_static import Cling
from whitenoise.django import DjangoWhiteNoise

os.environ.setdefault("DJANGO_SETTINGS_MODULE", "learning_log.settings")

# application = Cling(get_wsgi_application())
application = get_wsgi_application()
application = DjangoWhiteNoise(application)
```

#### Making a Directory for Static Files
On Heroku, Django collects all the static files and places them in one place so it can manage them efficiently. We’ll create a directory for these static files. Inside the learning_log folder we’ve been working from is another folder called learning_log. In this nested folder, make a new folder called static with the path learning_log/learning_log/static/. We also need to make a placeholder file to store in this directory for now, because empty directories won’t be included in the project when it’s pushed to Heroku. In the static/ directory, make a file called placeholder.txt:
```text
# placeholder.txt
This file ensures that learning_log/static/ will be added to the project.
Django will collect static files and place them in learning_log/static/.
```

#### Using the gunicorn Server Locally (Linux or OS X)
If you’re using Linux or OS X, you can try using the gunicorn server locally before deploying to Heroku. From an active virtual environment, run the command 'heroku local' to start the processes defined in Procfile:
```commandline
# The first time you run heroku local, a number of packages from the 
# Heroku Toolbelt will be installed. 
# gunicorn has been started with a process id of 25525. 
# gunicorn is listening for requests on port 5000. 
# gunicorn has started a worker process (25528) to help it serve requests
# http://localhost:5000/ to make sure everything is working; 
# you should see the Learning Log home page, just as it appears when 
# you use the Django server (runserver). Press CTRL-C to stop the 
# processes started by heroku local. You should continue to use runserver 
# for local development.

heroku local

    [WARN] No ENV file found
    23:27:05 web.1   |  [2017-12-11 23:27:05 +0900] [25525] [INFO] Starting gunicorn 19.7.1
    23:27:05 web.1   |  [2017-12-11 23:27:05 +0900] [25525] [INFO] Listening at: http://0.0.0.0:5000 (25525)
    23:27:05 web.1   |  [2017-12-11 23:27:05 +0900] [25525] [INFO] Using worker: sync
    23:27:05 web.1   |  [2017-12-11 23:27:05 +0900] [25528] [INFO] Booting worker with pid: 25528
```

#### Using Git to Track the Project’s Files

Installing Git
```commandline
# The Heroku Toolbelt includes Git, so it should already be installed on your system.
git --version
    git version 2.7.4
```

Configuring Git
```commandline
# Git keeps track of who makes changes to a project, even in cases like 
# this when there’s only one person working on the project. To do this, 
# Git needs to know your username and email. You have to provide a username, 
# but feel free to make up an email for your practice projects:

git config --global user.name 'fggo'
git config --global user.email 'a@a.com'
```

Ignoring Files
```
# .gitignore
# We don’t need Git to track every file in the project, so we’ll tell Git 
# to ignore some files. Make a file called .gitignore in the folder that 
# contains manage.py. Notice that this filename begins with a dot and 
# has no file extension. Here’s what goes in .gitignore:

ll_env/
__pycache_/
*.sqlite3
```

Committing the Project
```commandline
git init
git add .
git commit -am 'Ready for deployment to heroku.'
git status
```

#### Pushing to Heroku
'heroku create' tells Heroku to build an empty project; Heroku generates a name made up of two words and a number; you can change this later on. We then issue the command 'git push heroku master', which tells Git to push the master branch of the project to the repository Heroku just created. Heroku then builds the project on its servers using these files. Then it will show deployed URL we’ll use to access the live project.
```commandline
heroku login
    Enter your Heroku credentials:
    Email: example@mail.com
    Password: 
    Logged in as example@mail.com
heroku create
    Creating app... done, ⬢ lit-lowlands-35015
    https://lit-lowlands-35015.herokuapp.com/ | https://git.heroku.com/lit-lowlands-35015.git
git push heroku master
```

NOTE: using dj-static and static3 packages you will get an ERROR while git push command. You will not see these errors if you're using whitenoise package for handling static files. 'python manage.py collectstatic --noinput command' will be timed out Check [workaround](https://stackoverflow.com/questions/36665889/collectstatic-error-while-deploying-django-app-to-heroku) 
```commandline
# FIRST delete pkg-resources==0.0.0 line from requirements.txt and then:
heroku config:set DISABLE_COLLECTSTATIC=0
git push heroku master

heroku run python manage.py migrate

sudo npm install -g bower
sudo ln -s /usr/bin/nodejs /usr/bin/node
sudo bower --version
heroku run 'bower install --config.interactive=false;grunt prep;python manage.py collectstatic --noinput'

heroku config:unset DISABLE_COLLECTSTATIC
heroku run python manage.py collectstatic
```

When you’ve issued these commands, the project is deployed but not fully configured. To check that the server process started correctly, use the heroku ps command:
```commandline
heroku ps
    Free dyno hours quota remaining this month: 550h 0m (100%)
    For more information on dyno sleeping and how to upgrade, see:
    https://devcenter.heroku.com/articles/dyno-sleeping
    
    === web (Free): gunicorn learning_log.wsgi --log-file - (1)
    web.1: up 20xx/xx/xx 00:00:00 +0000 (~ 1m ago)
```

The output shows how much more time the project can be active in the next 24 hours. At the time of this writing, Heroku allows free deployments to be active for up to 18 hours in any 24-hour period. If a project exceeds these limits, a standard server error page will be displayed; we’ll customize this error page shortly. we see that the process defined in Procfile has been started.

Now we can open the app in a browser using the command heroku open:
```commandline
heroku open
```
This command spares you from opening a browser and entering the URL Heroku showed you, but that’s another way to open the site. You should see the home page for Learning Log, styled correctly. However, you can’t use the app yet because we haven’t set up the database.


#### Setting Up the Database on Heroku
We need to run migrate once to set up the live database and apply all the migrations we generated during development. You can run Django and Python commands on a Heroku project using the command heroku run. Here’s how to run migrate on the Heroku deployment:
```commandline
heroku run python manage.py migrate
```

Now when you visit your deployed app, you should be able to use it just as you did on your local system. However, you won’t see any of the data you entered on your local deployment, because we didn’t copy the data to the live server. This is normal practice: you don’t usually copy local data to a live deployment because the local data is usually test data.

#### Refining the Heroku Deployment
In this section we’ll refine the deployment by creating a superuser, just as we did locally. We’ll also make the project more secure by changing the setting DEBUG to False, so users won’t see any extra information in error messages that they could use to attack the server.

Creating a Superuser on Heroku
```commandline
heroku run bash

~$ ls
~$ python manage.py createsuperuser
        username: ll_admin
        email: a@mail.com
        password: 
~$ exit 
```
Now you can add /admin/ to the end of the URL for the live app and log in to the admin site. For example, the URL is https://afternoon-meadow-2775.herokuapp.com/admin/


Creating a User-Friendly URL on Heroku
```commandline
heroku apps:rename learning-log
```
This deployment now lives at https://learning-log-1337.herokuapp.com/. The project is no longer available at the previous URL; the apps:rename command completely moves the project to the new URL.


#### Securing the Live Project
One glaring security issue exists in the way our project is currently deployed: the setting DEBUG=True in settings.py, which provides debug messages when errors occur. Django’s error pages give you vital debugging information when you’re developing a project, but they give way too much information to attackers if you leave them enabled on a live server. We also need to make sure no one can get information or redirect requests by pretending to be the project’s host.

Let’s modify settings.py so we can see error messages locally but not on the live deployment:
```python
# settings.py
import os
# --snip--
# Heroku settings
cwd = os.getcwd()
print("--- CWD ---\n", cwd, "\n---\n")
if cwd == '/app' or cwd[:4] == '/tmp':
    # --snip--
    # Only allow heroku to host the project.
    ALLOWED_HOSTS = ['learning-log.herokuapp.com']
    
    DEBUG = False
    # --snip--
```

#### Committing and Pushing Changes
```commandline
git commit -am 'Set DEBUG=False for Heroku.'
git status

git push heroku master
```

#### Creating Custom Error Pages
A 404 error usually means your Django code is correct, but the object being requested doesn’t exist; a 500 error usually means there’s an error in the code you’ve written, such as an error in a function in views.py. Currently, Django returns the same generic error page in both situations, but we can write our own 404 and 500 error page templates that match the overall appearance of Learning Log. These templates must go in the root template directory.

Making Custom Templates
```html
<!--404.html-->
<!--in /learning_log/learning_log/templates/-->
{% extends 'learning_logs/base.html' %}

{% block header %}
  <h2>The item you requested is not available. (404)</h2>
{% endblock header %}
```
```html
<!--500.html-->
{% extends 'learning_logs/base.html' %}

{% block header %}
  <h2>There has been an internal error. (500)</h2>
{% endblock header %}
```

```python
# settings.py
# This change tells Django to look in the root template directory for 
# the error page templates.
# --snip--
TEMPLATES = [
    {
        'BACKEND': 'django.template.backends.django.DjangoTemplates',
        'DIRS': [os.path.join(BASE_DIR, 'learning_log/templates')],
        'APP_DIRS': True,
        # --snip--
    },
]
```

Viewing the Error Pages Locally (do not apply to Heroku Deployment)
```python
# settings.py
# for example try 'localhost:8000/letmein' in browser to see error page
# --snip--
# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = False

ALLOWED_HOSTS = ['localhost']
# --snip--
```
Pushing the Changes to Heroku
```commandline
git status
git add .
git commit -am 'added custom 404 500 error pages'
git push heroku master
```

At this point, if a user manually requests a topic or entry that doesn’t exist, they’ll get a 500 server error. Django tries to render the page but it doesn’t have enough information to do so, and the result is a 500 error. This situation is more accurately handled as a 404 error, and we can implement this behavior with the Django shortcut function get_object_or_404(). This function tries to get the requested object from the database, but if that object doesn’t exist, it raises a 404 exception. We’ll import this function into views.py and use it in place of get():

Using the get_object_or_404() Method
```python
# views.py
def topic(request, topic_id):
    topic = get_object_or_404(Topic, id=topic_id)
    if topic.owner != request.user:
        raise Http404

    entries = topic.entry_set.order_by('date_added')
    context = {'topic': topic, 'entries': entries}
    return render(request, 'learning_logs/topic.html', context)
```
Now when you request a topic that doesn’t exist (for example, http://localhost:8000/topics/999999/), you’ll see a 404 error page. To deploy this change, make a new commit, and then push the project to Heroku.
```commandline
git add .
git commit -m '404 handling'
git push heroku master
```

#### Ongoing Development

#### The SECRET_KEY Setting

#### Deleting a Project on Heroku
```commandline
heroku apps:destroy --app appname
```


```python
name = "A"
```


```python
name
```




    'A'




```python
name = "C"
```


```python
!pip3 list
```

    Package             Version
    ------------------- ---------
    appnope             0.1.0
    argon2-cffi         20.1.0
    astroid             2.4.2
    async-generator     1.10
    attrs               20.2.0
    backcall            0.2.0
    bleach              3.1.5
    certifi             2020.6.20
    cffi                1.14.2
    chardet             3.0.4
    cycler              0.10.0
    decorator           4.4.2
    defusedxml          0.6.0
    entrypoints         0.3
    idna                2.10
    importlib-metadata  1.7.0
    ipykernel           5.3.4
    ipython             7.18.1
    ipython-genutils    0.2.0
    ipywidgets          7.5.1
    isort               5.5.2
    jedi                0.17.2
    Jinja2              2.11.2
    jsonschema          3.2.0
    jupyter             1.0.0
    jupyter-client      6.1.7
    jupyter-console     6.2.0
    jupyter-core        4.6.3
    jupyterlab-pygments 0.1.1
    kiwisolver          1.2.0
    lazy-object-proxy   1.4.3
    MarkupSafe          1.1.1
    matplotlib          3.3.1
    mccabe              0.6.1
    mistune             0.8.4
    nbclient            0.5.0
    nbconvert           6.0.1
    nbformat            5.0.7
    nest-asyncio        1.4.0
    notebook            6.1.4
    numpy               1.19.2
    packaging           20.4
    pandocfilters       1.4.2
    parso               0.7.1
    pexpect             4.8.0
    pickleshare         0.7.5
    Pillow              7.2.0
    pip                 20.2.3
    prometheus-client   0.8.0
    prompt-toolkit      3.0.7
    ptyprocess          0.6.0
    pycparser           2.20
    Pygments            2.6.1
    pylint              2.6.0
    pyparsing           2.4.7
    pyrsistent          0.16.0
    python-dateutil     2.8.1
    pyzmq               19.0.2
    qtconsole           4.7.7
    QtPy                1.9.0
    requests            2.24.0
    Send2Trash          1.5.0
    setuptools          40.8.0
    six                 1.15.0
    terminado           0.8.3
    testpath            0.4.4
    toml                0.10.1
    tornado             6.0.4
    traitlets           5.0.4
    typed-ast           1.4.1
    urllib3             1.25.10
    wcwidth             0.2.5
    webencodings        0.5.1
    widgetsnbextension  3.5.1
    wrapt               1.12.1
    zipp                3.1.0



```python
%lsmagic
```




    Available line magics:
    %alias  %alias_magic  %autoawait  %autocall  %automagic  %autosave  %bookmark  %cat  %cd  %clear  %colors  %conda  %config  %connect_info  %cp  %debug  %dhist  %dirs  %doctest_mode  %ed  %edit  %env  %gui  %hist  %history  %killbgscripts  %ldir  %less  %lf  %lk  %ll  %load  %load_ext  %loadpy  %logoff  %logon  %logstart  %logstate  %logstop  %ls  %lsmagic  %lx  %macro  %magic  %man  %matplotlib  %mkdir  %more  %mv  %notebook  %page  %pastebin  %pdb  %pdef  %pdoc  %pfile  %pinfo  %pinfo2  %pip  %popd  %pprint  %precision  %prun  %psearch  %psource  %pushd  %pwd  %pycat  %pylab  %qtconsole  %quickref  %recall  %rehashx  %reload_ext  %rep  %rerun  %reset  %reset_selective  %rm  %rmdir  %run  %save  %sc  %set_env  %store  %sx  %system  %tb  %time  %timeit  %unalias  %unload_ext  %who  %who_ls  %whos  %xdel  %xmode
    
    Available cell magics:
    %%!  %%HTML  %%SVG  %%bash  %%capture  %%debug  %%file  %%html  %%javascript  %%js  %%latex  %%markdown  %%perl  %%prun  %%pypy  %%python  %%python2  %%python3  %%ruby  %%script  %%sh  %%svg  %%sx  %%system  %%time  %%timeit  %%writefile
    
    Automagic is ON, % prefix IS NOT needed for line magics.




```python
%pwd
```




    '/Users/agnt/jnuho.github.io'




```python
%ls -al
```

    total 16640
    drwxr-xr-x  46 agnt  staff     1472  9 13 01:37 [1m[36m.[m[m/
    drwxr-xr-x+ 73 agnt  staff     2336  9 13 01:31 [1m[36m..[m[m/
    -rw-r--r--@  1 agnt  staff     6148  9 12 18:23 .DS_Store
    drwxr-xr-x  15 agnt  staff      480  9 13 01:31 [1m[36m.git[m[m/
    -rw-r--r--   1 agnt  staff       80  9 12 20:51 .gitignore
    drwxr-xr-x   3 agnt  staff       96  9 13 00:47 [1m[36m.ipynb_checkpoints[m[m/
    -rw-r--r--   1 agnt  staff     2936  8  2 16:57 404.html
    -rw-r--r--   1 agnt  staff     1852  9 12 20:44 CODING_TEST.sql
    -rw-r--r--   1 agnt  staff      126  8  2 16:57 Gemfile
    -rw-r--r--   1 agnt  staff     1105  8  2 16:57 LICENSE
    -rw-r--r--   1 agnt  staff     1157  8  2 16:57 README.md
    -rw-r--r--   1 agnt  staff     7110  9 12 21:00 README_java.md
    -rw-r--r--   1 agnt  staff  7989959  8  2 16:57 Semi-Project.pdf
    -rw-r--r--   1 agnt  staff      920  9 13 00:05 _config.yml
    drwxr-xr-x   7 agnt  staff      224  8  2 16:57 [1m[36m_includes[m[m/
    drwxr-xr-x   5 agnt  staff      160  8  2 16:57 [1m[36m_layouts[m[m/
    drwxr-xr-x   5 agnt  staff      160  9 12 20:25 [1m[36m_posts[m[m/
    drwxr-xr-x   8 agnt  staff      256  8  2 16:57 [1m[36massets[m[m/
    -rw-r--r--   1 agnt  staff     1852  9 12 03:38 categories.html
    drwxr-xr-x   4 agnt  staff      128  9 13 00:48 [1m[36mchap02[m[m/
    -rw-r--r--   1 agnt  staff     1349  9 12 17:29 developmental.md
    -rw-r--r--   1 agnt  staff      734  9 12 20:43 doc_APNs.md
    -rw-r--r--@  1 agnt  staff     3881  9 12 20:43 doc_APPLE_login_test_starpass.md
    -rw-r--r--   1 agnt  staff     1313  9 12 20:43 doc_git.md
    -rw-r--r--   1 agnt  staff     1096  9 12 20:43 doc_rest_api.md
    -rw-r--r--   1 agnt  staff    25245  9 12 20:43 doc_thread.md
    -rw-r--r--@  1 agnt  staff    11793  9 12 20:43 doc_xssFilter.md
    -rw-r--r--   1 agnt  staff      769  8  2 16:57 feed.xml
    -rw-r--r--   1 agnt  staff       42  9 13 00:54 filename
    -rw-r--r--   1 agnt  staff     6357  9 12 20:29 food.md
    -rw-r--r--   1 agnt  staff     1689  8  2 16:57 index.html
    drwxr-xr-x   3 agnt  staff       96  9 12 21:02 [1m[36mjava[m[m/
    -rw-r--r--   1 agnt  staff      492  9 13 01:30 learn.md
    -rw-r--r--   1 agnt  staff       20  9 13 00:55 numbers.json
    -rw-r--r--   1 agnt  staff     2084  9 12 17:29 parking.md
    -rw-r--r--   1 agnt  staff       65  9 13 00:54 programming.txt
    -rw-r--r--   1 agnt  staff     7237  9 12 20:43 python_coding_test.ipynb
    -rw-r--r--@  1 agnt  staff     4005  9 13 01:00 python_coding_test.md
    -rw-r--r--   1 agnt  staff   274803  9 13 01:37 python_crash_course.ipynb
    -rw-r--r--   1 agnt  staff    21203  9 12 20:43 python_ds_algorithm.ipynb
    -rw-r--r--@  1 agnt  staff    11194  9 13 01:01 python_ds_algorithm.md
    -rw-r--r--   1 agnt  staff    39142  8  2 16:57 spring-set-up.md
    -rw-r--r--   1 agnt  staff     1541  8  2 16:57 tags.html
    -rw-r--r--   1 agnt  staff       53  9 13 00:49 test.py
    -rw-r--r--   1 agnt  staff        5  9 13 00:55 username.json
    drwxr-xr-x   8 agnt  staff      256  9 13 00:10 [1m[36mvenv[m[m/


``` bash
pip3 install matplotlib
```


```python
%matplotlib inline
```


```python
"""
Simple demo of a scatter plot
"""
import numpy as np
import matplotlib.pyplot as plt

N = 50
x = np.random.rand(N)
y = np.random.rand(N)
colors = np.random.rand(N)
area = np.pi * (15 * np.random.rand(N))**2 # 0 to 15 point radiuses

plt.scatter(x, y, s=area, c=colors, alpha=0.5)
plt.show()
```


    
![png](output_212_0.png)
    



```python
%%HTML
<iframe width="560" height="315" src="https://www.youtube.com/embed/HW29067qVWk" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
```


<iframe width="560" height="315" src="https://www.youtube.com/embed/HW29067qVWk" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>




```bash
%%bash
ls -al
```

    total 16640
    drwxr-xr-x  46 agnt  staff     1472  9 13 01:37 .
    drwxr-xr-x+ 73 agnt  staff     2336  9 13 01:31 ..
    -rw-r--r--@  1 agnt  staff     6148  9 12 18:23 .DS_Store
    drwxr-xr-x  15 agnt  staff      480  9 13 01:31 .git
    -rw-r--r--   1 agnt  staff       80  9 12 20:51 .gitignore
    drwxr-xr-x   3 agnt  staff       96  9 13 00:47 .ipynb_checkpoints
    -rw-r--r--   1 agnt  staff     2936  8  2 16:57 404.html
    -rw-r--r--   1 agnt  staff     1852  9 12 20:44 CODING_TEST.sql
    -rw-r--r--   1 agnt  staff      126  8  2 16:57 Gemfile
    -rw-r--r--   1 agnt  staff     1105  8  2 16:57 LICENSE
    -rw-r--r--   1 agnt  staff     1157  8  2 16:57 README.md
    -rw-r--r--   1 agnt  staff     7110  9 12 21:00 README_java.md
    -rw-r--r--   1 agnt  staff  7989959  8  2 16:57 Semi-Project.pdf
    -rw-r--r--   1 agnt  staff      920  9 13 00:05 _config.yml
    drwxr-xr-x   7 agnt  staff      224  8  2 16:57 _includes
    drwxr-xr-x   5 agnt  staff      160  8  2 16:57 _layouts
    drwxr-xr-x   5 agnt  staff      160  9 12 20:25 _posts
    drwxr-xr-x   8 agnt  staff      256  8  2 16:57 assets
    -rw-r--r--   1 agnt  staff     1852  9 12 03:38 categories.html
    drwxr-xr-x   4 agnt  staff      128  9 13 00:48 chap02
    -rw-r--r--   1 agnt  staff     1349  9 12 17:29 developmental.md
    -rw-r--r--   1 agnt  staff      734  9 12 20:43 doc_APNs.md
    -rw-r--r--@  1 agnt  staff     3881  9 12 20:43 doc_APPLE_login_test_starpass.md
    -rw-r--r--   1 agnt  staff     1313  9 12 20:43 doc_git.md
    -rw-r--r--   1 agnt  staff     1096  9 12 20:43 doc_rest_api.md
    -rw-r--r--   1 agnt  staff    25245  9 12 20:43 doc_thread.md
    -rw-r--r--@  1 agnt  staff    11793  9 12 20:43 doc_xssFilter.md
    -rw-r--r--   1 agnt  staff      769  8  2 16:57 feed.xml
    -rw-r--r--   1 agnt  staff       42  9 13 00:54 filename
    -rw-r--r--   1 agnt  staff     6357  9 12 20:29 food.md
    -rw-r--r--   1 agnt  staff     1689  8  2 16:57 index.html
    drwxr-xr-x   3 agnt  staff       96  9 12 21:02 java
    -rw-r--r--   1 agnt  staff      492  9 13 01:30 learn.md
    -rw-r--r--   1 agnt  staff       20  9 13 00:55 numbers.json
    -rw-r--r--   1 agnt  staff     2084  9 12 17:29 parking.md
    -rw-r--r--   1 agnt  staff       65  9 13 00:54 programming.txt
    -rw-r--r--   1 agnt  staff     7237  9 12 20:43 python_coding_test.ipynb
    -rw-r--r--@  1 agnt  staff     4005  9 13 01:00 python_coding_test.md
    -rw-r--r--   1 agnt  staff   274803  9 13 01:37 python_crash_course.ipynb
    -rw-r--r--   1 agnt  staff    21203  9 12 20:43 python_ds_algorithm.ipynb
    -rw-r--r--@  1 agnt  staff    11194  9 13 01:01 python_ds_algorithm.md
    -rw-r--r--   1 agnt  staff    39142  8  2 16:57 spring-set-up.md
    -rw-r--r--   1 agnt  staff     1541  8  2 16:57 tags.html
    -rw-r--r--   1 agnt  staff       53  9 13 00:49 test.py
    -rw-r--r--   1 agnt  staff        5  9 13 00:55 username.json
    drwxr-xr-x   8 agnt  staff      256  9 13 00:10 venv



```python
%%timeit
square_evens = [n*n for n in range(100)]
```

    5.09 µs ± 58.1 ns per loop (mean ± std. dev. of 7 runs, 100000 loops each)



```python
import pandas as pd
import numpy as np

df = pd.DataFrame(np.random.randn(10, 5))
df.head()
```




<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>0</th>
      <th>1</th>
      <th>2</th>
      <th>3</th>
      <th>4</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>0.011451</td>
      <td>0.566869</td>
      <td>-0.591890</td>
      <td>0.043932</td>
      <td>2.080241</td>
    </tr>
    <tr>
      <th>1</th>
      <td>-1.314721</td>
      <td>-0.673593</td>
      <td>-0.592165</td>
      <td>-1.409841</td>
      <td>-0.845797</td>
    </tr>
    <tr>
      <th>2</th>
      <td>-0.704676</td>
      <td>-0.026795</td>
      <td>-1.152982</td>
      <td>-0.936747</td>
      <td>-2.539400</td>
    </tr>
    <tr>
      <th>3</th>
      <td>-0.653081</td>
      <td>1.963302</td>
      <td>0.581187</td>
      <td>-0.204401</td>
      <td>-0.204504</td>
    </tr>
    <tr>
      <th>4</th>
      <td>-0.884161</td>
      <td>0.285411</td>
      <td>-1.577241</td>
      <td>-1.979155</td>
      <td>-0.132490</td>
    </tr>
  </tbody>
</table>
</div>


