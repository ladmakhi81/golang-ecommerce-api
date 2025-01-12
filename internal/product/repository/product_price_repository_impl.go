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
) ProductPriceRepository {
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
