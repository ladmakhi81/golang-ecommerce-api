package userrepository

import userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"

type IUserAddressRepository interface {
	CreateUserAddress(userAddress *userentity.UserAddress) error
	GetUserAddresses(userId uint) ([]*userentity.UserAddress, error)
}
