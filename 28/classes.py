class Animal:
    def sound(self):
        print("Animal sound")

class Dog(Animal):
    def __init__(self, name, age, color, breed):
        self.name = name
        self.age = age
        self.color = color
        self.breed = breed
    def __private_sound(self):
        print("Dog sounds: Woof Woof")
    def sound(self):
        self.__private_sound()
    def __str__(self):
        return f'Dog:(name={self.name}, age={self.age}, color={self.color}, breed={self.breed})'
    def __repr__(self):
        return f'Dog:(name={self.name!r}, age={self.age!r}, color={self.color!r}, breed={self.breed!r})'



# Define new objects
animal = Animal()
dog = Dog("Cupcake", 3, "Black", "Rottweiler")


# Call the sound method
animal.sound()
dog.sound()
# dog.__private_sound() # AttributeError

print(str(dog))
print(repr(dog))