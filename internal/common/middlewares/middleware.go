package middlewares

import "github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"

type Middleware struct {
	config config.MainConfig
}

func NewMiddleware(config config.MainConfig) Middleware {
	return Middleware{
		config: config,
	}
}
