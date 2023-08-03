package composite


import(
	"fmt"
)


type Circle struct {
	*Dot
	radius float32
}


func NewCircle(x, y, radius float32) *Circle {
	return &Circle{
		Dot: &Dot{x: x, y: y},
		radius: radius,
	}
}

func (self *Circle) Move(x, y float32) bool {
	fmt.Println("\nCircle Moved\nFrom\nx: ", self.x, "\ny: ", self.y, "\nTo\nx: ", x, "\ny: ", y)

    self.x = x
    self.y = y
    
    return true
}

func (self *Circle) Draw() bool {
	fmt.Println("\nCircle Drawn\nx: ", self.x, "\ny: ", self.y, "\nradius: ", self.radius)
    
    return true
}

func (self *Circle) GetRadius() float32 {
	return self.radius
}

func (self *Circle) SetRadius(radius float32) {
	self.radius = radius
}