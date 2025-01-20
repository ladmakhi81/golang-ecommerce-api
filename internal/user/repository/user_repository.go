package userrepository

import userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"

type IUserRepository interface {
	CreateUser(user *userentity.User) error
	IsEmailExist(email string) (bool, error)
	FindBasicInfoByEmail(email string) (*userentity.User, error)
	CompleteProfile(user *userentity.User) error
	FindBasicUserInfoById(id uint) (*userentity.User, error)
	UpdateVerificationState(adminId uint, vendorId uint) error
	SetActiveUserAddress(userId uint, addressId uint) error
}
