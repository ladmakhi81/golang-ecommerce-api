package userentity

import "github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"

type UserAddress struct {
	City         string `json:"city"`
	Province     string `json:"province"`
	Address      string `json:"address"`
	LicensePlate string `json:"licensePlate"`
	Description  string `json:"description"`
	User         *User  `json:"user,omitempty"`

	entity.BaseEntity
}

func NewUserAddress(
	city string,
	province string,
	address string,
	licensePlate string,
	description string,
	user *User,
) *UserAddress {
	return &UserAddress{
		City:         city,
		Province:     province,
		Address:      address,
		LicensePlate: licensePlate,
		Description:  description,
		User:         user,
	}
}
