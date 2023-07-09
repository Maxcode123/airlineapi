package main

/*
Get an AirlineTicketProvider instance.
DUFFEL_API must have been configured as an environmental variable and the
configuration must have been initialized through InitConfig().
*/
func GetAirlineTicketProvider() AirlineTicketProvider {
	return AirlineTicketProvider{Conf.DuffelAPI}
}
