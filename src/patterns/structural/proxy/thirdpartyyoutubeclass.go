package proxy


// std
import (
	"fmt"
)


type ThirdPartyYoutubeClass struct {}


func NewThirdPartyYoutubeClass() *ThirdPartyYoutubeClass {
	return &ThirdPartyYoutubeClass{}
}

func (self *ThirdPartyYoutubeClass) ListVideos() []string {
	return []string{
		"AAA",
		"BBB",
		"CCC",
		"DDD",
		"EEE",
		"FFF",
	}
}

func (self *ThirdPartyYoutubeClass) GetVideoInfo(id string) string {
	return fmt.Sprintf("Video Information\n ID: %s\n Name: AAA\n Description: Descripton of video", id)
}

func (self *ThirdPartyYoutubeClass) DownloadVideo(id string) []byte {
	return []byte("Video FROM http://localhost")
}