package main

// This function returns true if the server does not have enough memory to store the video which is passed in the parameters
func (cs CacheServer) isFull(video Video) bool {
	if video.Size > (y.Cap - cs.MemoryUsed) {
		return true
	}
	return false
}

// This function takes the requests following the array's order and add the videos into the VideoID array if possible
func (cs CacheServer) addVideos() {
	for _,request := range cs.Request {
		_, exist := cs.VideoID[request.VideoID]
		if !exist && !cs.isFull(y.Video[request.VideoID]) {
			cs.VideoID[request.VideoID] = true
			cs.MemoryUsed += y.Video[request.VideoID].Size
		}
	}
}

