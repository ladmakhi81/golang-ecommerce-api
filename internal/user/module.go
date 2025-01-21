package user

import (
	"fmt"

	"github.com/labstack/echo/v4"
	userrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/user/repository"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
	"go.uber.org/dig"
)

type UserModule struct {
	diContainer *dig.Container
	baseApi     *echo.Group
}

func NewUserModule(
	diContainer *dig.Container,
	baseApi *echo.Group,
) UserModule {
	return UserModule{
		diContainer: diContainer,
		baseApi:     baseApi,
	}
}

func (userModule UserModule) LoadModule() {
	userModule.diContainer.Provide(userservice.NewUserService)
	userModule.diContainer.Provide(userservice.NewUserAddressService)
	userModule.diContainer.Provide(userrepository.NewUserRepository)
	userModule.diContainer.Provide(userrepository.NewUserAddressRepository)
	userModule.diContainer.Provide(NewUserRouter)
	userModule.diContainer.Provide(NewUserHandler)
}

func (userModule UserModule) Run() {
	err := userModule.diContainer.Invoke(func(userRouter UserRouter) {
		userRouter.SetBaseApi(userModule.baseApi)
		userRouter.RegisterRoutes()
	})

	if err == nil {
		fmt.Println("UserModule Loaded Successfully")
	} else {
		fmt.Println("UserModule Not Loaded", err)
	}
}
