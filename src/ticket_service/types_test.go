package main

import (
	"errors"
	"testing"
)

func TestTicketRequestInvalidOriginAirport(t *testing.T) {
	_, err := NewTicketRequest("invalid", "CHQ", "2023-08-01")
	if err == nil {
		t.Fatal("error is not raised by invalid origin airport")
	}
	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatal("error is not of ValidationError type")
	}
}

func TestTicketRequestInvalidDepartureAirport(t *testing.T) {
	_, err := NewTicketRequest("ATH", "", "2023-08-01")
	if err == nil {
		t.Fatal("error is not raised by invalid departure airport")
	}
	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatal("error is not of ValidationError type")
	}
}
