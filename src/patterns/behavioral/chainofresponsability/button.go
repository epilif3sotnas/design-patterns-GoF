package chainofresponsability


type Button struct {
	Component
}


func NewButton(container *Container, tooltipText string) *Button {
	return &Button{
		Component{
			container: container,
			tooltipText: tooltipText,
		},
	}
}