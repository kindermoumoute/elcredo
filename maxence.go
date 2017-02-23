package main

func (y YoutubeCache) calcTotalScore() int {
	scoreTotal := 0
	for _, request := range y.Request {
		scoreTotal += request.Score
	}
	return scoreTotal
}