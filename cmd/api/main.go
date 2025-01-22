package main

import (
	"github.com/ladmakhi81/golang-ecommerce-api/bootstrap"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/logger"
	"go.uber.org/dig"
)

func main() {
	mainConfig := config.NewMainConfig()
	mainConfig.LoadConfigs()

	diContainer := dig.New()

	appLogger := logger.NewZapLogger()
	diContainer.Provide(func() logger.ILogger {
		return appLogger
	})

	appServer := bootstrap.NewAppServer(mainConfig, diContainer, appLogger)
	appServer.StartApp()
}
