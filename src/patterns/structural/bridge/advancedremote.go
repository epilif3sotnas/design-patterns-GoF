package bridge


type AdvancedRemote struct {
	*Remote
}


func NewAdvancedRemote(device Device) *AdvancedRemote {
	return &AdvancedRemote{
		&Remote{device: device},
	}
}

func (self *AdvancedRemote) Mute() bool {
	return self.device.SetVolume(0)
}