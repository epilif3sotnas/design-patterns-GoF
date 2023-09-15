package command


type PasteCommand struct {
	CommandAbstract
}


func NewPasteCommand(app *Application, editor *Editor) *PasteCommand {
	return &PasteCommand{
		CommandAbstract: CommandAbstract{
			app: app,
			editor: editor,
		},
	}
}

func (self *PasteCommand) Execute() bool {
	self.SaveBackup()
	self.editor.ReplaceSelection(self.app.clipboard)
	return true
}