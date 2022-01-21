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

			id := users.Group(":id", h.userIdentity)
			{
				id.GET("/", h.getUserName)
				id.POST("/play/:serverAddress", h.createSession)
			}
		}

		servers := api.Group("/servers")
		{
			servers.GET("/", h.getServersList)

			address := servers.Group(":address")
			{
				address.POST("/", h.addOrUpdateServer)

				address.GET("/session/:session", h.getSessionData)
			}
		}
	}

	return router
}

// Currently not supported
//       .post("/users/:id/verify", UserController.verify)
//       .post("/users/:id/reset-pin", UserController.resetPin)
//       .get("/enduser-verify/:email/:pin", UserController.verifyEnduser)
