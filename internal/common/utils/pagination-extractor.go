package utils

import (
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
)

func (Util) PaginationExtractor(c echo.Context) types.Pagination {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
	page := uint(0)
	limit := uint(10)
	if parsedPage, parsedErr := strconv.Atoi(pageParam); parsedErr == nil {
		page = uint(parsedPage)
	}
	if parsedLimit, parsedErr := strconv.Atoi(limitParam); parsedErr == nil {
		limit = uint(parsedLimit)
	}
	return types.NewPagination(uint(math.Ceil(float64(page)*float64(limit))), limit)
}
