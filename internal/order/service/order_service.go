package orderservice

import (
	orderdto "github.com/ladmakhi81/golang-ecommerce-api/internal/order/dto"
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
)

type IOrderService interface {
	SubmitOrder(customerId uint, reqBody orderdto.CreateOrderReqBody) (*orderdto.CreateOrderResponse, error)
	FindOrderItemsByOrderId(orderId uint) ([]*orderentity.OrderItem, error)
	ChangeOrderStatus(orderId uint, reqBody orderdto.ChangeOrderStatusReqBody) error
	FindOrderById(id uint) (*orderentity.Order, error)
	FindOrdersPage(page, limit uint) ([]*orderentity.Order, uint, error)
}
