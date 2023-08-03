package composite


import(
	"fmt"
)


type Dot struct {
	x float32
	y float32
}


func NewDot(x, y float32) *Dot {
	return &Dot{x: x, y: y}
}

func (self *Dot) Move(x, y float32) bool {
	fmt.Println("\nDot Moved\nFrom\nx: ", self.x, "\ny: ", self.y, "\nTo\nx: ", x, "\ny: ", y)

    self.x = x
    self.y = y
    
    return true
}

func (self *Dot) Draw() bool {
	fmt.Println("\nDot Drawn\nx: ", self.x, "\ny: ", self.y)
    
    return true
}

func (self *Dot) GetX() float32 {
	return self.x
}

func (self *Dot) SetX(x float32) {
	self.x = x
}

func (self *Dot) GetY() float32 {
	return self.y
}

func (self *Dot) SetY(y float32) {
	self.y = y
}