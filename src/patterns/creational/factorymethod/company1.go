package factorymethod


type Company1 struct {
	name string
	iban string
}


func NewCompany1(name string, iban string) *Company1 {
	return &Company1{
		name: name,
		iban:iban,
	}
}

func (self *Company1) SendMoney(ibanToSend string) bool {
	return true
}

func (self *Company1) GetCompanyInfo() string {
	return "Company Name: " + self.name + "\nIBAN: " + self.iban
}

func (self *Company1) GetName() string {
	return self.name
}

func (self *Company1) SetName(name string) {
	self.name = name
}

func (self *Company1) GetIban() string {
	return self.iban
}

func (self *Company1) SetIban(iban string) {
	self.iban = iban
}