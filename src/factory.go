package main

func GetAirlineTicketProvider() AirlineTicketProvider {
	return AirlineTicketProvider{Conf.DuffelAPI}
}
