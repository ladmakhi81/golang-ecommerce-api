package categoryrepository

import categoryentity "github.com/ladmakhi81/golang-ecommerce-api/internal/category/entity"

type ICategoryRepository interface {
	IsCategoryNameExist(name string) (bool, error)
	FindCategoryById(id uint) (*categoryentity.Category, error)
	CreateCategory(category *categoryentity.Category) error
	FindCategoriesTree() ([]*categoryentity.Category, error)
	FindCategoriesPage(page, limit uint) ([]*categoryentity.Category, error)
	GetCategoriesCount() (uint, error)
	DeleteCategoryById(id uint) error
	SetCategoryIcon(categoryId uint, filename string) error
}
