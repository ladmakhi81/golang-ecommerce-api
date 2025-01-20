package userdto

type CreateUserAddressReqBody struct {
	City         string `json:"city" validate:"required,alpha"`
	Province     string `json:"province" validate:"required,alpha"`
	Address      string `json:"address" validate:"required"`
	LicensePlate string `json:"licensePlate" validate:"required"`
	Description  string `json:"description,alpha"`
}
