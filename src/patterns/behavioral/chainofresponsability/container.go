package chainofresponsability


type Container struct {
	Component
	children []*Component
}


func NewContainer(container *Container, tooltipText string, children []*Component) *Container {
	return &Container{
		Component: Component{
			container: container,
			tooltipText: tooltipText,
		},
		children: children,
	}
}

func (self *Container) Add(child *Component) {
	self.children = append(self.children, child)

	child.container = self
}

func (self *Container) GetChildren() []*Component {
	return self.children
}

func (self *Container) SetChildren(children []*Component) {
	self.children = children
}