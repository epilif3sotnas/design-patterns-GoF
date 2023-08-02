package bridge


import(
	"strconv"
)


type Radio struct {
	isEnabled bool
	volume uint8
	channel uint16
}


func NewRadio() *Radio {
	return &Radio{
		isEnabled: false,
		volume: 0,
		channel: 1,
	}
}

func (self *Radio) IsEnabled() bool {
	return self.isEnabled
}

func (self *Radio) Enable() bool {
	self.isEnabled = true
    return true
}

func (self *Radio) Disable() bool {
	self.isEnabled = false
    return true
}

func (self *Radio) GetVolume() uint8 {
	return self.volume
}

func (self *Radio) SetVolume(volume uint8) bool {
	self.volume = volume
    return true
}

func (self *Radio) GetChannel() uint16 {
	return self.channel
}

func (self *Radio) SetChannel(channel uint16) bool {
	self.channel = channel
    return true
}

func (self *Radio) GetDeviceInfo() string {
	isEnabled := strconv.FormatBool(self.isEnabled)
	volume := strconv.FormatUint(uint64(self.volume), 10)
	channel := strconv.FormatUint(uint64(self.channel), 10)

	return "\nIs Enabled: " + isEnabled +
            "\nVolume: " + volume +
            "\nChannel: " + channel
}