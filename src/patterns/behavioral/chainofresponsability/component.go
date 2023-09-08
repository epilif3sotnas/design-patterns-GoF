package chainofresponsability


// std
import (
	"fmt"
)


type Component struct {
	container *Container
	tooltipText string
}


func NewComponent(container *Container, tooltipText string) *Component {
	return &Component{
		container: container,
		tooltipText: tooltipText,
	}
}

func (self *Component) ShowHelp() {
	if self.tooltipText == "" {
		self.container.ShowHelp()
		return
	}

	fmt.Println("Tool Tip: " + self.tooltipText)
}

func (self *Component) GetContainer() *Container {
	return self.container
}

func (self *Component) SetContainer(container *Container) {
	self.container = container
}

func (self *Component) GetTooltipText() string {
	return self.tooltipText
}

func (self *Component) SetTooltipText(tooltipText string) {
	self.tooltipText = tooltipText
}