package facade


// std
import (
	"fmt"
)


type Cpu struct {}


func NewCpu() *Cpu {
	return &Cpu{}
}

func (self *Cpu) Freeze() {
	fmt.Println("CPU Freeze")
}

func (self *Cpu) Jump(position int64) {
	fmt.Println("CPU Jump to position: ", position)
}

func (self *Cpu) Execute() {
	fmt.Println("CPU Execute")
}