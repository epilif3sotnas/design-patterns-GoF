package observer


import (
	// std
	"fmt"
)

type EmailAlertsListener struct {
	message string
}


func NewEmailAlertsListener() *EmailAlertsListener {
	return &EmailAlertsListener{
		message: "",
	}
}

func (self *EmailAlertsListener) Update(message string) {
	if message == "" { return }

	fmt.Println("Updating alert message -> FROM ", self.message, " TO ", message)

	self.SetMessage(message)
}

func (self *EmailAlertsListener) GetMessage() string {
	return self.message
}

func (self *EmailAlertsListener) SetMessage(message string) {
	self.message = message
}