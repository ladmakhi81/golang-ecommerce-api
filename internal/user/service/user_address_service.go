package userservice

import (
	userdto "github.com/ladmakhi81/golang-ecommerce-api/internal/user/dto"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type IUserAddressService interface {
	CreateUserAddress(userId uint, reqBody userdto.CreateUserAddressReqBody) (*userentity.UserAddress, error)
	GetUserAddresses(userId uint) ([]*userentity.UserAddress, error)
	AssignUserAddress(userId uint, reqBody userdto.AssignActiveAddressUserReqBody) error
	FindAddressById(addressId uint) (*userentity.UserAddress, error)
}
