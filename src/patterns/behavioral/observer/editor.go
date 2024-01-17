package observer


type Editor struct {
	eventManager *EventManager
}


func NewEditor(eventManager *EventManager) *Editor {
	return &Editor{
		eventManager: eventManager,
	}
}

func (self *Editor) OpenFile() {
	self.notifyEvents("Open file")
}

func (self *Editor) SaveFile() {
	self.notifyEvents("Save file")
}

func (self *Editor) RemoveFile() {
	self.notifyEvents("Remove file")
}

func (self *Editor) CloseFile() {
	self.notifyEvents("Close file")
}

func (self *Editor) notifyEvents(data string) {
	self.eventManager.Notify(EMAIL, data)
	self.eventManager.Notify(LOGGING, data)
}

func (self *Editor) GetEventManager() *EventManager {
	return self.eventManager
}

func (self *Editor) SetEventManager(eventManager *EventManager) {
	self.eventManager = eventManager
}