package user

import "github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"

type User struct {
	Email    string
	Password string

	// vendor
	FullName   string
	NationalID string
	PostalCode string
	Address    string
	IsVerified bool
	VerifiedBy *User

	entity.BaseEntity
}
