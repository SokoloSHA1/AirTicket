package handler

import (
	"net/http"
	"strconv"

	airticket "github.com/SokoloSHA/AirTicket"
	"github.com/gin-gonic/gin"
)

type getAllTicketsResponse struct {
	Data []airticket.Ticket `json:"data"`
}

func (h *Handler) getAllTicket(c *gin.Context) {
	tickets, err := h.service.TodoTicket.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTicketsResponse{
		Data: tickets,
	})
}

func (h *Handler) updateTicket(c *gin.Context) {
	var input airticket.Ticket

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	err := h.service.TodoTicket.UpdateTicket(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) deleteTicket(c *gin.Context) {
	ticketId, err := strconv.Atoi(c.Param("ticketId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.service.TodoTicket.DeleteTicket(ticketId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
