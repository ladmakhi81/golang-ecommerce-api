package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type JwtService struct {
	config config.MainConfig
}

func NewJwtService(config config.MainConfig) IJwtService {
	return JwtService{
		config,
	}
}

func (jwtService JwtService) getSecretKey() []byte {
	return []byte(jwtService.config.SecretKey)
}

func (jwtService JwtService) GenerateAccessToken(user *userentity.User) (string, error) {
	claim := types.NewAuthClaim(user.ID, user.Role)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	accessToken, accessTokenErr := token.SignedString(jwtService.getSecretKey())
	if accessTokenErr != nil {
		return "", types.NewServerError(
			"error in generating access token",
			"JwtService.GenerateAccessToken",
			accessTokenErr,
		)
	}
	return accessToken, nil
}
