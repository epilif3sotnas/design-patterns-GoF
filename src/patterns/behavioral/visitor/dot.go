package visitor


import (
    // std
    "fmt"
)


type Dot struct {
    Shape
    x, y int32
}


func NewDot(x, y int32) *Dot {
    return &Dot{
        x: x,
        y: y,
    }
}

func (self *Dot) Move(x, y int32) {
    self.x, self.y = x, y
}

func (self *Dot) Draw() {
    fmt.Println("Dot draw with (x,y) => (", self.x, ",", self.y, ")")
}

func (self *Dot) Accept(visitor Visitor) {
    visitor.VisitDot(self)
}

func (self *Dot) GetX() int32  {
    return self.x
}

func (self *Dot) SetX(x int32) {
    self.x = x
}

func (self *Dot) GetY() int32 {
    return self.y
}

func (self *Dot) SetY(y int32) {
    self.y = y
}
