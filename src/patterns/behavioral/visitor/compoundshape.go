package visitor


import (
    // std
    "fmt"
)


type CompoundShape struct {
    Shape
    x, y int32
}


func NewCompoundShape(x, y int32) *CompoundShape {
    return &CompoundShape{
        x: x,
        y: y,
    }
}

func (self *CompoundShape) Move(x, y int32) {
    self.x, self.y = x, y
}

func (self *CompoundShape) Draw() {
    fmt.Println("CompoundShape draw with (x,y) => (", self.x, ",", self.y, ")")
}

func (self *CompoundShape) Accept(visitor Visitor) {
    visitor.VisitCompoundShape(self)
}

func (self *CompoundShape) GetX() int32  {
    return self.x
}

func (self *CompoundShape) SetX(x int32) {
    self.x = x
}

func (self *CompoundShape) GetY() int32 {
    return self.y
}

func (self *CompoundShape) SetY(y int32) {
    self.y = y
}
