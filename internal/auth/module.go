package auth

import (
	"fmt"

	"github.com/labstack/echo/v4"
	authservice "github.com/ladmakhi81/golang-ecommerce-api/internal/auth/service"
	"go.uber.org/dig"
)

type AuthModule struct {
	diContainer *dig.Container
	baseApi     *echo.Group
}

func NewAuthModule(
	diContainer *dig.Container,
	baseApi *echo.Group,
) AuthModule {
	return AuthModule{
		diContainer: diContainer,
		baseApi:     baseApi,
	}
}

func (authModule AuthModule) LoadModule() {
	authModule.diContainer.Provide(NewAuthHandler)
	authModule.diContainer.Provide(authservice.NewAuthService)
	authModule.diContainer.Provide(authservice.NewJwtService)
	authModule.diContainer.Provide(NewAuthRouter)
}

func (authModule AuthModule) Run() {
	err := authModule.diContainer.Invoke(func(authRouter AuthRouter) {
		authRouter.SetBaseApi(authModule.baseApi)
		authRouter.RegisterRoutes()
	})

	if err == nil {
		fmt.Println("AuthModule Loaded Successfully")
	} else {
		fmt.Println("AuthModule Not Loaded", err)
	}
}
