package main

import (
	"litigation_backend/config"
	"litigation_backend/routes"
	"litigation_backend/services"
	"litigation_backend/utils"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
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
	e.Binder = &utils.CustomBinder{Validator: validator.New(validator.WithRequiredStructEnabled())}
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
	cronJob := cron.New()
	cronJob.AddFunc("0 8 * * *", func() {
		log.Println("Running at: ", time.Now())
		services.TestWhatsapp()
	})
	cronJob.Start()
	log.Println("CronJob started!")
	e.Logger.Fatal(e.Start(host + ":" + port))
}
