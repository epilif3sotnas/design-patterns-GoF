package command


type CutCommand struct {
	CommandAbstract
}


func NewCutCommand(app *Application, editor *Editor) *CutCommand {
	return &CutCommand{
		CommandAbstract: CommandAbstract{
			app: app,
			editor: editor,
		},
	}
}

func (self *CutCommand) Execute() bool {
	self.SaveBackup()
	self.app.clipboard = self.editor.GetSelection()
	self.editor.DeleteSelection()
	return true
}