package userservice

import userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"

type IUserService interface {
	CheckDuplicatedUserEmail(email string) error
	CreateBasicUser(email, password string, role userentity.UserRole) (*userentity.User, error)
	FindUserByEmailAndPassword(email, password string) (*userentity.User, error)
	FindUserByEmail(email string) (*userentity.User, error)
}
