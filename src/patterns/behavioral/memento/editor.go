package memento


type Editor struct {
	text string
	curX, curY, selectionWidth uint32
}


func NewEditor() *Editor {
	return &Editor{
		text: "",
		curX: 0,
		curY: 0,
		selectionWidth: 0,
	}
}

func (self *Editor) CreateSnapshot() *Snapshot {
	return NewSnapshot(
		self,
		self.text,
		self.curX,
		self.curY,
		self.selectionWidth,
	)
}

func (self *Editor) GetText() string {
	return self.text
}

func (self *Editor) SetText(text string) {
	self.text = text
}

func (self *Editor) GetCurX() uint32 {
	return self.curX
}

func (self *Editor) SetCurX(curX uint32) {
	self.curX = curX
}

func (self *Editor) GetCurY() uint32 {
	return self.curY
}

func (self *Editor) SetCurY(curY uint32) {
	self.curY = curY
}

func (self *Editor) GetCursor() (uint32, uint32) {
	return self.GetCurX(), self.GetCurY()
}

func (self *Editor) SetCursor(curX, curY uint32) {
	self.SetCurX(curX)
	self.SetCurY(curY)
}

func (self *Editor) GetSelectionWidth() uint32 {
	return self.selectionWidth
}

func (self *Editor) SetSelectionWidth(selectionWidth uint32) {
	self.selectionWidth = selectionWidth
}