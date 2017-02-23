package main

func addRequestNotSorted(casheServer *CacheServer, request *Request) {
	taille := y.Video[request.VideoID].Size
	nbrReq := request.NRequest
	latenceDateCenter := y.Endpoint[request.EndpointID].LatencyDataCenter
	latenceCasheServer := y.Endpoint[request.EndpointID].Connection[request.EndpointID].LatencyCacheServer
	latenceGagne := latenceDateCenter - latenceCasheServer

	request.Score = (nbrReq * latenceGagne) / taille

	casheServer.Request = append(casheServer.Request, request)
}
