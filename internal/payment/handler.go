package payment

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	paymentdto "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/dto"
	paymentservice "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/service"
)

type PaymentHandler struct {
	paymentService paymentservice.IPaymentService
}

func NewPaymentHandler(
	paymentService paymentservice.IPaymentService,
) PaymentHandler {
	return PaymentHandler{paymentService: paymentService}
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
