package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/swaggo/http-swagger"
	"log"
	"os"
	"pood/v2/app/middleware"
	"pood/v2/config"
	"pood/v2/config/router"
	_ "pood/v2/docs"
)

// @title                      Pood - just pood)
// @version                    1.0.0
// @in                         header
// @name                       Authorization
// @BasePath                   /
// @securityDefinitions.basic ApiKeyAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	//178.128.169.26:8080
	if err := godotenv.Load(); err != nil {
		log.Print("Not found .env file, err: ", err)
	}

	app := fiber.New(config.FiberConfig())
	middleware.AppMiddleware(app)
	router.AppRouter(app)

	config.AppConfig()

	err := app.Listen(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))

	if err != nil {
		log.Panicf("Server is not running: %v", err.Error())
	}

}

// $GOPATH/bin/swag init
