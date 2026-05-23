package main

import (
	"litigation_backend/config"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
	if os.Getenv("PORT") == "" {
		log.Println("Running locally")
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, running with existing env vars")
		} else {
			log.Println(".env file loaded")
		}
	}

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
}
