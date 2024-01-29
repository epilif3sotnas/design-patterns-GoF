package state


// std
import(
	"time"
)


type PlayingState struct {
	player *AudioPlayer
}


func NewPlayingState(player *AudioPlayer) *PlayingState {
	return &PlayingState{
		player: player,
	}
}

func (self *PlayingState) ClickLock() {
	self.player.SetState(NewLockedState(self.player))
}

func (self *PlayingState) ClickPlay() {
	self.player.StopPlayback()
	self.player.SetState(NewReadyState(self.player))
}

func (self *PlayingState) ClickNext() {
	if time.Now().Unix() % 2 == 0 {
		self.player.NextSong()
	} else {
		self.player.FastForward(5)
	}
}

func (self *PlayingState) ClickPrevious() {
	if time.Now().Unix() % 2 == 0 {
		self.player.PreviousSong()
	} else {
		self.player.Rewind(5)
	}
}