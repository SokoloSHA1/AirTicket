package handler

import (
	"net/http"
	airticket "test-module/src/AirTicket"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllUser(c *gin.Context) {

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
	id := c.Param("userId")
	// id приходит строкой, переделать в int
	err := h.service.TodoUser.DeleteUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
