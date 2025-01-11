package categoryservice

import (
	categorydto "github.com/ladmakhi81/golang-ecommerce-api/internal/category/dto"
	categoryentity "github.com/ladmakhi81/golang-ecommerce-api/internal/category/entity"
)

type ICategoryService interface {
	CreateCategory(reqBody categorydto.CreateCategoryReqBody) (*categoryentity.Category, error)
	FindCategoryById(id uint) (*categoryentity.Category, error)
}
