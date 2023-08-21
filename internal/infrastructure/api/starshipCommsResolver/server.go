// server.go

package starshipcommsresolver

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func CreateServer() *gin.Engine {
	server := gin.Default()

	server.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET,POST,DELETE,PUT",
		RequestHeaders: "Origin, Authorization, Content-Type, Access-Control-Allow-Origin",
		MaxAge:         50 * time.Second,
	}))

	RegisterRoutes2(server)

	return server
}

func RunServer() {
	server := CreateServer()
	server.Run(":8080")
}
