package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTickets(c *gin.Context) {
	var req TicketRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}
	tickets := GetAirlineTicketProvider().GetTickets(req)
	c.IndentedJSON(http.StatusOK, tickets)
}

func createEngine() *gin.Engine {
	router := gin.Default()
	router.GET("/tickets", getTickets)
	return router
}
