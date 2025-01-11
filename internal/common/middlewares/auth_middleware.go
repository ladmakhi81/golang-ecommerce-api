package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
)

func AuthMiddleware(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(secretKey),
		SigningMethod: jwt.SigningMethodHS256.Name,
		ContextKey:    "Auth",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &types.AuthClaim{}
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return types.NewClientError("Unauthorized", http.StatusUnauthorized)
		},
	})
}
