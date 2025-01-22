package transaction

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type TransactionRouter struct {
	baseApi    *echo.Group
	handler    TransactionHandler
	middleware middlewares.Middleware
}

func NewTransactionRouter(
	handler TransactionHandler,
	middleware middlewares.Middleware,
) TransactionRouter {
	return TransactionRouter{
		handler:    handler,
		middleware: middleware,
	}
}

func (transactionRouter *TransactionRouter) SetBaseApi(baseApi *echo.Group) {
	transactionRouter.baseApi = baseApi
}

func (transactionRouter TransactionRouter) RegisterRoutes() {
	transactionApi := transactionRouter.baseApi.Group("/transactions")

	transactionApi.Use(
		transactionRouter.middleware.AuthMiddleware(),
	)

	transactionApi.GET(
		"/page",
		transactionRouter.handler.GetTransactionsPage,
		transactionRouter.middleware.RoleMiddleware(userentity.AdminRole),
	)
}
