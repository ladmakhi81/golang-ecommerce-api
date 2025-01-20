package userservice

import (
	userdto "github.com/ladmakhi81/golang-ecommerce-api/internal/user/dto"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type IUserService interface {
	CheckDuplicatedUserEmail(email string) error
	CreateBasicUser(email, password string, role userentity.UserRole) (*userentity.User, error)
	FindUserByEmailAndPassword(email, password string) (*userentity.User, error)
	FindUserByEmail(email string) (*userentity.User, error)
	CompleteProfile(userId uint, data *userdto.CompleteProfileReqBody) (*userentity.User, error)
	FindBasicUserInfoById(id uint) (*userentity.User, error)
	VerifyAccountByAdmin(adminId uint, vendorId uint) error
	SetActiveUserAddress(userId uint, addressId uint) error
}
