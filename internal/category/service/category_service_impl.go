package categoryservice

import (
	"net/http"

	categorydto "github.com/ladmakhi81/golang-ecommerce-api/internal/category/dto"
	categoryentity "github.com/ladmakhi81/golang-ecommerce-api/internal/category/entity"
	categoryrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/category/repository"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
)

type CategoryService struct {
	categoryRepo categoryrepository.ICategoryRepository
}

func NewCategoryService(categoryRepo categoryrepository.ICategoryRepository) CategoryService {
	return CategoryService{
		categoryRepo,
	}
}

func (categoryService CategoryService) CreateCategory(reqBody categorydto.CreateCategoryReqBody) (*categoryentity.Category, error) {
	categoryExistByName, categoryExistenceErr := categoryService.categoryRepo.IsCategoryNameExist(reqBody.Name)
	if categoryExistenceErr != nil {
		return nil, types.NewServerError(
			"error in checking category name is exist",
			"CategoryService.CreateCategory.IsCategoryNameExist",
			categoryExistenceErr,
		)
	}
	if categoryExistByName {
		return nil, types.NewClientError(
			"category exist by name",
			http.StatusConflict,
		)
	}
	var parentCategory *categoryentity.Category
	if reqBody.ParentCategoryId != 0 {
		foundedCategory, foundedCategoryErr := categoryService.FindCategoryById(reqBody.ParentCategoryId)
		if foundedCategoryErr != nil {
			return nil, foundedCategoryErr
		}
		parentCategory = foundedCategory
	}
	category := categoryentity.NewCategory(reqBody.Name, reqBody.Icon, parentCategory)
	if createCategoryErr := categoryService.categoryRepo.CreateCategory(category); createCategoryErr != nil {
		return nil, types.NewServerError(
			"error in creating category",
			"CategoryService.CreateCategory",
			createCategoryErr,
		)
	}
	return category, nil
}
func (categoryService CategoryService) FindCategoryById(id uint) (*categoryentity.Category, error) {
	category, categoryErr := categoryService.categoryRepo.FindCategoryById(id)
	if categoryErr != nil {
		return nil, types.NewServerError(
			"error in finding category by id",
			"CategoryService.FindCategoryById",
			categoryErr,
		)
	}
	if category == nil {
		return nil, types.NewClientError(
			"category not found with this provided id",
			http.StatusNotFound,
		)
	}
	return category, nil
}
func (categoryService CategoryService) FindCategoriesTree() ([]*categoryentity.Category, error) {
	categories, categoriesErr := categoryService.categoryRepo.FindCategoriesTree()
	if categoriesErr != nil {
		return nil, types.NewServerError(
			"error in retuning categories as tree data",
			"CategoryService.FindCategoriesPage",
			categoriesErr,
		)
	}
	return categories, nil
}
func (categoryService CategoryService) FindCategoriesPage(page, limit uint) ([]*categoryentity.Category, error) {
	categories, categoriesErr := categoryService.categoryRepo.FindCategoriesPage(page, limit)
	if categoriesErr != nil {
		return nil, types.NewServerError(
			"error in finding categories page",
			"CategoryService.FindCategoriesPage",
			categoriesErr,
		)
	}
	return categories, nil
}
func (categoryService CategoryService) DeleteCategoryById(id uint) error {
	if _, err := categoryService.FindCategoryById(id); err != nil {
		return err
	}
	if deleteErr := categoryService.categoryRepo.DeleteCategoryById(id); deleteErr != nil {
		return types.NewServerError(
			"error in deleting category",
			"CategoryService.DeleteCategoryById",
			deleteErr,
		)
	}
	return nil
}
