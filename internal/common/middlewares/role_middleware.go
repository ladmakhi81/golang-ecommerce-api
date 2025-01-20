package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

func RoleMiddleware(allowedRole userentity.UserRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole := c.Get("AuthClaim").(*types.AuthClaim).Role
			if userRole != "admin" && userRole != allowedRole {
				return types.NewClientError(
					"Access Denied",
					http.StatusForbidden,
				)
			}
			return next(c)
		}
	}
}
