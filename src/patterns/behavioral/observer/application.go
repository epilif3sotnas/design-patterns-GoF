package observer


import(
	// std
	"fmt"
)


type Application struct {
	editor *Editor
}


func NewApplication(editor *Editor) *Application {
	return &Application{
		editor: editor,
	}
}

func (self *Application) Config() {
	self.editor.GetEventManager().Subscribe(
		EMAIL,
		NewEmailAlertsListener(),
	)

	self.editor.GetEventManager().Subscribe(
		LOGGING,
		NewLoggingListener(),
	)
}

func (self *Application) Run() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		self.editor.OpenFile()
		self.editor.SaveFile()
		self.editor.CloseFile()
		self.editor.RemoveFile()

		fmt.Println()
	}
}