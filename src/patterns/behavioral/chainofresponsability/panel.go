package chainofresponsability


// std
import (
	"fmt"
)


type Panel struct {
	Container
	modalHelpText string
}


func NewPanel(
	container *Container,
	tooltipText string,
	children []*Component,
	modalHelpText string,
) *Panel {
	return &Panel{
		Container: Container{
			Component: Component{
				container: container,
				tooltipText: tooltipText,
			},
			children: children,
		},
		modalHelpText: modalHelpText,
	}
}

func (self *Panel) ShowHelp() {
	if self.modalHelpText == "" {
		self.Component.ShowHelp()
		return
	}

	fmt.Println("Modal Help: " + self.modalHelpText)
}

func (self *Panel) GetModalHelpText() string {
	return self.modalHelpText
}

func (self *Panel) SetModalHelpText(modalHelpText string) {
	self.modalHelpText = modalHelpText
}