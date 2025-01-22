package main

import (
	"github.com/ladmakhi81/golang-ecommerce-api/bootstrap"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"go.uber.org/dig"
)

func main() {
	mainConfig := config.NewMainConfig()
	mainConfig.LoadConfigs()
	diContainer := dig.New()

	appServer := bootstrap.NewAppServer(mainConfig, diContainer)
	appServer.StartApp()
}
