package proxy


type ThirdPartyYoutubeLib interface {
	ListVideos() []string
	GetVideoInfo(id string) string
	DownloadVideo(id string) []byte
}