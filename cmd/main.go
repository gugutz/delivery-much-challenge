package main

import (
	"delivery-much-api/api"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := gin.Default()

	api.Start(app)

	port := os.Getenv("PORT")
	app.Run(":" + port)
}
