package productrepository

import (
	"database/sql"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	productentity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
)

type ProductPriceRepository struct {
	storage *storage.Storage
}

func NewProductPriceRepository(
	storage *storage.Storage,
) IProductPriceRepository {
	return ProductPriceRepository{
		storage: storage,
	}
}

func (productPriceRepository ProductPriceRepository) CreateProductPrice(productPrice *productentity.ProductPrice) error {
	command := `
		INSERT INTO _product_prices 
		(key, value, extra_price, product_id) 
		VALUES 
		($1, $2, $3, $4)
		RETURNING id, created_at, updated_at;
	`
	row := productPriceRepository.storage.DB.QueryRow(
		command,
		productPrice.Key,
		productPrice.Value,
		productPrice.ExtraPrice,
		productPrice.ProductID,
	)
	scanErr := row.Scan(
		&productPrice.ID,
		&productPrice.CreatedAt,
		&productPrice.UpdatedAt,
	)
	if scanErr != nil {
		return scanErr
	}
	return nil
}
func (productPriceRepository ProductPriceRepository) DeleteProductPriceById(id uint) error {
	command := `
		DELETE FROM _product_prices WHERE id = $1
	`
	row := productPriceRepository.storage.DB.QueryRow(command, id)
	return row.Err()
}
func (productPriceRepository ProductPriceRepository) IsProductPriceItemExist(id uint) (bool, error) {
	command := `
		SELECT COUNT(*) FROM _product_prices WHERE id = $1
	`
	row := productPriceRepository.storage.DB.QueryRow(command, id)
	var count int
	scanErr := row.Scan(&count)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return false, nil
		}
		return false, scanErr
	}
	return count > 0, nil
}
func (productPriceRepository ProductPriceRepository) FindPriceItemById(priceItemID uint) (*productentity.ProductPrice, error) {
	command := `
		SELECT id, key, value, extra_price FROM _product_prices WHERE id = $1
	`
	row := productPriceRepository.storage.DB.QueryRow(command, priceItemID)
	priceItem := new(productentity.ProductPrice)
	scanErr := row.Scan(
		&priceItem.ID,
		&priceItem.Key,
		&priceItem.Value,
		&priceItem.ExtraPrice,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		}
		return nil, scanErr
	}
	return priceItem, nil
}
func (productPriceRepository ProductPriceRepository) FindPricesByProductId(productId uint) (*[]productentity.ProductPrice, error) {
	command := `
		SELECT id, key, value, extra_price FROM _product_prices WHERE product_id = $1
	`
	rows, rowsErr := productPriceRepository.storage.DB.Query(command, productId)
	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()
	productPrices := []productentity.ProductPrice{}
	for rows.Next() {
		priceItem := productentity.ProductPrice{}
		scanErr := rows.Scan(
			&priceItem.ID,
			&priceItem.Key,
			&priceItem.Value,
			&priceItem.ExtraPrice,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		productPrices = append(productPrices, priceItem)
	}
	return &productPrices, nil
}
