package composite


type CoumpoundGraphic struct {
	children []Graphic
}


func NewCoumpoundGraphic() *CoumpoundGraphic {
	return &CoumpoundGraphic{}
}

func (self *CoumpoundGraphic) AddChild(child Graphic) {
	self.children = append(self.children, child)
}

func (self *CoumpoundGraphic) Move(x,y float32) bool {
	for _, child := range self.children {
		child.Move(x,y)
	}

	return true
}

func (self *CoumpoundGraphic) Draw() bool {
	for _, child := range self.children {
		child.Draw()
	}

	return true
}