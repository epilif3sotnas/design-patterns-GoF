package bridge


type Device interface {
	IsEnabled() bool
	Enable() bool
	Disable() bool
	GetVolume() uint8
	SetVolume(uint8) bool
	GetChannel() uint16
	SetChannel(uint16) bool
	GetDeviceInfo() string
}