package mediator


type Button struct {
	Component
}


func NewButton(dialog Mediator) *Button {
	return &Button{
		Component {
			componentType: BUTTON,
			dialog: dialog,
		},
	}
}