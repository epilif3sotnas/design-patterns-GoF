package interpreter


type OrExpression struct {
	expr1 Expression
	expr2 Expression
}


func NewOrExpression(expr1, expr2 Expression) *OrExpression {
	return &OrExpression{
		expr1: expr1,
		expr2: expr2,
	}
}

func (self *OrExpression) Interpreter(context string) bool {
	return self.expr1.Interpreter(context) || self.expr2.Interpreter(context)
}