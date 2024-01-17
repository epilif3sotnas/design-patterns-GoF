package observer


import (
	// std
	"fmt"
)

type LoggingListener struct {
	message string
}


func NewLoggingListener() *LoggingListener {
	return &LoggingListener{
		message: "",
	}
}

func (self *LoggingListener) Update(message string) {
	if message == "" { return }

	fmt.Println("Updating logging message -> FROM ", self.message, " TO ", message)

	self.SetMessage(message)
}

func (self *LoggingListener) GetMessage() string {
	return self.message
}

func (self *LoggingListener) SetMessage(message string) {
	self.message = message
}