package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"golangPrac/controllers"
	"golangPrac/models"
)

func main() {
	godotenv.Load()
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
