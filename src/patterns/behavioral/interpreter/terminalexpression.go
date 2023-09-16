package interpreter


// std
import (
	"strings"
)


type TerminalExpression struct {
	data string
}


func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{
		data: data,
	}
}

func (self *TerminalExpression) Interpreter(context string) bool {
	if (!strings.Contains(context, self.data)) {
		return false
	}

	return true
}