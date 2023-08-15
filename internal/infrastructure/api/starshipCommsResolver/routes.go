package starshipcommsresolver

import "github.com/gin-gonic/gin"


func RegisterRoutes2(e *gin.Engine, authService auth.AuthService) {

	//TopsecrestService := order.NewServiceOrder(orderRepo)
	//topsecret := newHandler(orderService, nil)


	
	// Registramos los endpoints y sus respectivos handlers
	e.POST("/topsecret/", topsecret.)
//	e.GET("/orders", authMiddleware.Authorize(permissions["GET:/orders"]), orderHandler.Get)

}