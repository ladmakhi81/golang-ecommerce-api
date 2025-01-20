package orderrepository

import (
	"database/sql"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	productentity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
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
		INSERT INTO _orders (customer_id, status, final_price, address_id) VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at;
	`
	row := orderRepo.storage.DB.QueryRow(
		command,
		order.Customer.ID,
		orderentity.OrderStatusPending,
		order.FinalPrice,
		order.Address.ID,
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
func (orderRepo OrderRepository) FindOrderItemsByOrderId(orderId uint) ([]*orderentity.OrderItem, error) {
	command := `
		SELECT
			i.id, i.quantity, --order_items
			p.id, p.name, p.description, p.base_price, p.fee, -- product
			pp.id, pp.key, pp.value, pp.extra_price, -- price item
			u.id, u.email, -- customer
			uu.id, uu.email, uu.full_name, uu.national_id, uu.postal_code, uu.address, -- vendor
			o.id, o.final_price -- order
		FROM _order_items i
			INNER JOIN _products p ON p.id = i.product_id
			INNER JOIN _product_prices pp ON pp.id = i.price_item_id
			INNER JOIN _users u ON u.id = i.vendor_id
			INNER JOIN _users uu ON uu.id = i.customer_id
			INNER JOIN _orders o ON o.id = i.order_id
		WHERE order_id = $1
	`
	rows, rowsErr := orderRepo.storage.DB.Query(command, orderId)
	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()
	orderItems := []*orderentity.OrderItem{}
	for rows.Next() {
		item := new(orderentity.OrderItem)
		item.Product = new(productentity.Product)
		item.PriceItem = new(productentity.ProductPrice)
		item.Customer = new(userentity.User)
		item.Vendor = new(userentity.User)
		item.Order = new(orderentity.Order)
		scanErr := rows.Scan(
			&item.ID, &item.Quantity,
			&item.Product.ID, &item.Product.Name, &item.Product.Description, &item.Product.BasePrice, &item.Product.Fee,
			&item.PriceItem.ID, &item.PriceItem.Key, &item.PriceItem.Value, &item.PriceItem.ExtraPrice,
			&item.Customer.ID, &item.Customer.Email,
			&item.Vendor.ID, &item.Vendor.Email, &item.Vendor.FullName, &item.Vendor.NationalID, &item.Vendor.PostalCode, &item.Vendor.Address,
			&item.Order.ID, &item.Order.FinalPrice,
		)
		if scanErr != nil {
			return nil, scanErr
		}

		orderItems = append(orderItems, item)
	}
	return orderItems, nil
}
func (orderRepo OrderRepository) ChanegOrderStatus(order *orderentity.Order) error {
	command := `
		UPDATE _orders SET
		status = $1,
		status_changed_at = $2
		WHERE id = $3;
	`
	row := orderRepo.storage.DB.QueryRow(command, order.Status, order.StatusChangedAt, order.ID)
	return row.Err()
}
func (orderRepo OrderRepository) FindOrderById(id uint) (*orderentity.Order, error) {
	command := `
		SELECT 
		o.id, o.status, o.final_price, o.status_changed_at ,
		u.id, u.email
		FROM _orders o
		INNER JOIN _users u ON o.customer_id = u.id
		WHERE o.id = $1
	`
	row := orderRepo.storage.DB.QueryRow(command, id)
	order := new(orderentity.Order)
	order.Customer = new(userentity.User)
	scanErr := row.Scan(
		&order.ID,
		&order.Status,
		&order.FinalPrice,
		&order.StatusChangedAt,
		&order.Customer.ID,
		&order.Customer.Email,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		}
		return nil, scanErr
	}
	return order, nil
}
