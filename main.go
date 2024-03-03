package main

import (
	"aditya-coding-task/routes"
	"os"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	r := gin.New()
	routes.RegisterRoutes(r)
	r.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
