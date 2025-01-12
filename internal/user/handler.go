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
	userService userservice.IUserService
	util        utils.Util
}

func NewUserHandler(userService userservice.IUserService) UserHandler {
	return UserHandler{
		userService: userService,
		util:        utils.NewUtil(),
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
