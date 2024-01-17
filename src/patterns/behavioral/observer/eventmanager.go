package observer


type EventManager struct {
	listeners map[EventType][]EventListener
}


func NewEventManager() *EventManager {
	return &EventManager{
		listeners: make(map[EventType][]EventListener),
	}
}

func (self *EventManager) Subscribe(eventType EventType, listener EventListener) {
	self.listeners[eventType] = append(self.listeners[eventType], listener)
}

func (self *EventManager) UnSubscribe(eventType EventType, listener EventListener) {
	for idx, element := range self.listeners[eventType] {
		if element != listener { continue }

		self.listeners[eventType] = append(self.listeners[eventType][:idx], self.listeners[eventType][idx+1:]...)
	}
}

func (self *EventManager) Notify(eventType EventType, data string) {
	for _, element := range self.listeners[eventType] {
		element.Update(data)
	}
}

func (self *EventManager) GetListeners() map[EventType][]EventListener {
	return self.listeners
}

func (self *EventManager) SetListeners(listeners map[EventType][]EventListener) {
	self.listeners = listeners
}