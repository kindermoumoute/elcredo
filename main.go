package main

import (
	"fmt"
	"sort"
)

var (
	fileName = "trending_today"
	y        YoutubeCache
)

func main() {
	read()
	for i := range y.Cache {
		feedCasheServerRequests(i)
		sort.Sort(ByScore(y.Cache[i].Request))
		(&y.Cache[i]).addVideos()
		
	}
	fmt.Println(calcTotalScore())
	write()
}
