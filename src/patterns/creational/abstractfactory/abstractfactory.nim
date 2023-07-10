# internal
import
    patterns/creational/abstractfactory/animal


type
    AbstactFactory* = ref object

# type
#     AbstactFactory* = ref object
#         animal: Animal


proc newAbstactFactory*(): AbstactFactory =
    return AbstactFactory()

# proc newAbstactFactory*(animal: Animal): AbstactFactory =
#     # return AbstactFactory(animal: animal)

proc getAnimalInfo*(self: AbstactFactory, animal: Animal): string =
    return "Name: " & animal.getName() & "\nSpecie Name: " & animal.getSpecieName()

# proc getAnimalInfo*(self: AbstactFactory): string =
#     return "Name: " & self.animal.getName() & "\nSpecie Name: " & self.animal.getSpecieName()

# proc getAnimal*(self: Animal): Animal =
#     return self.animal

# proc setAnimal*(self: Animal, animal: Animal) =
#     self.animal = animal