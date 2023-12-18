package mediator


type TextBox struct {
	Component
}


func NewTextBox(dialog Mediator) *TextBox {
	return &TextBox{
		Component {
			componentType: TEXTBOX,
			dialog: dialog,
		},
	}
}