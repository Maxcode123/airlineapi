package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AirlineTicketProvider struct {
	Url string
}

type Ticket struct {
	TotalAmount   float32 `json:"total_amount"`
	TotalCurrency string  `json:"total_currency"`
	Departure     string  `json:"departure"`
	Arrival       string  `json:"arrival"`
	Airline       string  `json:"airline"`
}

type TicketRequest struct {
	origin      string
	destination string
	date        string
}

type Operator struct {
	ConditionsOfCarriageUrl any    `json:"conditions_of_carriage_url"`
	IataCode                string `json:"iata_code"`
	Id                      string `json:"id"`
	LogoLockupUrl           any    `json:"logo_lockup_url"`
	LogoSymbolUrl           string `json:"logo_symbol_url"`
	Name                    string `json:"name"`
}

type Segment struct {
	Aircraft                     map[string]any `json:"aircraft"`
	ArrivingAt                   string         `json:"arriving_at"`
	DepartingAt                  string         `json:"departing_at"`
	Destination                  map[string]any `json:"destination"`
	DestinationTerminal          string         `json:"destination_terminal"`
	Distance                     float32        `json:"distance"`
	Duration                     string         `json:"duration"`
	Id                           string         `json:"id"`
	MarketingCarrier             map[string]any `json:"marketing_carrier"`
	MarketingCarrierFlightNumber string         `json:"marketing_carrier_flight_number"`
	OperatingCarrier             Operator       `json:"operating_carrier"`
	OperatingCarrierFlightNumber string         `json:"operating_carrier_flight_number"`
	Origin                       map[string]any `json:"origin"`
	OriginTerminal               string         `json:"origin_terminal"`
	Passengers                   []any          `json:"passengers"`
	Stops                        []any          `json:"stops"`
}

type Slice struct {
	Conditions      map[string]any `json:"conditions"`
	Destination     map[string]any `json:"destination"`
	DestinationType string         `json:"destination_type"`
	Duration        string         `json:"duration"`
	FareBrandName   string         `json:"fare_brand_name"`
	Id              string         `json:"id"`
	Origin          map[string]any `json:"origin"`
	OriginType      string         `json:"origin_type"`
	Segments        []Segment      `json:"segments"`
}

type Offer struct {
	AllowedPassengerIdentityDocumentTypes   []string          `json:"allowed_passenger_identity_document_types"`
	BaseAmount                              float32           `json:"base_amount"`
	BaseCurrency                            string            `json:"base_currency"`
	Conditions                              map[string]any    `json:"conditions"`
	CreatedAt                               string            `json:"created_at"`
	ExpiresAt                               string            `json:"expires_at"`
	Id                                      string            `json:"id"`
	LiveMode                                bool              `json:"live_mode"`
	Owner                                   map[string]string `json:"owner"`
	Partial                                 bool              `json:"partial"`
	PassengerIdentityDocumentsRequired      bool              `json:"passenger_identity_documents_required"`
	Passengers                              []map[string]any  `json:"passengers"`
	PaymentRequirements                     map[string]any    `json:"payment_requirements"`
	PrivateFares                            []any             `json:"private_fares"`
	Slices                                  []Slice           `json:"slices"`
	SupportedPassengerIdentityDocumentTypes []string          `json:"suppoerted_passenger_identity_document_types"`
	TaxAmount                               float32           `json:"tax_amount"`
	TaxCurrency                             string            `json:"tax_currency"`
	TotalAmount                             float32           `json:"total_amount"`
	TotalCurrency                           string            `json:"total_currency"`
	TotalEmissionsKg                        string            `json:"total_emissions_kg"`
	UpdatedAt                               string            `json:"updated_at"`
}

type Response struct {
	Data struct {
		Slices     []any   `json:"slices"`
		Passengers []any   `json:"passengers"`
		Offers     []Offer `json:"offers"`
		LiveMode   bool    `json:"live_mode"`
		Id         string  `json:"id"`
		CreatedAt  string  `json:"created_at"`
		ClientKey  string  `json:"client_key"`
		CabinClass string  `json:"cabin_class"`
	} `json:"data"`
}

func (provider AirlineTicketProvider) GetTickets(
	departureAirport string, arrivalAirport string) {
}

func (provider AirlineTicketProvider) _MakeOfferRequest(request TicketRequest) {
	data := []byte(`{"data": {"slices": ` + _CreateSlices(request) + `, "passengers": [{"type": "adult"}], "cabin_class": "economy"}}`)
	client := _CreateHttpClient()
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

	defer resp.Body.Close()
	respObj := _ParseResponse(resp)
	var tickets []Ticket
	for i := 0; i < len(respObj.Data.Offers); i++ {
		tickets = append(tickets, _ParseTicket(respObj.Data.Offers[i]))
	}
	// r, _ := json.Marshal(tickets[0])
	println(len(tickets))

}

func _CreateSlices(request TicketRequest) string {
	return `[{"origin": "` + request.origin + `", "destination": "` + request.destination + `", "departure_date": "` + request.date + `"}]`
}

func _CreateHttpClient() *http.Client {
	client := &http.Client{}
	return client
}

func _ParseResponse(resp *http.Response) Response {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Response{}
	}
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cannot unmarshal JSON" + err.Error())
	}
	return result
}

func _ParseTicket(offer Offer) Ticket {
	return Ticket{
		TotalAmount:   float32(offer.TotalAmount),
		TotalCurrency: offer.TotalCurrency,
		Departure:     offer.Slices[0].Segments[0].DepartingAt,
		Arrival:       offer.Slices[0].Segments[0].ArrivingAt,
		Airline:       offer.Slices[0].Segments[0].OperatingCarrier.Name,
	}
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
