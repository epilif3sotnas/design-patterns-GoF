package visitor


type Shape interface {
    Move(x, y int32)
    Draw()
    Accept(visitor Visitor)
}
