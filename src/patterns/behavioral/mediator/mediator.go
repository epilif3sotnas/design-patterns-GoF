package mediator


type Mediator interface {
	Notify(sender *Component, event string)
}