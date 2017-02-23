package main

import (
	"fmt"
	"sort"
)

var (
	fileName = "kittens"
	y        YoutubeCache
)

func main() {
	read()
	for i := range y.Cache {
		fmt.Println(i)
		feedCasheServerRequests(i)
		sort.Sort(ByScore(y.Cache[i].Request))
		(&y.Cache[i]).addVideos()
	}
	write()
}
