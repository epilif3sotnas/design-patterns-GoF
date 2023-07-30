package abstractfactory


type Cat struct {
	name string
	specieName string
}


func NewCat(name string, specieName string) *Cat {
	return &Cat{
		name: name,
		specieName: specieName,
	}
}

func (self *Cat) GetName() string {
	return self.name
}

func (self *Cat) SetName(name string) {
	self.name = name
}

func (self *Cat) GetSpecieName() string {
	return self.specieName
}

func (self *Cat) SetSpecieName(specieName string) {
	self.specieName = specieName
}