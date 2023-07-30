package singleton


type CarStorage struct {
	cars []string
}


func newCarStorage() *CarStorage {
	return &CarStorage{
		cars: []string{},
	}
}


var instance *CarStorage


func GetInstance() *CarStorage {
	if instance == nil {
		instance = newCarStorage()
		return instance
	}
	return instance
}

func (self *CarStorage) GetCars() []string {
	return self.cars
}

func (self *CarStorage) AddCar(car string) {
	self.cars = append(self.cars, car)
}