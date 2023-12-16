package main

import (
	"log"

	"fastprint/configs"
	"fastprint/middlewares"
	"fastprint/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func LoadConfig(port string) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	configs.InitDB()

	e := echo.New()
	middlewares.Cors(e)
	middlewares.BasicLogger(e)
	routes.InitRouter(e)
	e.Logger.Fatal(e.Start(port))
}
