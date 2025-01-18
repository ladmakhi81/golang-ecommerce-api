package orderservice

import (
	orderdto "github.com/ladmakhi81/golang-ecommerce-api/internal/order/dto"
)

type IOrderService interface {
	SubmitOrder(customerId uint, reqBody orderdto.CreateOrderReqBody) (*orderdto.CreateOrderResponse, error)
}
