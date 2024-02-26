package visitor


import (
    // std
    "fmt"
)


type Rectangle struct {
    Shape
    x, y int32
}


func NewRectangle(x, y int32) *Rectangle {
    return &Rectangle{
        x: x,
        y: y,
    }
}

func (self *Rectangle) Move(x, y int32) {
    self.x, self.y = x, y
}

func (self *Rectangle) Draw() {
    fmt.Println("Rectangle draw with (x,y) => (", self.x, ",", self.y, ")")
}

func (self *Rectangle) Accept(visitor Visitor) {
    visitor.VisitRectangle(self)
}

func (self *Rectangle) GetX() int32  {
    return self.x
}

func (self *Rectangle) SetX(x int32) {
    self.x = x
}

func (self *Rectangle) GetY() int32 {
    return self.y
}

func (self *Rectangle) SetY(y int32) {
    self.y = y
}
