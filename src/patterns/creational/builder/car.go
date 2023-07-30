package builder


import(
	"strconv"
)


type Car struct {
	engine string
	horsepower uint16
}


func NewCar(engine string, horsepower uint16) *Car {
	return &Car{
		engine: engine,
		horsepower: horsepower,
	}
}

func (self *Car) GetCarInfo() string {
	return "Engine: " + self.engine + "\nHorsepower: " + strconv.FormatUint(uint64(self.horsepower), 10)
}

func (self *Car) GetEngine() string {
	return self.engine
}

func (self *Car) SetEngine(engine string) {
	self.engine = engine
}

func (self *Car) GetHorsepower() uint16 {
	return self.horsepower
}

func (self *Car) SetHorsepower(horsepower uint16) {
	self.horsepower = horsepower
}