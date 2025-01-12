package utils

import (
	"net/http"
	"strconv"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
)

func (Util) NumericParamConvertor(paramValue, errorMessage string) (uint, error) {
	value, parsedErr := strconv.Atoi(paramValue)
	if parsedErr != nil {
		return 0, types.NewClientError(errorMessage, http.StatusBadRequest)
	}
	return uint(value), nil
}
