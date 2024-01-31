package strategy


type Divide struct {}


func NewDivide() *Divide {
	return &Divide{}
}

func (self *Divide) Execute(a int32, b int32) float64 {
	return float64(a/b)
}