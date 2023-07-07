package main

func main() {
	InitConfig()
	provider := AirlineTicketProvider{Url: Conf.DuffelAPI}
	provider._MakeOfferRequest(TicketRequest{origin: "LHR", destination: "JFK", date: "2023-08-08"})
}
