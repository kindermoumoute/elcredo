package main

func addRequestNotSorted(cashServer *CashServer, request *Request) {
	cashServer.request.append(request)
}
