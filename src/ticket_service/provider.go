package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// Provider responsible for fetching airline tickets.
type AirlineTicketProvider struct {
	Url string
}

// Fetch all the available Ticket(s) matching the given request.
func (provider AirlineTicketProvider) GetTickets(request TicketRequest) ([]Ticket, error) {
	data := []byte(`{"data": {"slices": ` + createSlices(request) + `, "passengers": [{"type": "adult"}], "cabin_class": "economy"}}`)
	client := createHttpClient()
	req, err := http.NewRequest(
		"POST",
		provider.Url+"/air/offer_requests?return_offers=true&supplier_timeout=3000",
		bytes.NewBuffer(data))
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to create http request object: "+err.Error())
		return nil, err
	}

	addHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to make http request: "+err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	respObj := parseResponse(resp)
	var tickets []Ticket
	for i := 0; i < len(respObj.Data.Offers); i++ {
		tickets = append(tickets, parseTicket(respObj.Data.Offers[i]))
	}
	return tickets, nil
}

func createSlices(request TicketRequest) string {
	return `[{"origin": "` + request.Origin + `", "destination": "` + request.Destination + `", "departure_date": "` + request.Date + `"}]`
}

func createHttpClient() *http.Client {
	client := &http.Client{}
	return client
}

func addHeaders(req *http.Request) {
	// req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Duffel-Version", "v1")
	req.Header.Add("Authorization", "Bearer "+Conf.DuffelAPIToken)
}

func parseResponse(resp *http.Response) response {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return response{}
	}
	var result response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cannot unmarshal JSON" + err.Error())
	}
	return result
}

func parseTicket(offerObj offer) Ticket {
	f, err := strconv.ParseFloat(offerObj.TotalAmount, 32)
	if err != nil {
		fmt.Println(err)
	}
	return Ticket{
		TotalAmount:   float32(f),
		TotalCurrency: offerObj.TotalCurrency,
		Departure:     offerObj.Slices[0].Segments[0].DepartingAt,
		Arrival:       offerObj.Slices[0].Segments[0].ArrivingAt,
		Airline:       offerObj.Slices[0].Segments[0].OperatingCarrier.Name,
	}
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
