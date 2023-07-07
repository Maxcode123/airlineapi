package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type AirlineTicketProvider struct {
	Url string
}

type Ticket struct{}

func (provider AirlineTicketProvider) GetTickets(
	departureAirport string, arrivalAirport string) {
}

func (provider AirlineTicketProvider) _MakeOfferRequest() {
	data := []byte(`{"data": {"slices": [{"origin": "LHR", "destination": "JFK", "departure_date": "2023-08-08"}], "passengers": [{"type": "adult"}]}}`)
	client := provider._CreateClient()
	req, err := http.NewRequest("POST", "https://api.duffel.com/air/offer_requests?return_offers=true&supplier_timeout=3000", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	// req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Duffel-Version", "v1")
	req.Header.Add("Authorization", "Bearer ")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp.Body.Close()

	fmt.Println(string(body))
}

func (provider AirlineTicketProvider) _GetOffers(
	departureAirport string, arrivalAirport string, offerRequestId string) {
}

func (provider AirlineTicketProvider) _MakeRequest(url string) {}

func (provider AirlineTicketProvider) _ParseResponse() {}

func (provider AirlineTicketProvider) _CreateClient() *http.Client {
	client := &http.Client{}
	return client
}
