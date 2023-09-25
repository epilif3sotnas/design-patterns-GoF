package iterator


// std
import (
	"fmt"
)


type SocialSpammer struct {}


func NewSocialSpammer() *SocialSpammer {
	return &SocialSpammer{}
}

func (self *SocialSpammer) Send(iterator ProfileIterator, message string) {
	for iterator.HasMore() {
		var profile *Profile = iterator.GetNext()

		fmt.Println("Email sent to", profile.GetEmail(), "with a message:", message)
	}
}