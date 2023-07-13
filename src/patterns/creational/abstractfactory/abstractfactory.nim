# internal
import
    patterns/creational/abstractfactory/animal


type
    AbstactFactory*[TAnimal: Animal] = ref object
        animal: TAnimal


proc newAbstactFactory*[T: Animal](animal: T): AbstactFactory[T] =
    return AbstactFactory[T](animal: animal)

proc getAnimalInfo*(self: AbstactFactory): string =
    return "Name: " & self.animal.getName() & "\nSpecie Name: " & self.animal.getSpecieName()