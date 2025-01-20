package orderrepository

import orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"

type IOrderRepository interface {
	CreateOrder(order *orderentity.Order) error
	CreateOrderItem(orderItem *orderentity.OrderItem) error
	FindOrderItemsByOrderId(orderId uint) ([]*orderentity.OrderItem, error)
	ChanegOrderStatus(order *orderentity.Order) error
	FindOrderById(id uint) (*orderentity.Order, error)
	FindOrdersPage(page, limit uint) ([]*orderentity.Order, error)
	GetOrdersCount() (uint, error)
}
