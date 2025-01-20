package userservice

import (
	"net/http"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	userdto "github.com/ladmakhi81/golang-ecommerce-api/internal/user/dto"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
	userrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/user/repository"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
)

type UserAddressService struct {
	userAddressRepo userrepository.IUserAddressRepository
	userService     IUserService
	translation     translations.ITranslation
}

func NewUserAddressService(
	userAddressRepo userrepository.IUserAddressRepository,
	userService IUserService,
	translation translations.ITranslation,
) UserAddressService {
	return UserAddressService{
		userAddressRepo: userAddressRepo,
		userService:     userService,
		translation:     translation,
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
func (userAddressService UserAddressService) AssignUserAddress(userId uint, reqBody userdto.AssignActiveAddressUserReqBody) error {
	address, addressErr := userAddressService.FindAddressById(reqBody.AddressId)
	if addressErr != nil {
		return addressErr
	}
	if err := userAddressService.userService.SetActiveUserAddress(userId, address.ID); err != nil {
		return err
	}
	return nil
}
func (userAddressService UserAddressService) FindAddressById(addressId uint) (*userentity.UserAddress, error) {
	address, addressErr := userAddressService.userAddressRepo.FindAddressById(addressId)
	if addressErr != nil {
		return nil, types.NewServerError(
			"error in finding address by address id",
			"UserAddressService.FindAddressById",
			addressErr,
		)
	}
	if address == nil {
		return nil, types.NewClientError(
			userAddressService.translation.Message("user.user_address_not_found_id"),
			http.StatusNotFound,
		)
	}
	return address, nil
}
