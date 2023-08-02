package bridge


type Remote struct {
	device Device
}


func NewRemote(device Device) *Remote {
	return &Remote{device: device}
}

func (self *Remote) TooglePower() bool {
	if (!self.device.IsEnabled()) {
		return self.device.Enable()
	}
	return self.device.Enable()
}

func (self *Remote) VolumeDown() bool {
	return self.device.SetVolume(self.device.GetVolume() - 1)
}

func (self *Remote) VolumeUp() bool {
	return self.device.SetVolume(self.device.GetVolume() + 1)
}

func (self *Remote) ChannelDown() bool {
	return self.device.SetChannel(self.device.GetChannel() - 1)
}

func (self *Remote) ChannelUp() bool {
	return self.device.SetChannel(self.device.GetChannel() + 1)
}

func (self *Remote) GetDevice() Device {
	return self.device
}

func (self *Remote) SetDevice(device Device) {
	self.device = device
}