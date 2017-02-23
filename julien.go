package main

import "fmt"

func addRequestNotSorted(cacheServer *CacheServer, request *Request) {
	size := y.Video[request.VideoID].Size
	nbrReq := request.NRequest
	latenceDateCenter := y.Endpoint[request.EndpointID].LatencyDataCenter
	latenceCacheServer := y.Endpoint[request.EndpointID].Connection[request.EndpointID].LatencyCacheServer
	latenceWon := latenceDateCenter - latenceCacheServer

	request.Score = (nbrReq * latenceWon) / size

	cacheServer.Request = append(cacheServer.Request, request)
}

func feedCasheServerRequests(cacheID int) {
	for numEndPoint := 0; numEndPoint < y.NEndpoints; numEndPoint++ {
		for numConnection := 0; numConnection < len(y.Endpoint[numEndPoint].Connection); numConnection++ {
			if y.Endpoint[numEndPoint].Connection[numConnection].CacheID == cacheID {
				for numRequest := 0; numRequest < y.NRequests; numRequest++ {
					if y.Request[numRequest].EndpointID == numEndPoint {
						fmt.Println(len(y.Request), numRequest)
						addRequestNotSorted(&(y.Cache[cacheID]),
							&(y.Request[numRequest]))
					}
				}
			}
		}
	}
}
