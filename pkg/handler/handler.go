package handler

import (
	"test-module/src/AirTicket/pkg/service"

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
		user.GET("/", h.getAllUser)
		user.POST("/update", h.updateUser)
		user.DELETE("/:userId", h.deleteUser)
	}

	ticket := router.Group("ticket")
	{
		ticket.POST("/update", h.updateTicket)
		ticket.DELETE("/:ticketId", h.deleteTicket)
	}

	return router
}
