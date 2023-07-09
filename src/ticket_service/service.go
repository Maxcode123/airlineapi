package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketResponse struct {
	Data []Ticket `json:"data"`
}

func getTickets(c *gin.Context) {
	var req TicketRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}
	tickets := GetAirlineTicketProvider().GetTickets(req)
	c.IndentedJSON(http.StatusOK, TicketResponse{tickets})
}

func createEngine() *gin.Engine {
	router := gin.Default()
	router.GET("/tickets", getTickets)
	return router
}
