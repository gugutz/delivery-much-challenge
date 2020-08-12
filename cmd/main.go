package main

import (
	"delivery-much-api/api"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading environment variables", err)
		panic(err)
	}

	app := gin.Default()

	api.Start(app)

	port := os.Getenv("PORT")
	app.Run(":" + port)
}
