package orderservice

import (
	orderdto "github.com/ladmakhi81/golang-ecommerce-api/internal/order/dto"
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
)

type IOrderService interface {
	SubmitOrder(customerId uint, reqBody orderdto.CreateOrderReqBody) (*orderdto.CreateOrderResponse, error)
	FindOrderItemsByOrderId(orderId uint) ([]*orderentity.OrderItem, error)
}
