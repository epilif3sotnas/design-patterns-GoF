package iterator


type FacebookIterator struct {
	facebook *Facebook
	profileId string
	type_ string
	currentPosition int
	cache []*Profile
}


func NewFacebookIterator(facebook *Facebook, profileId string, type_ string) *FacebookIterator {
	return &FacebookIterator{
		facebook: facebook,
		profileId: profileId,
		type_: type_,
	}
}

func (self *FacebookIterator) lazyInit() {
	if self.cache == nil {
		self.cache = self.facebook.GetSocialGraph(self.profileId, self.type_)
	}
}

func (self *FacebookIterator) GetNext() *Profile {
	if !self.HasMore() {
		return nil
	}

	var nextElement *Profile = self.cache[self.currentPosition]
	self.currentPosition++
	return nextElement
}

func (self *FacebookIterator) HasMore() bool {
	self.lazyInit()
	return self.currentPosition < len(self.cache)
}

func (self *FacebookIterator) GetFacebook() *Facebook {
	return self.facebook
}

func (self *FacebookIterator) SetFacebook(facebook *Facebook) {
	self.facebook = facebook
}

func (self *FacebookIterator) GetProfileId() string {
	return self.profileId
}

func (self *FacebookIterator) SetProfileId(profileId string) {
	self.profileId = profileId
}

func (self *FacebookIterator) GetType() string {
	return self.type_
}

func (self *FacebookIterator) SetType(type_ string) {
	self.type_ = type_
}

func (self *FacebookIterator) GetCurrentPosition() int {
	return self.currentPosition
}

func (self *FacebookIterator) GetCache() []*Profile {
	return self.cache
}

func (self *FacebookIterator) SetCache(cache []*Profile) {
	self.cache = cache
}