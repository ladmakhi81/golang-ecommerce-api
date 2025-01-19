package productrepository

import (
	"database/sql"

	categoryentity "github.com/ladmakhi81/golang-ecommerce-api/internal/category/entity"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	productentity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
	"github.com/lib/pq"
)

type ProductRepository struct {
	storage *storage.Storage
}

func NewProductRepository(storage *storage.Storage) ProductRepository {
	return ProductRepository{
		storage: storage,
	}
}

func (productRepo ProductRepository) CreateProduct(product *productentity.Product) error {
	command := `
		INSERT INTO _products 
		(name, description, category_id, vendor_id, base_price, tags)
		VALUES
		($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at;
	`
	row := productRepo.storage.DB.QueryRow(
		command,
		product.Name,
		product.Description,
		product.Category.ID,
		product.Vendor.ID,
		product.BasePrice,
		pq.Array(product.Tags),
	)
	scanErr := row.Scan(
		&product.ID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if scanErr != nil {
		return scanErr
	}
	return nil
}
func (productRepo ProductRepository) UpdateProductById(product *productentity.Product) error {
	command := `
		UPDATE _products SET
			name = $1,
			description = $2,
			tags = $3,
			is_confirmed = $4,
			confirmed_by_id = $5,
			confirmed_at = $6,
			fee = $7
		WHERE id = $8
	`
	var confirmedById *uint
	if product.ConfirmedBy != nil {
		confirmedById = &product.ConfirmedBy.ID
	}
	row := productRepo.storage.DB.QueryRow(
		command,
		product.Name,
		product.Description,
		pq.Array(product.Tags),
		product.IsConfirmed,
		confirmedById,
		product.ConfirmedAt,
		product.Fee,
		product.ID,
	)
	return row.Err()
}
func (productRepo ProductRepository) FindProductById(id uint) (*productentity.Product, error) {
	command := `
		SELECT
			p.id, p.fee, p.name, p.description,
			p.base_price, p.tags, p.is_confirmed, p.confirmed_at, p.created_at, p.updated_at,
			uc.id, uc.full_name, uc.email,
			c.id, c.name, c.icon, c.created_at, c.updated_at,
			u.id, u.full_name, u.email
		FROM _products p 
			LEFT JOIN _categories c ON c.id = p.category_id
			LEFT JOIN _users u ON u.id = p.vendor_id
			LEFT JOIN _users uc ON uc.id = p.confirmed_by_id
		WHERE p.id = $1
	`
	row := productRepo.storage.DB.QueryRow(
		command,
		id,
	)
	product := new(productentity.Product)
	product.Category = new(categoryentity.Category)
	product.ConfirmedBy = new(userentity.User)
	product.Vendor = new(userentity.User)

	var tags []byte
	var confirmedAt sql.NullTime
	var confirmedByEmail, confirmedByFullName sql.NullString
	var confirmedById sql.NullInt16

	scanErr := row.Scan(
		&product.ID,
		&product.Fee,
		&product.Name,
		&product.Description,
		&product.BasePrice,
		&tags,
		&product.IsConfirmed,
		&confirmedAt,
		&product.CreatedAt,
		&product.UpdatedAt,
		&confirmedById,
		&confirmedByFullName,
		&confirmedByEmail,
		&product.Category.ID,
		&product.Category.Name,
		&product.Category.Icon,
		&product.Category.CreatedAt,
		&product.Category.UpdatedAt,
		&product.Vendor.ID,
		&product.Vendor.FullName,
		&product.Vendor.Email,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		}
		return nil, scanErr
	}
	if len(tags) > 0 {
		pq.Array(&product.Tags).Scan(tags)
	}
	if confirmedByEmail.Valid {
		product.ConfirmedBy.Email = confirmedByEmail.String
	}
	if confirmedByFullName.Valid {
		product.ConfirmedBy.FullName = confirmedByFullName.String
	}
	if confirmedById.Valid {
		product.ConfirmedBy.ID = uint(confirmedById.Int16)
	}
	if confirmedAt.Valid {
		product.ConfirmedAt = confirmedAt.Time
	}
	return product, nil
}
func (productRepo ProductRepository) FindProductsPage(page, limit uint) ([]*productentity.Product, error) {
	command := `
		SELECT
			p.id, p.fee, p.name, p.description,
			p.base_price, p.tags, p.is_confirmed, p.confirmed_at, p.created_at, p.updated_at,
			uc.id, uc.full_name, uc.email,
			c.id, c.name, c.icon, c.created_at, c.updated_at,
			u.id, u.full_name, u.email
		FROM _products p 
			LEFT JOIN _categories c ON c.id = p.category_id
			LEFT JOIN _users u ON u.id = p.vendor_id
			LEFT JOIN _users uc ON uc.id = p.confirmed_by_id
		ORDER BY p.id DESC
		LIMIT $1 OFFSET $2
`
	rows, rowsErr := productRepo.storage.DB.Query(command, limit, page)
	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()
	products := []*productentity.Product{}
	for rows.Next() {
		product := new(productentity.Product)
		product.Category = new(categoryentity.Category)
		product.ConfirmedBy = new(userentity.User)
		product.Vendor = new(userentity.User)

		var tags []byte
		var confirmedAt sql.NullTime
		var confirmedByEmail, confirmedByFullName sql.NullString
		var confirmedById sql.NullInt16

		scanErr := rows.Scan(
			&product.ID,
			&product.Fee,
			&product.Name,
			&product.Description,
			&product.BasePrice,
			&tags,
			&product.IsConfirmed,
			&confirmedAt,
			&product.CreatedAt,
			&product.UpdatedAt,
			&confirmedById,
			&confirmedByFullName,
			&confirmedByEmail,
			&product.Category.ID,
			&product.Category.Name,
			&product.Category.Icon,
			&product.Category.CreatedAt,
			&product.Category.UpdatedAt,
			&product.Vendor.ID,
			&product.Vendor.FullName,
			&product.Vendor.Email,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		if len(tags) > 0 {
			pq.Array(&product.Tags).Scan(tags)
		}
		if confirmedByEmail.Valid {
			product.ConfirmedBy.Email = confirmedByEmail.String
		}
		if confirmedByFullName.Valid {
			product.ConfirmedBy.FullName = confirmedByFullName.String
		}
		if confirmedById.Valid {
			product.ConfirmedBy.ID = uint(confirmedById.Int16)
		}
		if confirmedAt.Valid {
			product.ConfirmedAt = confirmedAt.Time
		}
		products = append(products, product)
	}

	return products, nil
}
func (productRepo ProductRepository) DeleteProductById(id uint) error {
	command := `
		DELETE FROM _products WHERE id = $1
	`
	row := productRepo.storage.DB.QueryRow(command, id)
	return row.Err()
}
