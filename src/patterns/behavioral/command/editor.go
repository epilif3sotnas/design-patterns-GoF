package command


type Editor struct {
	text string
}


func NewEditor(text string) *Editor {
	return &Editor{
		text: text,
	}
}

func (self *Editor) GetSelection() string {
	return self.text
}

func (self *Editor) DeleteSelection() {
	self.text = ""
}

func (self *Editor) ReplaceSelection(text string) {
	self.text = text
}