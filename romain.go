package main

// This function returns true if the server does not have enough memory to store the video which is passed in the parameters
func (cs CacheServer) isFull(video *Video) bool {
	if video.Size > y.Cap {
		return true
	}

	return false
}

