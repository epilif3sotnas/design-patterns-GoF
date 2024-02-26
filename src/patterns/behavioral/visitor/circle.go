package visitor


import (
    // std
    "fmt"
)


type Circle struct {
    Shape
    x, y int32
}


func NewCircle(x, y int32) *Circle {
    return &Circle{
        x: x,
        y: y,
    }
}

func (self *Circle) Move(x, y int32) {
    self.x, self.y = x, y
}

func (self *Circle) Draw() {
    fmt.Println("Circle draw with (x,y) => (", self.x, ",", self.y, ")")
}

func (self *Circle) Accept(visitor Visitor) {
    visitor.VisitCircle(self)
}

func (self *Circle) GetX() int32  {
    return self.x
}

func (self *Circle) SetX(x int32) {
    self.x = x
}

func (self *Circle) GetY() int32 {
    return self.y
}

func (self *Circle) SetY(y int32) {
    self.y = y
}
