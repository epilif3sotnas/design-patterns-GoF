package command


type Command interface {
	SaveBackup()
	Undo()
	Execute() bool
}