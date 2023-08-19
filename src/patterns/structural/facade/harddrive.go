package facade


// std
import (
	"fmt"
)


type HardDrive struct {}


func NewHardDrive() *HardDrive {
	return &HardDrive{}
}

func (self *HardDrive) Read(lba int64, size int) []byte {
	fmt.Println("Hard Drive Read: lba ", lba, " size ", size)
	return []byte("Init command")
}