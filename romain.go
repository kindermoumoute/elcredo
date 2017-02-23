package main

// This function returns true if the server does not have enough memory to store the video which is passed in the parameters
func (cs CacheServer) isFull(video Video) bool {
	if video.Size > (y.Cap - cs.MemoryUsed) {
		return true
	}
	return false
}

func (cs CacheServer) addVideos() {
	for _,request := range cs.Request {
		if !cs.isFull(y.Video[request.VideoID]) {
			cs.VideoID = append(cs.VideoID, request.VideoID)
		}
	}
}

