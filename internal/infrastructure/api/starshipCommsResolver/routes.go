package starshipcommsresolver

import (
	"github.com/gin-gonic/gin"

	services "starshipCommsResolver/internal/pkg/service"
)

func RegisterRoutes2(e *gin.Engine) {
	// Crear una instancia de CommunicationServices usando la función NewService
	topsecretService := services.NewService(nil) // Puedes proporcionar un valor real si es necesario
	topsecretHandler := newHandler(topsecretService)
	var topsecretsplitService = services.NewServiceSplit(nil)
	topsecretSplitHandler := newSplitHandler(topsecretsplitService)

	// Registramos los endpoints y sus respectivos handlerss
	e.POST("/topsecret/", topsecretHandler.topsecret)
	e.POST("/topsecret_split/:satellite_name", topsecretSplitHandler.topsecretSplit)
	e.GET("/topsecret_split/:satellite_name", topsecretSplitHandler.GetsecretSplit)
	e.PATCH("/topsecret_split/:satellite_name", topsecretSplitHandler.PatchtopsecretSplit)

	//	e.GET("/orders", authMiddleware.Authorize(permissions["GET:/orders"]), orderHandler.Get)
}
