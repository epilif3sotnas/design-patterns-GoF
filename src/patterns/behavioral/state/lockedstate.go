package state


type LockedState struct {
	player *AudioPlayer
}


func NewLockedState(player *AudioPlayer) *LockedState {
	return &LockedState{
		player: player,
	}
}

func (self *LockedState) ClickLock() {
	self.player.SetState(NewReadyState(self.player))
}

func (self *LockedState) ClickPlay() {}

func (self *LockedState) ClickNext() {}

func (self *LockedState) ClickPrevious() {}