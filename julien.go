package main

func addRequestNotSorted(cacheID, requestID int) {
	size := y.Video[y.Request[requestID].VideoID].Size
	nbrReq := y.Request[requestID].NRequest
	latenceDateCenter := y.Endpoint[y.Request[requestID].EndpointID].LatencyDataCenter
	var connection Connection
	for _, c := range y.Endpoint[y.Request[requestID].EndpointID].Connection {
		if c.CacheID == cacheID {
			connection = c
			break
		}
	}
	latenceWon := latenceDateCenter - connection.LatencyCacheServer

	y.Request[requestID].Score = (nbrReq * latenceWon) / size

	y.Cache[cacheID].Request = append(y.Cache[cacheID].Request, &y.Request[requestID])
}

func feedCasheServerRequests(cacheID int) {
	for numEndPoint := 0; numEndPoint < y.NEndpoints; numEndPoint++ {
		for numConnection := 0; numConnection < len(y.Endpoint[numEndPoint].Connection); numConnection++ {
			if y.Endpoint[numEndPoint].Connection[numConnection].CacheID == cacheID {
				for numRequest := 0; numRequest < y.NRequests; numRequest++ {
					if y.Request[numRequest].EndpointID == numEndPoint {
						addRequestNotSorted(cacheID, numRequest)
					}
				}
			}
		}
	}
}
