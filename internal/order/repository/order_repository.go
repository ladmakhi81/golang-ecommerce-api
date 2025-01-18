package orderrepository

import orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"

type IOrderRepository interface {
	CreateOrder(order *orderentity.Order) error
	CreateOrderItem(orderItem *orderentity.OrderItem) error
}
