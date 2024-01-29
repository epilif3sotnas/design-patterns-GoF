package state


type ReadyState struct {
	player *AudioPlayer
}


func NewReadyState(player *AudioPlayer) *ReadyState {
	return &ReadyState{
		player: player,
	}
}

func (self *ReadyState) ClickLock() {
	self.player.SetState(NewLockedState(self.player))
}

func (self *ReadyState) ClickPlay() {
	self.player.StartPlayback()
	self.player.SetState(NewPlayingState(self.player))
}

func (self *ReadyState) ClickNext() {
	self.player.NextSong()
}

func (self *ReadyState) ClickPrevious() {
	self.player.PreviousSong()
}