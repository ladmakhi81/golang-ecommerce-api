package responsehandling

import (
	"time"

	"github.com/labstack/echo/v4"
)

func ResponseJSON(c echo.Context, statusCode int, data any) {
	c.JSON(statusCode, map[string]any{
		"success":   true,
		"timestamp": time.Now(),
		"data":      data,
	})
}
