package factorymethod


type Creator struct {}


func NewCreator() *Creator {
	return &Creator{}
}

func (self *Creator) Create() Company {
	return NewCompany1("Klarna", "SE4550000000058398257466")
}