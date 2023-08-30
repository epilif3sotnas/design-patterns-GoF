package proxy


type CachedYoutubeClass struct {
	service ThirdPartyYoutubeLib
	isDataValid bool
	listVideosCache []string
	videoInfoCache string
}


func NewCachedYoutubeClass(service ThirdPartyYoutubeLib) *CachedYoutubeClass {
	return &CachedYoutubeClass{
		service: service,
		isDataValid: false,
	}
}

func (self *CachedYoutubeClass) ListVideos() []string {
	if !self.isDataValid || self.listVideosCache == nil {
		self.listVideosCache = self.service.ListVideos()
		self.isDataValid = true
	}

	return self.listVideosCache
}

func (self *CachedYoutubeClass) GetVideoInfo(id string) string {
	if !self.isDataValid || self.videoInfoCache == "" {
		self.videoInfoCache = self.service.GetVideoInfo(id)
		self.isDataValid = true
	}

	return self.videoInfoCache
}

func (self *CachedYoutubeClass) DownloadVideo(id string) []byte {
	return self.service.DownloadVideo(id)
}

func (self *CachedYoutubeClass) GetService() *ThirdPartyYoutubeLib {
	return &self.service
}

func (self *CachedYoutubeClass) SetService(service ThirdPartyYoutubeLib) {
	self.service = service
}

func (self *CachedYoutubeClass) GetIsDataValid() bool {
	return self.isDataValid
}

func (self *CachedYoutubeClass) SetIsDataValid(isDataValid bool) {
	self.isDataValid = isDataValid
}

func (self *CachedYoutubeClass) GetListVideosCache() []string {
	return self.listVideosCache
}

func (self *CachedYoutubeClass) SetListVideosCache(listVideosCache []string) {
	self.listVideosCache = listVideosCache
}

func (self *CachedYoutubeClass) GetVideoInfoCache() string {
	return self.videoInfoCache
}

func (self *CachedYoutubeClass) SetVideoInfoCache(videoInfoCache string) {
	self.videoInfoCache = videoInfoCache
}