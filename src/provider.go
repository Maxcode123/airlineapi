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

type TicketRequest struct {
	origin      string
	destination string
	date        string
}

func (provider AirlineTicketProvider) GetTickets(
	departureAirport string, arrivalAirport string) {
}

func (provider AirlineTicketProvider) _MakeOfferRequest(request TicketRequest) {
	data := []byte(`{"data": {"slices": ` + _CreateSlices(request) + `, "passengers": [{"type": "adult"}]}}`)
	client := provider._CreateClient()
	req, err := http.NewRequest(
		"POST",
		provider.Url+"/air/offer_requests?return_offers=true&supplier_timeout=3000",
		bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	// req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Duffel-Version", "v1")
	req.Header.Add("Authorization", "Bearer "+Conf.DuffelAPIToken)

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

func _CreateSlices(request TicketRequest) string {
	return `[{"origin": "` + request.origin + `", "destination": "` + request.destination + `", "departure_date": "` + request.date + `"}]`
}

func (provider AirlineTicketProvider) _CreateClient() *http.Client {
	client := &http.Client{}
	return client
}
