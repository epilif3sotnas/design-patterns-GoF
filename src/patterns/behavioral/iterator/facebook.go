package iterator


type Facebook struct {}


func NewFacebook() *Facebook {
	return &Facebook{}
}

func (self *Facebook) CreateFriendsIterator(profileId string) ProfileIterator {
	return NewFacebookIterator(self, profileId, "friends")
}

func (self *Facebook) CreateCoworkersIterator(profileId string) ProfileIterator {
	return NewFacebookIterator(self, profileId, "coworkers")
}

func (self *Facebook) GetSocialGraph(profileId, type_ string) []*Profile {
	return []*Profile{
		NewProfile("1", "email1@example.com"),
		NewProfile("2", "email2@example.com"),
		NewProfile("3", "email3@example.com"),
		NewProfile("4", "email4@example.com"),
		NewProfile("5", "email5@example.com"),
		NewProfile("6", "email6@example.com"),
		NewProfile("7", "email7@example.com"),
		NewProfile("8", "email8@example.com"),
		NewProfile("9", "email9@example.com"),
	}
}