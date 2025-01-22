package transaction

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type TransactionRouter struct {
	baseApi *echo.Group
	config  config.MainConfig
	handler TransactionHandler
}

func NewTransactionRouter(
	config config.MainConfig,
	handler TransactionHandler,
) TransactionRouter {
	return TransactionRouter{
		config:  config,
		handler: handler,
	}
}

func (transactionRouter *TransactionRouter) SetBaseApi(baseApi *echo.Group) {
	transactionRouter.baseApi = baseApi
}

func (transactionRouter TransactionRouter) RegisterRoutes() {
	transactionApi := transactionRouter.baseApi.Group("/transactions")

	transactionApi.Use(
		middlewares.AuthMiddleware(
			transactionRouter.config.SecretKey,
		),
	)

	transactionApi.GET(
		"/page",
		transactionRouter.handler.GetTransactionsPage,
		middlewares.RoleMiddleware(userentity.AdminRole),
	)
}
