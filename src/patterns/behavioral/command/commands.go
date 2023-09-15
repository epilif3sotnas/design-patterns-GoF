package command


type Commands int


const (
	COPY Commands = iota
	CUT
	PASTE
	UNDO
)