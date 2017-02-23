package main

// This function returns true if the server does not have enough memory to store the video which is passed in the parameters
func (CacheServer) isServerFull(video *Video) bool {
	if (y.Cap - cs.UsedMemory) < video.Size {
		return true
	}

	return false
}

