package strategy


type Add struct {}


func NewAdd() *Add {
	return &Add{}
}

func (self *Add) Execute(a int32, b int32) float64 {
	return float64(a+b)
}