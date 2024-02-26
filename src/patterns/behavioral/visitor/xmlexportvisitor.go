package visitor


type XmlExportVisitor struct{
    Visitor
}


func NewXmlExportVisitor() *XmlExportVisitor {
    return &XmlExportVisitor{}
}

func (self *XmlExportVisitor) VisitDot(dot *Dot) {
    dot.Draw()
}

func (self *XmlExportVisitor) VisitCircle(circle *Circle) {
    circle.Draw()
}

func (self *XmlExportVisitor) VisitRectangle(rectangle *Rectangle) {
    rectangle.Draw()
}

func (self *XmlExportVisitor) VisitCompoundShape(compoundShape *CompoundShape) {
    compoundShape.Draw()
}
