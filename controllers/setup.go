package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "golangPrac/controllers/v1"
)

// CORSMiddleware ...
// CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
}

type Controller struct {
	v1.UserController
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	controller := Controller{}
	user := router.Group("/user")
	{
		// user.GET("/{id}", GetUser)
		user.POST("/", controller.CreateUser)
	}

	return router
}
