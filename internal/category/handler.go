package category

import (
	"net/http"

	"github.com/labstack/echo/v4"
	categorydto "github.com/ladmakhi81/golang-ecommerce-api/internal/category/dto"
	categoryservice "github.com/ladmakhi81/golang-ecommerce-api/internal/category/service"
	responsehandling "github.com/ladmakhi81/golang-ecommerce-api/internal/common/response_handling"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
)

type CategoryHandler struct {
	categoryService categoryservice.ICategoryService
	util            utils.Util
	translation     translations.ITranslation
}

func NewCategoryHandler(
	categoryService categoryservice.ICategoryService,
	translation translations.ITranslation,
) CategoryHandler {
	return CategoryHandler{
		categoryService: categoryService,
		util:            utils.NewUtil(),
		translation:     translation,
	}
}

func (categoryHandler CategoryHandler) CreateCategory(c echo.Context) error {
	var reqBody categorydto.CreateCategoryReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError(
			categoryHandler.translation.Message("errors.invalid_request_body"),
			http.StatusBadRequest,
		)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	createdCategory, createdCategoryErr := categoryHandler.categoryService.CreateCategory(
		reqBody,
	)
	if createdCategoryErr != nil {
		return createdCategoryErr
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusCreated,
		createdCategory,
	)
	return nil
}
func (categoryHandler CategoryHandler) GetCategoriesTree(c echo.Context) error {
	categories, categoriesErr := categoryHandler.categoryService.FindCategoriesTree()
	if categoriesErr != nil {
		return categoriesErr
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		categories,
	)
	return nil
}
func (categoryHandler CategoryHandler) GetCategoriesPage(c echo.Context) error {
	pagination := categoryHandler.util.PaginationExtractor(c)
	categories, categoriesCount, categoriesErr := categoryHandler.categoryService.FindCategoriesPage(pagination.Page, pagination.Limit)
	if categoriesErr != nil {
		return categoriesErr
	}
	paginatedResponse := types.NewPaginationResponse(
		categoriesCount,
		pagination,
		categories,
	)
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		paginatedResponse,
	)
	return nil
}
func (categoryHandler CategoryHandler) DeleteCategoryById(c echo.Context) error {
	categoryId, parsedCategoryErr := categoryHandler.util.NumericParamConvertor(
		c.Param("id"),
		categoryHandler.translation.Message("category.invalid_id"),
	)
	if parsedCategoryErr != nil {
		return parsedCategoryErr
	}
	deleteErr := categoryHandler.categoryService.DeleteCategoryById(categoryId)
	if deleteErr != nil {
		return deleteErr
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		nil,
	)
	return nil
}
func (categoryHandler CategoryHandler) UploadCategoryIcon(c echo.Context) error {
	categoryId, parsedCategoryIdErr := categoryHandler.util.NumericParamConvertor(
		c.Param("categoryId"),
		categoryHandler.translation.Message("category.invalid_id"),
	)
	if parsedCategoryIdErr != nil {
		return parsedCategoryIdErr
	}
	fileHeader, fileHeaderErr := c.FormFile("icon")
	if fileHeaderErr != nil {
		return types.NewClientError(
			categoryHandler.translation.Message("category.invalid_icon"),
			http.StatusBadRequest,
		)
	}
	file, fileErr := fileHeader.Open()
	if fileErr != nil {
		return types.NewServerError(
			"error in opening file header as file",
			"CategoryHandler.UploadCategoryIcon.FileHeader.Open",
			fileErr,
		)
	}
	defer file.Close()
	uploadedFilename, uploadedFileErr := categoryHandler.categoryService.UploadCategoryIcon(categoryId, fileHeader.Filename, file)
	if uploadedFileErr != nil {
		return uploadedFileErr
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		uploadedFilename,
	)
	return nil
}
