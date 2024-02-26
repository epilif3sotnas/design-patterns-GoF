package visitor


type Visitor interface {
    VisitDot(dot *Dot)
    VisitCircle(circle *Circle)
    VisitRectangle(rectangle *Rectangle)
    VisitCompoundShape(compoundShape *CompoundShape)
}
