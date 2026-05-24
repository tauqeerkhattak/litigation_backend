package main

import (
	"litigation_backend/config"
	"litigation_backend/routes"
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
	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOriginFunc: func(origin string) (bool, error) {
				return true, nil
			},
			AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "ngrok-skip-browser-warning"},
		}),
	)

	router := e.Group("/api/v1")
	routes.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := os.Getenv("HOST_URL")
	e.Logger.Fatal(e.Start(host + ":" + port))
}
