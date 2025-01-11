package authservice

import userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"

type IJwtService interface {
	GenerateAccessToken(user *userentity.User) (string, error)
}
