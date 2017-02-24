package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func read() {
	b, err := ioutil.ReadFile("input/" + fileName + ".in")
	if err != nil {
		panic(err)
	}
	decode(b)
}

func decode(b []byte) {
	tmp := bytes.NewReader(b)
	s := bufio.NewScanner(tmp)
	s.Scan()
	args := toSliceOfInt(s.Text())
	y = YoutubeCache{
		NVideos:    args[0],
		NEndpoints: args[1],
		NRequests:  args[2],
		NCaches:    args[3],
		Cap:        args[4],
		Video:      make([]Video, args[0]),
		Endpoint:   make([]Endpoint, args[1]),
		Request:    make([]Request, args[2]),
		Cache:      make([]CacheServer, args[3]),
	}
	for i := range y.Cache {
		y.Cache[i].VideoID = make(map[int]bool)
	}
	s.Scan()
	args = toSliceOfInt(s.Text())
	for i := range y.Video {
		y.Video[i].Size = args[i]
	}
	for i := range y.Endpoint {
		s.Scan()
		args = toSliceOfInt(s.Text())
		y.Endpoint[i].LatencyDataCenter = args[0]
		y.Endpoint[i].NServer = args[1]
		y.Endpoint[i].Connection = make([]Connection, args[1])
		for j := range y.Endpoint[i].Connection {
			s.Scan()
			args = toSliceOfInt(s.Text())
			y.Endpoint[i].Connection[j].CacheID = args[0]
			y.Endpoint[i].Connection[j].LatencyCacheServer = args[1]
		}
	}
	for i := range y.Request {
		s.Scan()
		args = toSliceOfInt(s.Text())
		y.Request[i].VideoID = args[0]
		y.Request[i].EndpointID = args[1]
		y.Request[i].NRequest = args[2]
	}
}

func encode() string {
	output := fmt.Sprintf("%d", len(y.Cache))
	//fmt.Printf("%d", len(y.Cache))
	for i, server := range y.Cache {
		output += fmt.Sprintf("\n%d", i)
		//fmt.Printf("\n%d", i)
		for videoID := range server.VideoID {
			output += fmt.Sprintf(" %d", videoID)
			//fmt.Printf(" %d", videoID)
		}
	}
	return output
}

func write() {
	output := encode()
	err := ioutil.WriteFile("output/"+fileName+".out", []byte(output), 0644)
	if err != nil {
		return
	}
}

func toSliceOfInt(line string) []int {
	args := strings.Split(line, " ")
	rep := make([]int, len(args))
	var err error
	for i, v := range args {
		rep[i], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}
	return rep
}
