package state


// std
import(
	"fmt"
)


type AudioPlayer struct {
	state State
	volume uint8
	playlist []string
	currentSong string
	currentSongIdx uint
}


func NewAudioPlayer(volume uint8, playlist []string, currentSong string, currentSongIdx uint) *AudioPlayer {
	return &AudioPlayer{
		state: nil,
		volume: volume,
		playlist: playlist,
		currentSong: currentSong,
		currentSongIdx: currentSongIdx,
	}
}

func (self *AudioPlayer) ClickLock() {
	self.state.ClickLock()
}

func (self *AudioPlayer) ClickPlay() {
	self.state.ClickPlay()
}

func (self *AudioPlayer) ClickNext() {
	self.state.ClickNext()
}

func (self *AudioPlayer) ClickPrevious() {
	self.state.ClickPrevious()
}

func (self *AudioPlayer) StartPlayback() {
	fmt.Println("Start playback")
}

func (self *AudioPlayer) StopPlayback() {
	fmt.Println("Stop playback")
}

func (self *AudioPlayer) NextSong() {
	fmt.Println("Next song")
	fmt.Println("From:", self.currentSong)
	self.currentSongIdx += 1
	self.currentSong = self.playlist[self.currentSongIdx]
	fmt.Println("To:", self.currentSong)
}

func (self *AudioPlayer) PreviousSong() {
	fmt.Println("Previous song")
	fmt.Println("From:", self.currentSong)
	self.currentSongIdx -= 1
	self.currentSong = self.playlist[self.currentSongIdx]
	fmt.Println("To:", self.currentSong)
}

func (self *AudioPlayer) FastForward(time uint8) {
	fmt.Println("Fast forward", time, "seconds")
}

func (self *AudioPlayer) Rewind(time uint8) {
	fmt.Println("Rewind", time, "seconds")
}

func (self *AudioPlayer) GetState() State {
	return self.state
}

func (self *AudioPlayer) SetState(state State) {
	self.state = state
}

func (self *AudioPlayer) GetVolume() uint8 {
	return self.volume
}

func (self *AudioPlayer) SetVolume(volume uint8) {
	self.volume = volume
}

func (self *AudioPlayer) GetPlaylist() []string {
	return self.playlist
}

func (self *AudioPlayer) SetPlaylist(playlist []string) {
	self.playlist = playlist
}

func (self *AudioPlayer) GetCurrentSong() string {
	return self.currentSong
}

func (self *AudioPlayer) SetCurrentSong(currentSong string) {
	self.currentSong = currentSong
}

func (self *AudioPlayer) GetCurrentSongIdx() uint {
	return self.currentSongIdx
}

func (self *AudioPlayer) SetCurrentSongIdx(currentSongIdx uint) {
	self.currentSongIdx = currentSongIdx
}