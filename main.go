package main

import (
	"log"
	"net/http"

	"BoardGame/configs"
	"BoardGame/controllers"
	"BoardGame/models"
)

func main() {
	log.Println("Server starting...")
	// configs.LoadConfig("./env/")
	configs.LoadConfig("./")
	log.Println("JWT_SECRET_KEY: ", configs.Cfg.JWT_SECRET_KEY)
	handler := controllers.SetupRouter()
	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}
	models.ConnectDatabase()
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Server failed to start: ", err)
		}
	}()
	log.Println("Server running on port 8008")
	select {} // Keep the main goroutine running indefinitely

}
