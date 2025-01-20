package categoryservice

import (
	"mime/multipart"

	categorydto "github.com/ladmakhi81/golang-ecommerce-api/internal/category/dto"
	categoryentity "github.com/ladmakhi81/golang-ecommerce-api/internal/category/entity"
)

type ICategoryService interface {
	CreateCategory(reqBody categorydto.CreateCategoryReqBody) (*categoryentity.Category, error)
	FindCategoryById(id uint) (*categoryentity.Category, error)
	FindCategoriesTree() ([]*categoryentity.Category, error)
	FindCategoriesPage(page, limit uint) ([]*categoryentity.Category, uint, error)
	DeleteCategoryById(id uint) error
	UploadCategoryIcon(categoryId uint, filename string, file multipart.File) (string, error)
}
