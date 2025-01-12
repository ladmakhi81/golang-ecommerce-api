package category

import (
	"net/http"

	"github.com/labstack/echo/v4"
	categorydto "github.com/ladmakhi81/golang-ecommerce-api/internal/category/dto"
	categoryservice "github.com/ladmakhi81/golang-ecommerce-api/internal/category/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
)

type CategoryHandler struct {
	categoryService categoryservice.ICategoryService
	util            utils.Util
}

func NewCategoryHandler(categoryService categoryservice.ICategoryService) CategoryHandler {
	return CategoryHandler{
		categoryService: categoryService,
		util:            utils.NewUtil(),
	}
}

func (categoryHandler CategoryHandler) CreateCategory(c echo.Context) error {
	var reqBody categorydto.CreateCategoryReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError("invalid request body", http.StatusBadRequest)
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
	c.JSON(http.StatusCreated, map[string]any{
		"category": createdCategory,
		"success":  true,
	})
	return nil
}
func (categoryHandler CategoryHandler) GetCategoriesTree(c echo.Context) error {
	categories, categoriesErr := categoryHandler.categoryService.FindCategoriesTree()
	if categoriesErr != nil {
		return categoriesErr
	}
	c.JSON(http.StatusOK, map[string]any{
		"data": categories,
	})
	return nil
}
func (categoryHandler CategoryHandler) GetCategoriesPage(c echo.Context) error {
	pagination := categoryHandler.util.PaginationExtractor(c)
	categories, categoriesErr := categoryHandler.categoryService.FindCategoriesPage(pagination.Page, pagination.Limit)
	if categoriesErr != nil {
		return categoriesErr
	}
	c.JSON(http.StatusOK, map[string]any{
		"data": categories,
	})
	return nil
}
func (categoryHandler CategoryHandler) DeleteCategoryById(c echo.Context) error {
	categoryId, parsedCategoryErr := categoryHandler.util.NumericParamConvertor(c.Param("id"), "the provided category id has wrong format")
	if parsedCategoryErr != nil {
		return parsedCategoryErr
	}
	deleteErr := categoryHandler.categoryService.DeleteCategoryById(categoryId)
	if deleteErr != nil {
		return deleteErr
	}
	c.JSON(http.StatusOK, map[string]any{
		"message": "delete successfully",
	})
	return nil
}
