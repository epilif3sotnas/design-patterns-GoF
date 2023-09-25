package iterator


type Profile struct {
	id string
	email string
}


func NewProfile(id string, email string) *Profile {
	return &Profile{
		id: id,
		email: email,
	}
}

func (self *Profile) GetId() string {
	return self.id
}

func (self *Profile) SetId(id string) {
	self.id = id
}

func (self *Profile) GetEmail() string {
	return self.email
}

func (self *Profile) SetEmail(email string) {
	self.email = email
}