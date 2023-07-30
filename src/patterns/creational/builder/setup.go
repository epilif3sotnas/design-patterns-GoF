package builder


type Setup struct {}


func NewSetup() *Setup {
	return &Setup{}
}

func (self *Setup) BuildCar() *Car {
	return NewCar("Toyota", 184)
}