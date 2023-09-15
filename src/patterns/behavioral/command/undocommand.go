package command


type UndoCommand struct {
	CommandAbstract
}


func NewUndoCommand(app *Application, editor *Editor) *UndoCommand {
	return &UndoCommand{
		CommandAbstract: CommandAbstract{
			app: app,
			editor: editor,
		},
	}
}

func (self *UndoCommand) Execute() bool {
	self.app.ExecuteCommand(UNDO)
	return false
}