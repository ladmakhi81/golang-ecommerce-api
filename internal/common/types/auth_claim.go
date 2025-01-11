package types

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type AuthClaim struct {
	ID   uint
	Role userentity.UserRole
	jwt.RegisteredClaims
}

func NewAuthClaim(id uint, role userentity.UserRole) AuthClaim {
	return AuthClaim{
		ID:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}
}
