package payment

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	paymentdto "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/dto"
	paymentservice "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/service"
)

type PaymentHandler struct {
	paymentService paymentservice.IPaymentService
	util           utils.Util
}

func NewPaymentHandler(
	paymentService paymentservice.IPaymentService,
) PaymentHandler {
	return PaymentHandler{
		paymentService: paymentService,
		util:           utils.NewUtil(),
	}
}

func (paymentHandler PaymentHandler) VerifyPayment(c echo.Context) error {
	customerId := c.Get("AuthClaim").(*types.AuthClaim).ID
	var reqBody paymentdto.VerifyPaymentReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError("invalid request body", http.StatusBadRequest)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	if verifyErr := paymentHandler.paymentService.VerifyPayment(customerId, reqBody); verifyErr != nil {
		return verifyErr
	}
	return nil
}
func (paymentHandler PaymentHandler) GetPaymentsPage(c echo.Context) error {
	pagination := paymentHandler.util.PaginationExtractor(c)
	payments, paymentsErr := paymentHandler.paymentService.GetPaymentsPage(
		pagination.Page,
		pagination.Limit,
	)
	if paymentsErr != nil {
		return paymentsErr
	}
	c.JSON(http.StatusOK, map[string]any{"data": payments})
	return nil
}
