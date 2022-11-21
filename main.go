package main

import (
	"Pood/app/middleware"
	"Pood/config"
	"Pood/config/router"
	_ "Pood/docs"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// @title                      Pood
// @version                    1.0.0
// @in                         header
// @name                       Authorization
// @host                       localhost:8080
// @BasePath                   /
// @securityDefinitions.basic ApiKeyAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("Not found .env.example file ")
	}

	config.AppConfig()

	app := fiber.New(config.FiberConfig())
	middleware.AppMiddleware(app)
	router.AppRouter(app)

	err := app.Listen(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Panicf("Server is not running: %v", err.Error())
	}

}
