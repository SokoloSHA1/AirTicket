package handler

import (
	"github.com/SokoloSHA/AirTicket/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/report", h.reportUser)
		user.GET("/:ticketId", h.getAllUser)
		user.POST("/update", h.updateUser)
		user.DELETE("/:userId", h.deleteUser)
	}

	ticket := router.Group("ticket")
	{
		ticket.GET("/", h.getAllTicket)
		ticket.GET("/:ticketId", h.getAllInfoAboutTicket)
		ticket.POST("/update", h.updateTicket)
		ticket.DELETE("/:ticketId", h.deleteTicket)
	}

	document := router.Group("/document")
	{
		document.GET("/:userId", h.getAllDocument)
		document.POST("/update", h.updateDocument)
		document.DELETE("/:documentId", h.deleteDocument)
	}

	return router
}
