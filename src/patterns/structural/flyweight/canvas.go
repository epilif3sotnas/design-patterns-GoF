package flyweight


import (
	"fmt"
)


type Canvas struct {
	width float32
	height float32
}


func NewCanvas(width, height float32) *Canvas {
	return &Canvas{
		width: width,
		height: height,
	}
}

func (self *Canvas) Draw(x, y float32) {
	fmt.Println(
		"\nDrawing:\nX: ", x,
					"\nY: ", y,
					"\nWidth: ", self.width,
					"\nHeight: ", self.height,
		)
}