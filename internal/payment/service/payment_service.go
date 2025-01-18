package paymentservice

import (
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	paymentdto "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/dto"
	paymententity "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/entity"
)

type IPaymentService interface {
	CreatePayment(order *orderentity.Order) (*paymententity.Payment, error)
	GetPayLink(payment *paymententity.Payment) string
	VerifyPayment(customerId uint, reqBody paymentdto.VerifyPaymentReqBody) error
}
