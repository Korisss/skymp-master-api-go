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

			id := users.Group("/:id")
			{
				id.GET("/", h.getUser)
				id.POST("/play/:serverAddress", h.createSession)

				//id.POST("/auth", h.auth)
			}
		}

		//gamedata := api.Group("/gamedata")
		//{
		//	gamedata.PUT("/:serverAddress")
		//	gamedata.GET("/:serverAddress")
		//}

		servers := api.Group("/servers")
		{
			servers.GET("/", h.getServersList)

			address := servers.Group("/:address")
			{
				address.POST("/", h.addServer)

				session := address.Group("/session")
				{
					session.GET("/:session", h.getSessionData)
				}
			}
		}
	}

	return router
}
