greeting = "Hello"
name = "John"

formal_greeting = greeting + ", " + name + "!"
print(formal_greeting)

format_greeting = f"{greeting}, {name}!"
print(format_greeting)

formatted_greeting = "{}, {}!".format(greeting, name)
print(formatted_greeting)

greet = "Hello, {}!"

print(greet.format(name))



repeat = "Hello " * 3
print(repeat)

# A box 
print("*" * 50) # Top
print("*" + " Some text ".rjust(48) + "*")
print("*" + " Some text ".center(48).upper() + "*")
print("*" * 50) # Bottom

hello_world = "Hello World"
print(hello_world.replace("World", "Python"))

# if, elif, else
for i in range(100):
    if i % 2 == 0:
        print(i," is even")
    elif i % 2 == 1:
        print(i," is odd")
    else:
        print(i," is not a number")
    

for i in range(100):
    if i % 2 == 0:
        continue
    print(i, " is odd")

# for loop for list
fruits = ["apple", "banana", "cherry", "orange", "kiwi", "melon", "mango"]

for fruit in fruits:
    print(fruit)


for fruit in fruits:
    print("")
    for char in fruit:
        print(char)