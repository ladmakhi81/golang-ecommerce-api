package user

import (
	"fmt"

	"github.com/labstack/echo/v4"
	userevent "github.com/ladmakhi81/golang-ecommerce-api/internal/user/event"
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
	userModule.diContainer.Provide(userevent.NewUserEventsSubscriber)
	userModule.diContainer.Provide(userevent.NewUserEventsContainer)
}

func (userModule UserModule) Run() {
	err := userModule.diContainer.Invoke(func(userRouter UserRouter) {
		userRouter.SetBaseApi(userModule.baseApi)
		userRouter.RegisterRoutes()
	})

	userEventsErr := userModule.diContainer.Invoke(func(userEventsContainer userevent.UserEventsContainer) {
		userEventsContainer.RegisterEvents()
	})

	if err == nil && userEventsErr == nil {
		fmt.Println("UserModule Loaded Successfully")
	} else {
		fmt.Println("UserModule Not Loaded")
	}
}
