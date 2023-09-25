package iterator


type Application struct {
	spammer *SocialSpammer
	network SocialNetwork
}


func NewApplication(spammer *SocialSpammer, network SocialNetwork) *Application {
	return &Application{
		spammer: spammer,
		network: network,
	}
}

func (self *Application) SendSpamToFriends(profile *Profile) {
	iterator := self.network.CreateFriendsIterator(profile.GetId())
	self.spammer.Send(iterator, "Hi Friends!")
}

func (self *Application) SendSpamToCoworkers(profile *Profile) {
	iterator := self.network.CreateCoworkersIterator(profile.GetId())
	self.spammer.Send(iterator, "Hi Coworkers!")
}

func (self *Application) GetSpammer() *SocialSpammer {
	return self.spammer
}

func (self *Application) SetSpammer(spammer *SocialSpammer) {
	self.spammer = spammer
}

func (self *Application) GetNetwork() SocialNetwork {
	return self.network
}

func (self *Application) SetNetwork(network SocialNetwork) {
	self.network = network
}