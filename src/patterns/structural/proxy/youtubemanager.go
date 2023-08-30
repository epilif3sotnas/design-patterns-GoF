package proxy


// std
import (
	"fmt"
)


type YoutubeManager struct {
	service ThirdPartyYoutubeLib
}


func NewYoutubeManager(service ThirdPartyYoutubeLib) *YoutubeManager {
	return &YoutubeManager{
		service: service,
	}
}

func (self *YoutubeManager) RenderVideoPage(id string) {
	fmt.Println("Rendering Video Page\n", self.service.GetVideoInfo(id))
}

func (self *YoutubeManager) RenderListPanel() {
	fmt.Println("Rendering List of Videos\n", self.service.ListVideos())
}

func (self *YoutubeManager) RenderVideo(id string) {
	fmt.Println("Rendering Video\n", self.service.DownloadVideo(id))
}

func (self *YoutubeManager) GetService() *ThirdPartyYoutubeLib {
	return &self.service
}

func (self *YoutubeManager) SetService(service ThirdPartyYoutubeLib) {
	self.service = service
}