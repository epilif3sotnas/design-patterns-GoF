package command


// std
import (
	"fmt"
)


type Application struct {
	clipboard string
	editors []*Editor
	activeEditor *Editor
	history *CommandHistory
	commands []Command
}


func NewApplication(
	clipboard string,
	editors []*Editor,
	activeEditor *Editor,
	history *CommandHistory,
) *Application {
	return &Application{
		clipboard: clipboard,
		editors: editors,
		activeEditor: activeEditor,
		history: history,
	}
}

func (self *Application) InitCommands() {
	self.commands = append(self.commands, NewCopyCommand(self, self.activeEditor))
	self.commands = append(self.commands, NewCutCommand(self, self.activeEditor))
	self.commands = append(self.commands, NewPasteCommand(self, self.activeEditor))
	self.commands = append(self.commands, NewUndoCommand(self, self.activeEditor))
}

func (self *Application) ExecuteCommand(command Commands) {
	switch command {
	case COPY, CUT, PASTE:
		if (!self.commands[command].Execute()) {
			return
		}
	
		self.history.Push(self.commands[command])
	
	case UNDO:
		self.history.Pop().Undo()
	
	default:
		fmt.Println("Command not supported: ", command)
	}
}

func (self *Application) GetClipboard() string {
	return self.clipboard
}

func (self *Application) SetClipboard(clipboard string) {
	self.clipboard = clipboard
}

func (self *Application) GetEditors() []*Editor {
	return self.editors
}

func (self *Application) SetEditors(editors []*Editor) {
	self.editors = editors
}

func (self *Application) GetActiveEditor() *Editor {
	return self.activeEditor
}

func (self *Application) SetActiveEditor(activeEditor *Editor) {
	self.activeEditor = activeEditor
}

func (self *Application) GetHistory() *CommandHistory {
	return self.history
}

func (self *Application) SetHistory(history *CommandHistory) {
	self.history = history
}