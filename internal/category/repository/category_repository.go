package categoryrepository

import categoryentity "github.com/ladmakhi81/golang-ecommerce-api/internal/category/entity"

type ICategoryRepository interface {
	IsCategoryNameExist(name string) (bool, error)
	FindCategoryById(id uint) (*categoryentity.Category, error)
	CreateCategory(category *categoryentity.Category) error
}
