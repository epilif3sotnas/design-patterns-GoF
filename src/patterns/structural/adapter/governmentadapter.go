package adapter


type GovernmentAdapter interface {
	SendGainsInfo([]byte) bool
}