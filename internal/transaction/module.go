package transaction

import (
	"fmt"

	"github.com/labstack/echo/v4"
	transactionrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/repository"
	transactionservice "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/service"
	"go.uber.org/dig"
)

type TransactionModule struct {
	diContainer *dig.Container
	baseApi     *echo.Group
}

func NewTransactionModule(
	diContainer *dig.Container,
	baseApi *echo.Group,
) TransactionModule {
	return TransactionModule{
		diContainer: diContainer,
		baseApi:     baseApi,
	}
}

func (transactionModule TransactionModule) LoadModule() {
	transactionModule.diContainer.Provide(NewTransactionHandler)
	transactionModule.diContainer.Provide(NewTransactionRouter)
	transactionModule.diContainer.Provide(transactionservice.NewTransactionService)
	transactionModule.diContainer.Provide(transactionrepository.NewTransactionRepository)
}

func (transactionModule TransactionModule) Run() {
	err := transactionModule.diContainer.Invoke(func(transactionRouter TransactionRouter) {
		transactionRouter.SetBaseApi(transactionModule.baseApi)
		transactionRouter.RegisterRoutes()
	})

	if err == nil {
		fmt.Println("TransactionModule Loaded Successfully")
	} else {
		fmt.Println("TransactionModule Not Loaded", err)
	}
}
