package interpreter


type Expression interface {
	Interpreter(string) bool
}