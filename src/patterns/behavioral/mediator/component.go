package mediator


type Component struct {
	componentType ComponentType
	dialog Mediator
}


func NewComponent(dialog Mediator) *Component {
	return &Component{
		dialog: dialog,
	}
}


func (self *Component) Click() {
	self.dialog.Notify(self, "click")
}

func (self *Component) KeyPress() {
	self.dialog.Notify(self, "keypress")
}

func (self *Component) GetDialog() Mediator {
	return self.dialog
}

func (self *Component) SetDialog(dialog Mediator) {
	self.dialog = dialog
}

func (self *Component) GetComponentType() ComponentType {
	return self.componentType
}

func (self *Component) SetComponentType(componentType ComponentType) {
	self.componentType = componentType
}