package main

func main() {
	InitConfig()
	provider := AirlineTicketProvider{Url: Conf.DuffelAPI}
	tickets := provider.GetTickets("LHR", "JFK", "2023-08-08")
	PrettyPrint(tickets[12])
}
