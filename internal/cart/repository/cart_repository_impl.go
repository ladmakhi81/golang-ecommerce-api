package cartrepository

import (
	"database/sql"

	cartentity "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/entity"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	productentity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type CartRepository struct {
	storage *storage.Storage
}

func NewCartRepository(storage *storage.Storage) CartRepository {
	return CartRepository{
		storage: storage,
	}
}

func (cartRepository CartRepository) CreateProductCart(cart *cartentity.Cart) error {
	command := `
		INSERT INTO _carts 
		(product_id, customer_id, quantity, price_item_id) 
		VALUES 
		($1, $2, $3, $4)
		RETURNING id, created_at, updated_at;
	`
	row := cartRepository.storage.DB.QueryRow(
		command,
		cart.Product.ID, cart.Customer.ID, cart.Quantity, cart.PriceItem.ID,
	)
	scanErr := row.Scan(
		&cart.ID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
	return scanErr
}
func (cartRepository CartRepository) FindCustomerCartByProductId(customerID, productID uint) (*cartentity.Cart, error) {
	command := `
		SELECT id FROM _carts WHERE product_id = $1 AND customer_id = $2
	`
	row := cartRepository.storage.DB.QueryRow(command, productID, customerID)
	var cart = new(cartentity.Cart)
	scanErr := row.Scan(
		&cart.ID,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		}
		return nil, scanErr
	}
	return cart, nil
}
func (cartRepository CartRepository) UpdateCartById(cart *cartentity.Cart) error {
	command := `
		UPDATE _carts SET
		quantity = $1
		WHERE id = $2;
	`
	row := cartRepository.storage.DB.QueryRow(command, cart.Quantity, cart.ID)
	return row.Err()
}
func (cartRepository CartRepository) DeleteCartById(cartID uint) error {
	command := `
		DELETE FROM _carts
		WHERE id = $1
	`
	row := cartRepository.storage.DB.QueryRow(command, cartID)
	return row.Err()
}
func (cartRepository CartRepository) FindCartById(cartID uint) (*cartentity.Cart, error) {
	command := `
		SELECT id, product_id, customer_id, price_item_id, quantity FROM _carts WHERE id = $1
	`
	row := cartRepository.storage.DB.QueryRow(command, cartID)
	var cart = new(cartentity.Cart)
	cart.Product = new(productentity.Product)
	cart.Customer = new(userentity.User)
	cart.PriceItem = new(productentity.ProductPrice)
	scanErr := row.Scan(
		&cart.ID,
		&cart.Product.ID,
		&cart.Customer.ID,
		&cart.PriceItem.ID,
		&cart.Quantity,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		}
		return nil, scanErr
	}
	return cart, nil
}
func (cartRepository CartRepository) FindCustomerCart(customerId uint) ([]*cartentity.Cart, error) {
	command := `
		SELECT
		c.id, c.created_at, c.quantity,
		p.id, p.name, p.description, p.base_price,
		pp.id, pp.key, pp.value, pp.extra_price
		FROM _carts c
		INNER JOIN _products p ON p.id = c.product_id
		INNER JOIN _product_prices pp ON c.price_item_id = pp.id
		WHERE c.customer_id = $1

	`
	rows, rowsErr := cartRepository.storage.DB.Query(command, customerId)
	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()
	carts := []*cartentity.Cart{}
	for rows.Next() {
		cart := new(cartentity.Cart)
		cart.Product = new(productentity.Product)
		cart.PriceItem = new(productentity.ProductPrice)
		scanErr := rows.Scan(
			&cart.ID,
			&cart.CreatedAt,
			&cart.Quantity,
			&cart.Product.ID,
			&cart.Product.Name,
			&cart.Product.Description,
			&cart.Product.BasePrice,
			&cart.PriceItem.ID,
			&cart.PriceItem.Key,
			&cart.PriceItem.Value,
			&cart.PriceItem.ExtraPrice,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		carts = append(carts, cart)
	}
	return carts, nil
}
