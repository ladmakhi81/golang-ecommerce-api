package category

import (
	"net/http"

	"github.com/labstack/echo/v4"
	categorydto "github.com/ladmakhi81/golang-ecommerce-api/internal/category/dto"
	categoryservice "github.com/ladmakhi81/golang-ecommerce-api/internal/category/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
)

type CategoryHandler struct {
	categoryService categoryservice.ICategoryService
}

func NewCategoryHandler(categoryService categoryservice.ICategoryService) CategoryHandler {
	return CategoryHandler{
		categoryService,
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
