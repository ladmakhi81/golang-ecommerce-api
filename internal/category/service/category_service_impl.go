package categoryservice

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	categorydto "github.com/ladmakhi81/golang-ecommerce-api/internal/category/dto"
	categoryentity "github.com/ladmakhi81/golang-ecommerce-api/internal/category/entity"
	categoryrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/category/repository"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
)

type CategoryService struct {
	categoryRepo categoryrepository.ICategoryRepository
	config       config.MainConfig
}

func NewCategoryService(
	categoryRepo categoryrepository.ICategoryRepository,
	config config.MainConfig,
) CategoryService {
	return CategoryService{
		categoryRepo: categoryRepo,
		config:       config,
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
func (categoryService CategoryService) FindCategoriesPage(page, limit uint) ([]*categoryentity.Category, uint, error) {
	categories, categoriesErr := categoryService.categoryRepo.FindCategoriesPage(page, limit)
	if categoriesErr != nil {
		return nil, 0, types.NewServerError(
			"error in finding categories page",
			"CategoryService.FindCategoriesPage",
			categoriesErr,
		)
	}
	categoriesCount, categoriesCountErr := categoryService.categoryRepo.GetCategoriesCount()
	if categoriesCountErr != nil {
		return nil, 0, types.NewServerError(
			"error in get count of categories",
			"CategoryService.FindCategoriesPage.GetCategoriesCount",
			categoriesCountErr,
		)
	}
	return categories, categoriesCount, nil
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
func (categoryService CategoryService) UploadCategoryIcon(categoryId uint, filename string, file multipart.File) (string, error) {
	category, categoryErr := categoryService.FindCategoryById(categoryId)
	if categoryErr != nil {
		return "", categoryErr
	}
	fileExtname := filepath.Ext(filename)
	outputFilename := fmt.Sprintf("%d-%d%s", rand.Intn(10000000000000), int(time.Now().Unix()), fileExtname)
	outputDestination := path.Join(categoryService.config.UploadDirectory, outputFilename)
	outputFile, outputFileErr := os.Create(outputDestination)
	if outputFileErr != nil {
		return "", types.NewServerError(
			"error in creating output file",
			"CategoryService.UploadCategoryIcon.Create",
			outputFileErr,
		)
	}
	if _, copyErr := io.Copy(outputFile, file); copyErr != nil {
		return "", types.NewServerError(
			"error in copy file into output file",
			"CategoryService.UploadCategoryIcon.Copy",
			copyErr,
		)
	}
	if setErr := categoryService.categoryRepo.SetCategoryIcon(category.ID, outputFilename); setErr != nil {
		return "", types.NewServerError(
			"error in setting category icon",
			"CategoryService.UploadCategoryIcon.SetCategoryIcon",
			setErr,
		)
	}
	return outputFilename, nil
}
