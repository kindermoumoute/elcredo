package main

func (y YoutubeCache) calcTotalScore() int {
	scoreTotal := 0
	for _, request := range y.Request {
		scoreTotal += request.Score
	}
	return scoreTotal
}

func (y YoutubeCache) getNumberOfUsedCacheServers() int {
	numberOfUsedCacheServers := 0
	for _, server := range y.Cache
	return 0
}