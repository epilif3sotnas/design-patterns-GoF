package memento


type Snapshot struct {
	editor *Editor
	text string
	curX, curY, selectionWidth uint32
}


func NewSnapshot(editor *Editor, text string, curX, curY, selectionWidth uint32) *Snapshot {
	return &Snapshot{
		editor: editor,
		text: text,
		curX: curX,
		curY: curY,
		selectionWidth: selectionWidth,
	}
}

func (self *Snapshot) Restore() {
	self.editor.SetText(self.text)
	self.editor.SetCursor(self.curX, self.curY)
	self.editor.SetSelectionWidth(self.selectionWidth)
}

func (self *Snapshot) GetEditor() *Editor {
	return self.editor
}

func (self *Snapshot) GetText() string {
	return self.text
}

func (self *Snapshot) GetCurX() uint32 {
	return self.curX
}

func (self *Snapshot) GetCurY() uint32 {
	return self.curY
}

func (self *Snapshot) GetCursor() (uint32, uint32) {
	return self.GetCurX(), self.GetCurY()
}

func (self *Snapshot) GetSelectionWidth() uint32 {
	return self.selectionWidth
}
