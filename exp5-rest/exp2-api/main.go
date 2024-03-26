package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"api/handlers"
	"api/utils"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db, err := utils.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	r.GET("/protected/resource", handlers.AuthMiddleware(), handlers.ProtectedResource)

	http.ListenAndServe(":8080", r)
}
