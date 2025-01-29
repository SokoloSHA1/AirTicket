package handler

import (
	"net/http"
	"strconv"

	airticket "github.com/SokoloSHA/AirTicket"
	"github.com/gin-gonic/gin"
)

func (h *Handler) updateDocument(c *gin.Context) {
	var input airticket.Document

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	err := h.service.TodoDocument.UpdateDocument(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) deleteDocument(c *gin.Context) {
	documentId, err := strconv.Atoi(c.Param("documentId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.service.TodoDocument.DeleteDocument(documentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
