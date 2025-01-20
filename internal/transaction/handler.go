package transaction

import (
	"net/http"

	"github.com/labstack/echo/v4"
	responsehandling "github.com/ladmakhi81/golang-ecommerce-api/internal/common/response_handling"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	transactionservice "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/service"
)

type TransactionHandler struct {
	transactionService transactionservice.ITransactionService
	util               utils.Util
}

func NewTransactionHandler(
	transactionService transactionservice.ITransactionService,
) TransactionHandler {
	return TransactionHandler{
		transactionService: transactionService,
		util:               utils.NewUtil(),
	}
}

func (transactionHandler TransactionHandler) GetTransactionsPage(c echo.Context) error {
	pagination := transactionHandler.util.PaginationExtractor(c)
	transactions, err := transactionHandler.transactionService.GetTransactionsPage(pagination.Page, pagination.Limit)
	if err != nil {
		return err
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		transactions,
	)
	return nil
}
