package controllers

import (
	v1 "BoardGame/controllers/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware ...
// CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Upgrade", "Connection", "Origin"},
		AllowCredentials: true,
	})
}

type RestAPIController struct {
	v1.UserController
	v1.FormationController
}

type WebSocketController struct {
	v1.WebSocketController
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	// WebSocket endpoint
	wsAPI := WebSocketController{}
	router.GET("/ws", wsAPI.HandleWebSocket)

	// REST endpoints
	restAPI := RestAPIController{}
	user := router.Group("/user")
	{
		user.GET("/", restAPI.GetUser)
		user.POST("/", restAPI.CreateUser)
		user.PUT("/", restAPI.UpdateUser)
		user.DELETE("/", restAPI.DeleteUser)
	}

	formation := router.Group("/formation")
	{
		formation.GET("/", restAPI.GetFormation)
		// formation.POST("/", restAPI.CreateFormation)
		// formation.PUT("/", restAPI.UpdateFormation)
		// formation.DELETE("/", restAPI.DeleteFormation)
	}

	return router
}
