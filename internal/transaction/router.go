package transaction

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	transactionservice "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/service"
)

type TransactionRouter struct {
	apiRoute           *echo.Group
	config             config.MainConfig
	transactionHandler TransactionHandler
	transactionService transactionservice.ITransactionService
}

func NewTransactionRouter(
	apiRoute *echo.Group,
	config config.MainConfig,
	transactionService transactionservice.ITransactionService,
) TransactionRouter {
	return TransactionRouter{
		apiRoute:           apiRoute,
		config:             config,
		transactionHandler: NewTransactionHandler(transactionService),
	}
}

func (transactionRouter TransactionRouter) Setup() {
	transactionApi := transactionRouter.apiRoute.Group("/transactions")

	transactionApi.Use(
		middlewares.AuthMiddleware(
			transactionRouter.config.SecretKey,
		),
	)

	transactionApi.GET(
		"/page",
		transactionRouter.transactionHandler.GetTransactionsPage,
	)
}
