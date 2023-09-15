package command


type CopyCommand struct {
	CommandAbstract
}


func NewCopyCommand(app *Application, editor *Editor) *CopyCommand {
	return &CopyCommand{
		CommandAbstract: CommandAbstract{
			app: app,
			editor: editor,
		},
	}
}

func (self *CopyCommand) Execute() bool {
	self.app.clipboard = self.editor.GetSelection()
	return false
}