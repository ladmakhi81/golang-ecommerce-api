package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	responsehandling "github.com/ladmakhi81/golang-ecommerce-api/internal/common/response_handling"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	userdto "github.com/ladmakhi81/golang-ecommerce-api/internal/user/dto"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
)

type UserHandler struct {
	userService        userservice.IUserService
	userAddressService userservice.IUserAddressService
	util               utils.Util
	translation        translations.ITranslation
}

func NewUserHandler(
	userService userservice.IUserService,
	userAddressService userservice.IUserAddressService,
	translation translations.ITranslation,
) UserHandler {
	return UserHandler{
		userService:        userService,
		util:               utils.NewUtil(),
		userAddressService: userAddressService,
		translation:        translation,
	}
}

func (userHandler UserHandler) VerifyAccountByAdmin(c echo.Context) error {
	adminId := c.Get("AuthClaim").(*types.AuthClaim).ID
	vendorId, parsedVendorErr := userHandler.util.NumericParamConvertor(
		c.Param("id"),
		userHandler.translation.Message("user.invalid_id"),
	)
	if parsedVendorErr != nil {
		return parsedVendorErr
	}
	err := userHandler.userService.VerifyAccountByAdmin(adminId, uint(vendorId))
	if err != nil {
		return err
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		nil,
	)
	return nil
}
func (userHandler UserHandler) CompleteProfile(c echo.Context) error {
	auth := c.Get("AuthClaim").(*types.AuthClaim)
	var reqBody userdto.CompleteProfileReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError(
			userHandler.translation.Message("errors.invalid_request_body"),
			http.StatusBadGateway,
		)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	_, updateUserErr := userHandler.userService.CompleteProfile(auth.ID, &reqBody)
	if updateUserErr != nil {
		return updateUserErr
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		nil,
	)
	return nil
}
func (userHandler UserHandler) CreateUserAddress(c echo.Context) error {
	auth := c.Get("AuthClaim").(*types.AuthClaim)
	var reqBody userdto.CreateUserAddressReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError(
			userHandler.translation.Message("errors.invalid_request_body"),
			http.StatusBadGateway,
		)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	createdAddress, createdAddressErr := userHandler.userAddressService.CreateUserAddress(auth.ID, reqBody)
	if createdAddressErr != nil {
		return createdAddressErr
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusCreated,
		createdAddress,
	)
	return nil
}
func (userHandler UserHandler) GetUserAddresses(c echo.Context) error {
	auth := c.Get("AuthClaim").(*types.AuthClaim)
	addresses, err := userHandler.userAddressService.GetUserAddresses(auth.ID)
	if err != nil {
		return err
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		addresses,
	)
	return nil
}
func (userHandler UserHandler) AssignActiveAddressUser(c echo.Context) error {
	auth := c.Get("AuthClaim").(*types.AuthClaim)
	var reqBody userdto.AssignActiveAddressUserReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError(
			userHandler.translation.Message("errors.invalid_request_body"),
			http.StatusBadGateway,
		)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	if err := userHandler.userAddressService.AssignUserAddress(auth.ID, reqBody); err != nil {
		return err
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		nil,
	)
	return nil
}
