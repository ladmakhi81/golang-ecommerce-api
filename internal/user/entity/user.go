package user_entity

import "github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"

type User struct {
	Role UserRole

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
