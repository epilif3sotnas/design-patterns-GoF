package abstractfactory


type AbstractFactory struct {
	animal Animal
}


func NewAbstactFactory(animal Animal) *AbstractFactory {
	return &AbstractFactory{
		animal: animal,
	}
}

func (self *AbstractFactory) GetAnimalInfo() string {
	return "Name: " + self.animal.GetName() + "\nSpecie Name: " + self.animal.GetSpecieName()
}