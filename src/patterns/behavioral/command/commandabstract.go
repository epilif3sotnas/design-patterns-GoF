package command


type CommandAbstract struct {
	app *Application
	editor *Editor
	backup string
}


func NewCommandAbstract(app *Application, editor *Editor) *CommandAbstract {
	return &CommandAbstract{
		app: app,
		editor: editor,
	}
}

func (self *CommandAbstract) SaveBackup() {
	self.backup = self.editor.GetSelection()
}

func (self *CommandAbstract) Undo() {
	self.editor.ReplaceSelection(self.backup)
}

func (self *CommandAbstract) GetApp() *Application {
	return self.app
}

func (self *CommandAbstract) SetApp(app *Application) {
	self.app = app
}

func (self *CommandAbstract) GetEditor() *Editor {
	return self.editor
}

func (self *CommandAbstract) SetEditor(editor *Editor) {
	self.editor = editor
}

func (self *CommandAbstract) GetBackup() string {
	return self.backup
}

func (self *CommandAbstract) SetBackup(backup string) {
	self.backup = backup
}