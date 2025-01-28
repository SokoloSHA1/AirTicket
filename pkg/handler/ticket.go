package handler

import (
	"net/http"
	airticket "test-module/src/AirTicket"

	"github.com/gin-gonic/gin"
)

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

}
