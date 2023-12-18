package mediator


type CheckBox struct {
	Component
	checked bool
}


func NewCheckBox(dialog Mediator) *CheckBox {
	return &CheckBox{
		checked: true,
		Component: Component {
			componentType: CHECKBOX,
			dialog: dialog,
		},
	}
}

func (self *Component) Check() {
	self.dialog.Notify(self, "check")
}

func (self *CheckBox) IsChecked() bool {
	return self.checked
}

func (self *CheckBox) SetChecked(checked bool) {
	self.checked = checked
}