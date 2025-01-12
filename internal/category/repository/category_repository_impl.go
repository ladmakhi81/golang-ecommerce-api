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
func (categoryRepo CategoryRepository) FindCategoriesTree() ([]*categoryentity.Category, error) {
	command := `
		SELECT id, name, icon, parent_category_id, created_at, updated_at FROM _categories;
	`
	rows, rowsErr := categoryRepo.storage.DB.Query(command)

	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()
	var categories []*categoryentity.Category
	for rows.Next() {
		category := categoryentity.Category{}
		var parentId *int
		rowErr := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Icon,
			&parentId,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if rowErr != nil {
			return nil, rowErr
		}
		if parentId != nil {
			category.ParentCategory = new(categoryentity.Category)
			category.ParentCategory.ID = uint(*parentId)
		}
		categories = append(categories, &category)
	}

	return removeParentFromNestedDataCategories(buildNestedDataFromCategories(categories, nil)), nil
}
func (categoryRepo CategoryRepository) FindCategoriesPage(page, limit uint) ([]*categoryentity.Category, error) {
	command := `
		SELECT id, name, icon, created_at, updated_at FROM _categories
		ORDER BY id DESC LIMIT $1 OFFSET $2
	`
	rows, rowsErr := categoryRepo.storage.DB.Query(command, limit, page)
	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()
	var categories []*categoryentity.Category
	for rows.Next() {
		category := categoryentity.Category{}
		scanErr := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Icon,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
func (categoryRepo CategoryRepository) DeleteCategoryById(id uint) error {
	command := `
		DELETE FROM _categories WHERE id = $1
	`
	row := categoryRepo.storage.DB.QueryRow(command, id)
	return row.Err()
}

func buildNestedDataFromCategories(categories []*categoryentity.Category, parentCategoryID *uint) []*categoryentity.Category {
	result := []*categoryentity.Category{}

	for _, category := range categories {
		if (category.ParentCategory == nil && parentCategoryID == nil) ||
			(category.ParentCategory != nil && parentCategoryID != nil && category.ParentCategory.ID == *parentCategoryID) {
			category.SubCategory = buildNestedDataFromCategories(categories, &category.ID)
			result = append(result, category)
		}
	}

	return result
}

func removeParentFromNestedDataCategories(categories []*categoryentity.Category) []*categoryentity.Category {
	for _, category := range categories {
		category.ParentCategory = nil
		category.SubCategory = removeParentFromNestedDataCategories(category.SubCategory)
	}
	return categories
}
