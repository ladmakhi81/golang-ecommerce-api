package userentity

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

func NewUser(email string, password string, role UserRole) *User {
	return &User{
		Role:     role,
		Email:    email,
		Password: password,
	}
}
