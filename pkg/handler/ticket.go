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

func (h *Handler) getAllInfoAboutTicket(c *gin.Context) {
	ticketId, err := strconv.Atoi(c.Param("ticketId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	ticket, err := h.service.TodoTicket.GetTicket(ticketId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	users, err := h.service.TodoUser.GetAll(ticketId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userDocuments, err := h.service.TodoDocument.GetAllDocumentsTicket(users)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ticket":        ticket,
		"userDocuments": userDocuments,
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

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
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
