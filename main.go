package main

import (
	"fmt"
	"sort"
)

var (
	fileName = "me_at_the_zoo"
	y        YoutubeCache
)

func main() {
	read()
	fmt.Println(y)
	fmt.Println("=================================================================")
	for i := range y.Cache {
		feedCasheServerRequests(i)
		sort.Sort(ByScore(y.Cache[i].Request))
		y.Cache[i].addVideos()
	}
	fmt.Println(len(y.Cache))
	encode()
}
