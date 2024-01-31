package strategy


type Subtract struct {}


func NewSubtract() *Subtract {
	return &Subtract{}
}

func (self *Subtract) Execute(a int32, b int32) float64 {
	return float64(a-b)
}