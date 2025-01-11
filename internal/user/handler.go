package user

import "github.com/labstack/echo/v4"

type UserHandler struct{}

func NewUserHandler() UserHandler {
	return UserHandler{}
}

func (userHandler UserHandler) VerifyAccountByAdmin(c echo.Context) error {
	c.JSON(200, map[string]string{"message": "verify account by admin"})
	return nil
}

func (userHandler UserHandler) CompleteProfile(c echo.Context) error {
	c.JSON(200, map[string]string{"message": "complete profile"})
	return nil
}
