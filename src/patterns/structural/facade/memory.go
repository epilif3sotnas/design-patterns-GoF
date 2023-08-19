package facade


// std
import (
	"fmt"
)

type Memory struct {}


func NewMemory() *Memory {
	return &Memory{}
}

func (self *Memory) Load(position int64, data []byte) {
	fmt.Println("Memory Load: position ", position, " data ", string(data))
}