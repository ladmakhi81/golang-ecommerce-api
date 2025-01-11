package categoryrepository

import (
	"database/sql"

	categoryentity "github.com/ladmakhi81/golang-ecommerce-api/internal/category/entity"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
)

type CategoryRepository struct {
	storage *storage.Storage
}

func NewCategoryRepository(storage *storage.Storage) CategoryRepository {
	return CategoryRepository{
		storage,
	}
}

func (categoryRepo CategoryRepository) IsCategoryNameExist(name string) (bool, error) {
	command := `
		SELECT COUNT(*) FROM _categories WHERE name = $1
	`
	var count uint
	row := categoryRepo.storage.DB.QueryRow(command, name)
	scanErr := row.Scan(&count)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return false, nil
		}
		return false, scanErr
	}
	return count > 0, nil
}
func (categoryRepo CategoryRepository) FindCategoryById(id uint) (*categoryentity.Category, error) {
	command := `
		SELECT id, name, icon, created_at, updated_at FROM _categories WHERE id = $1;
	`
	category := new(categoryentity.Category)
	row := categoryRepo.storage.DB.QueryRow(command, id)
	scanErr := row.Scan(
		&category.ID,
		&category.Name,
		&category.Icon,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		}
		return nil, scanErr
	}
	return category, nil
}
func (categoryRepo CategoryRepository) CreateCategory(category *categoryentity.Category) error {
	command := `
		INSERT INTO _categories(name, icon, parent_category_id) VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at;
	`
	var parentCategoryId *uint
	if category.ParentCategory != nil {
		parentCategoryId = &category.ParentCategory.ID
	}
	row := categoryRepo.storage.DB.QueryRow(command, category.Name, category.Icon, parentCategoryId)
	scanErr := row.Scan(
		&category.ID,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if scanErr != nil {
		return scanErr
	}
	return nil
}
