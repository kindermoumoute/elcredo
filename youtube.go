package main

type YoutubeCache struct {
	NVideos    int
	NEndpoints int
	NRequests  int
	NCaches    int
	Cap        int
	Video      []Video
	Endpoint   []Endpoint
	Request    []Request
	Cache      []CacheServer
}

type Video struct {
	Size int
}

type Endpoint struct {
	LatencyDataCenter int
	NServer           int
	Connection        []Connection
}

type Connection struct {
	CacheID            int
	LatencyCacheServer int
}

type Request struct {
	VideoID    int
	EndpointID int
	NRequest   int
	Score      int
}

type CacheServer struct {
	MemoryUsed int
	Request    []*Request
	VideoID    map[int]bool
}
