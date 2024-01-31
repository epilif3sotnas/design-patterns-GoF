package strategy


// std
import (
	"fmt"
)


type Context struct {
	strategy Strategy
}


func NewContext(strategy Strategy) *Context {
	return &Context{
		strategy: strategy,
	}
}

func (self *Context) ExecuteOperation(a int32, b int32) {
	fmt.Println("Strategy from", a, b, "returned", self.strategy.Execute(a, b))
}

func (self *Context) GetStrategy() Strategy {
	return self.strategy
}

func (self *Context) SetStrategy(strategy Strategy) {
	self.strategy = strategy
}