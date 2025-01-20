package userservice

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	userdto "github.com/ladmakhi81/golang-ecommerce-api/internal/user/dto"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
	userrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/user/repository"
)

type UserAddressService struct {
	userAddressRepo userrepository.IUserAddressRepository
	userService     IUserService
}

func NewUserAddressService(
	userAddressRepo userrepository.IUserAddressRepository,
	userService IUserService,
) UserAddressService {
	return UserAddressService{
		userAddressRepo: userAddressRepo,
		userService:     userService,
	}
}

func (userAddressService UserAddressService) CreateUserAddress(userId uint, reqBody userdto.CreateUserAddressReqBody) (*userentity.UserAddress, error) {
	user, userErr := userAddressService.userService.FindBasicUserInfoById(userId)
	if userErr != nil {
		return nil, userErr
	}
	userAddress := userentity.NewUserAddress(
		reqBody.City,
		reqBody.Province,
		reqBody.Address,
		reqBody.LicensePlate,
		reqBody.Description,
		user,
	)
	if createErr := userAddressService.userAddressRepo.CreateUserAddress(userAddress); createErr != nil {
		return nil, types.NewServerError(
			"error in creating user address",
			"UserAddressService.CreateUserAddress",
			createErr,
		)
	}
	return userAddress, nil
}
func (userAddressService UserAddressService) GetUserAddresses(userId uint) ([]*userentity.UserAddress, error) {
	addresses, addressesErr := userAddressService.userAddressRepo.GetUserAddresses(userId)
	if addressesErr != nil {
		return nil, types.NewServerError(
			"error in finding user addresses",
			"UserAddressService.GetUserAddresses",
			addressesErr,
		)
	}
	return addresses, nil
}
