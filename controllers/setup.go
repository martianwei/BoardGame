package controllers

import (
	v1 "BoardGame/controllers/v1"
	"BoardGame/middleware"

	"github.com/gin-gonic/gin"
)

type RestAPIController struct {
	v1.UserController
	v1.FormationController
	v1.LoginController
	v1.GameController
	v1.GameHistoryController
}

type WebSocketController struct {
	v1.WebSocketController
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// // CORS middleware
	router.Use(middleware.CORSMiddleware())

	// WebSocket endpoint
	wsAPI := WebSocketController{}
	router.GET("/ws", wsAPI.HandleWebSocket)

	// REST endpoints
	restAPI := RestAPIController{}
	user := router.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("", restAPI.GetUser)
		user.POST("", restAPI.CreateUser)
		user.PUT("", restAPI.UpdateUser)
		user.DELETE("", restAPI.DeleteUser)
	}

	formation := router.Group("/formation")
	formation.Use(middleware.AuthMiddleware())
	{
		formation.GET("", restAPI.GetFormation)
		formation.POST("", restAPI.CreateFormation)
		// formation.PUT("/", restAPI.UpdateFormation)
		// formation.DELETE("/", restAPI.DeleteFormation)
	}

	login := router.Group("/login")
	{
		login.POST("", restAPI.Login)
	}

	game := router.Group("/game")
	formation.Use(middleware.AuthMiddleware())
	{
		game.POST("", restAPI.CreateGame)
	}

	gamehistory := router.Group("/gamehistory")
	formation.Use(middleware.AuthMiddleware())
	{
		gamehistory.GET("", restAPI.GetGameHistories)
		gamehistory.POST("", restAPI.CreateGameHistories)
	}

	return router
}
