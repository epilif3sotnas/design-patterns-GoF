package iterator


type ProfileIterator interface {
	GetNext() *Profile
	HasMore() bool
}