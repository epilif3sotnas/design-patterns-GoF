package interpreter


type AndExpression struct {
	expr1 Expression
	expr2 Expression
}


func NewAndExpression(expr1, expr2 Expression) *AndExpression {
	return &AndExpression{
		expr1: expr1,
		expr2: expr2,
	}
}

func (self *AndExpression) Interpreter(context string) bool {
	return self.expr1.Interpreter(context) && self.expr2.Interpreter(context)
}