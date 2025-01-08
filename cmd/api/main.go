package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
)

func main() {
	mainConfig := config.NewMainConfig()
	mainConfig.LoadConfigs()

	port := mainConfig.GetAppPort()
	server := echo.New()

	log.Println("the server is running")

	log.Fatalln(server.Start(port))
}
