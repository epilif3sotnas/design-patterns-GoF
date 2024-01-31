package strategy


type Multiply struct {}


func NewMultiply() *Multiply {
	return &Multiply{}
}

func (self *Multiply) Execute(a int32, b int32) float64 {
	return float64(a*b)
}