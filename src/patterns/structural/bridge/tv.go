package bridge


import(
	"strconv"
)


type Tv struct {
	isEnabled bool
	volume uint8
	channel uint16
}


func NewTv() *Tv {
	return &Tv{
		isEnabled: false,
		volume: 0,
		channel: 1,
	}
}

func (self *Tv) IsEnabled() bool {
	return self.isEnabled
}

func (self *Tv) Enable() bool {
	self.isEnabled = true
    return true
}

func (self *Tv) Disable() bool {
	self.isEnabled = false
    return true
}

func (self *Tv) GetVolume() uint8 {
	return self.volume
}

func (self *Tv) SetVolume(volume uint8) bool {
	self.volume = volume
    return true
}

func (self *Tv) GetChannel() uint16 {
	return self.channel
}

func (self *Tv) SetChannel(channel uint16) bool {
	self.channel = channel
    return true
}

func (self *Tv) GetDeviceInfo() string {
	isEnabled := strconv.FormatBool(self.isEnabled)
	volume := strconv.FormatUint(uint64(self.volume), 10)
	channel := strconv.FormatUint(uint64(self.channel), 10)

	return "\nIs Enabled: " + isEnabled +
            "\nVolume: " + volume +
            "\nChannel: " + channel
}