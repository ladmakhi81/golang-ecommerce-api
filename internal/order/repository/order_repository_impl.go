package orderrepository

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
)

type OrderRepository struct {
	storage *storage.Storage
}

func NewOrderRepository(storage *storage.Storage) OrderRepository {
	return OrderRepository{
		storage: storage,
	}
}

func (orderRepo OrderRepository) CreateOrder(order *orderentity.Order) error {
	command := `
		INSERT INTO _orders (customer_id, status, final_price) VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at;
	`
	row := orderRepo.storage.DB.QueryRow(
		command,
		order.Customer.ID,
		orderentity.OrderStatusPending,
		order.FinalPrice,
	)
	scanErr := row.Scan(
		&order.ID,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if scanErr != nil {
		return scanErr
	}
	return nil
}

func (orderRepo OrderRepository) CreateOrderItem(orderItem *orderentity.OrderItem) error {
	command := `
		INSERT INTO _order_items (product_id, price_item_id, vendor_id, customer_id, order_id, quantity) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	row := orderRepo.storage.DB.QueryRow(
		command,
		orderItem.Product.ID,
		orderItem.PriceItem.ID,
		orderItem.Vendor.ID,
		orderItem.Customer.ID,
		orderItem.Order.ID,
		orderItem.Quantity,
	)
	scanErr := row.Scan(&orderItem.ID)
	if scanErr != nil {
		return scanErr
	}
	return nil
}
