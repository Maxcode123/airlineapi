package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type TicketResponse struct {
	Data []Ticket `json:"data"`
}

func getTickets(c *gin.Context) {
	var req TicketRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	tickets, err := GetAirlineTicketProvider().GetTickets(req)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		c.AbortWithStatus(500)
	}
	c.IndentedJSON(http.StatusOK, TicketResponse{tickets})
}

func createEngine() *gin.Engine {
	router := gin.Default()
	router.GET("/tickets", getTickets)
	return router
}
