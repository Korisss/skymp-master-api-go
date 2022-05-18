package handler

import (
	"github.com/Korisss/skymp-master-api-go/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.register)
			users.POST("/login", h.login)
			users.POST("/reset-password", h.resetPassword)
			users.GET("/verify", h.sendVerificationCode)
			users.POST("/verify", h.verify)

			id := users.Group(":id")
			{
				id.GET("/", h.getUserName)
				id.POST("/play/:serverAddress", h.userIdentity, h.createSession)
			}
		}

		servers := api.Group("/servers")
		{
			servers.GET("/", h.getServersList)

			address := servers.Group(":address")
			{
				address.POST("/", h.addOrUpdateServer)

				address.GET("/sessions/:session", h.getSessionData)
			}
		}
	}

	return router
}
