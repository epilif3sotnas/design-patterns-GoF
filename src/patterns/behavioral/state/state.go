package state


type State interface {
	ClickLock()
	ClickPlay()
	ClickNext()
	ClickPrevious()
}