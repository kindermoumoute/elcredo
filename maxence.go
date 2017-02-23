package main

func calcTotalScore() int {
	scoreTotal := 0
	for _, request := range y {
		scoreTotal += request.score
	}
	return scoreTotal
}