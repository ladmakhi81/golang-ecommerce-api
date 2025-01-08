package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
)

func main() {
	mainConfig := config.NewMainConfig()
	mainConfig.LoadConfigs()

	port := mainConfig.GetAppPort()
	server := echo.New()

	log.Println("the server is running")

	server.Use(middleware.Logger())
	server.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{"*"},
			},
		),
	)
	log.Fatalln(server.Start(port))
}
