package handler

import (
	"net/http"
	"strconv"

	airticket "github.com/SokoloSHA/AirTicket"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllUser(c *gin.Context) {
	ticketId, err := strconv.Atoi(c.Param("ticketId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	users, err := h.service.TodoUser.GetAll(ticketId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

func (h *Handler) updateUser(c *gin.Context) {
	var input airticket.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	err := h.service.TodoUser.UpdateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.service.TodoUser.DeleteUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
