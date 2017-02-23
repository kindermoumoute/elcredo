package main

func calcTotalScore() int {
	scoreTotal := 0
	nbReq := 0
	for cacheID, server := range y.Cache {
		for requestID := range server.Request {
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
			scoreTotal += (nbrReq * latenceWon)
			nbReq += nbrReq
		}
	}
	return scoreTotal / nbReq
}
