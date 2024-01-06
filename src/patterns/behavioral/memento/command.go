package memento


type Command struct {
	backup *Snapshot
}


func NewCommand() *Command {
	return &Command{
		backup: nil,
	}
}

func (self *Command) MakeBackup(editor *Editor) {
	self.backup = editor.CreateSnapshot()
}

func (self *Command) CleanBackup() {
	self.backup = nil
}

func (self *Command) Undo() {
	if self.backup != nil {
		self.backup.Restore()

		self.CleanBackup()
	}
}