package errorhandling

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
)

func GlobalErrorHandling(err error, c echo.Context) {
	if errors, ok := err.(types.ClientValidationError); ok {
		c.JSON(
			http.StatusBadRequest,
			map[string]any{
				"message": "input validation failed",
				"success": false,
				"errors":  errors.Errors,
			},
		)
	}

	if clientErr, ok := err.(types.ClientError); ok {
		c.JSON(
			clientErr.StatusCode,
			map[string]any{
				"success": false,
				"message": clientErr.Message,
			},
		)
	}

	if serverErrs, ok := err.(types.ServerError); ok {
		c.JSON(
			http.StatusInternalServerError,
			map[string]any{
				"message": "something went wrong",
				"success": false,
			},
		)

		fmt.Println(serverErrs)
	}
}
