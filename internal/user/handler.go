package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	userdto "github.com/ladmakhi81/golang-ecommerce-api/internal/user/dto"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
)

type UserHandler struct {
	userService        userservice.IUserService
	userAddressService userservice.IUserAddressService
	util               utils.Util
}

func NewUserHandler(
	userService userservice.IUserService,
	userAddressService userservice.IUserAddressService,
) UserHandler {
	return UserHandler{
		userService:        userService,
		util:               utils.NewUtil(),
		userAddressService: userAddressService,
	}
}

func (userHandler UserHandler) VerifyAccountByAdmin(c echo.Context) error {
	adminId := c.Get("AuthClaim").(*types.AuthClaim).ID
	vendorId, parsedVendorErr := userHandler.util.NumericParamConvertor(c.Param("id"), "the provided id has wrong format")
	if parsedVendorErr != nil {
		return parsedVendorErr
	}
	err := userHandler.userService.VerifyAccountByAdmin(adminId, uint(vendorId))
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, map[string]string{"message": "verify successfully ..."})
	return nil
}
func (userHandler UserHandler) CompleteProfile(c echo.Context) error {
	auth := c.Get("AuthClaim").(*types.AuthClaim)
	var reqBody userdto.CompleteProfileReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError("invalid request body", http.StatusBadGateway)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	_, updateUserErr := userHandler.userService.CompleteProfile(auth.ID, &reqBody)
	if updateUserErr != nil {
		return updateUserErr
	}
	c.JSON(http.StatusOK, map[string]any{
		"data":    map[string]string{"message": "profile completed"},
		"success": true,
	})
	return nil
}
func (userHandler UserHandler) CreateUserAddress(c echo.Context) error {
	auth := c.Get("AuthClaim").(*types.AuthClaim)
	var reqBody userdto.CreateUserAddressReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError("invalid request body", http.StatusBadRequest)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	createdAddress, createdAddressErr := userHandler.userAddressService.CreateUserAddress(auth.ID, reqBody)
	if createdAddressErr != nil {
		return createdAddressErr
	}
	c.JSON(http.StatusCreated, map[string]any{"success": true, "data": createdAddress})
	return nil
}
func (userHandler UserHandler) GetUserAddresses(c echo.Context) error {
	auth := c.Get("AuthClaim").(*types.AuthClaim)
	addresses, err := userHandler.userAddressService.GetUserAddresses(auth.ID)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, map[string]any{"data": addresses})
	return nil
}
