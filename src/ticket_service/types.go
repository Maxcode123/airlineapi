package main

// A Ticket holds the price for a flight along with departure/arrival data.
type Ticket struct {
	TotalAmount   float32 `json:"total_amount"`
	TotalCurrency string  `json:"total_currency"`
	Departure     string  `json:"departure"`
	Arrival       string  `json:"arrival"`
	Airline       string  `json:"airline"`
}

// Data needed to request a Ticket.
type TicketRequest struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
}

type operator struct {
	ConditionsOfCarriageUrl any    `json:"conditions_of_carriage_url"`
	IataCode                string `json:"iata_code"`
	Id                      string `json:"id"`
	LogoLockupUrl           any    `json:"logo_lockup_url"`
	LogoSymbolUrl           string `json:"logo_symbol_url"`
	Name                    string `json:"name"`
}

type segment struct {
	Aircraft                     map[string]any `json:"aircraft"`
	ArrivingAt                   string         `json:"arriving_at"`
	DepartingAt                  string         `json:"departing_at"`
	Destination                  map[string]any `json:"destination"`
	DestinationTerminal          string         `json:"destination_terminal"`
	Distance                     string         `json:"distance"`
	Duration                     string         `json:"duration"`
	Id                           string         `json:"id"`
	MarketingCarrier             map[string]any `json:"marketing_carrier"`
	MarketingCarrierFlightNumber string         `json:"marketing_carrier_flight_number"`
	OperatingCarrier             operator       `json:"operating_carrier"`
	OperatingCarrierFlightNumber string         `json:"operating_carrier_flight_number"`
	Origin                       map[string]any `json:"origin"`
	OriginTerminal               string         `json:"origin_terminal"`
	Passengers                   []any          `json:"passengers"`
	Stops                        []any          `json:"stops"`
}

type slice struct {
	Conditions      map[string]any `json:"conditions"`
	Destination     map[string]any `json:"destination"`
	DestinationType string         `json:"destination_type"`
	Duration        string         `json:"duration"`
	FareBrandName   string         `json:"fare_brand_name"`
	Id              string         `json:"id"`
	Origin          map[string]any `json:"origin"`
	OriginType      string         `json:"origin_type"`
	Segments        []segment      `json:"segments"`
}

type offer struct {
	AllowedPassengerIdentityDocumentTypes   []string          `json:"allowed_passenger_identity_document_types"`
	BaseAmount                              string            `json:"base_amount"`
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
	Slices                                  []slice           `json:"slices"`
	SupportedPassengerIdentityDocumentTypes []string          `json:"suppoerted_passenger_identity_document_types"`
	TaxAmount                               string            `json:"tax_amount"`
	TaxCurrency                             string            `json:"tax_currency"`
	TotalAmount                             string            `json:"total_amount"`
	TotalCurrency                           string            `json:"total_currency"`
	TotalEmissionsKg                        string            `json:"total_emissions_kg"`
	UpdatedAt                               string            `json:"updated_at"`
}

type response struct {
	Data struct {
		Slices     []any   `json:"slices"`
		Passengers []any   `json:"passengers"`
		Offers     []offer `json:"offers"`
		LiveMode   bool    `json:"live_mode"`
		Id         string  `json:"id"`
		CreatedAt  string  `json:"created_at"`
		ClientKey  string  `json:"client_key"`
		CabinClass string  `json:"cabin_class"`
	} `json:"data"`
}
