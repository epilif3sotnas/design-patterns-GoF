package factorymethod


type Company interface {
	GetName() string
	SendMoney(string) bool
	GetCompanyInfo() string
}