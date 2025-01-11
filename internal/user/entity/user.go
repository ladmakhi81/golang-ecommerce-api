package userentity

import (
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
)

type User struct {
	Role UserRole `json:"role"`

	Email    string `json:"email"`
	Password string `json:"-"`

	// vendor
	FullName          string    `json:"fullName"`
	NationalID        string    `json:"nationalId"`
	PostalCode        string    `json:"postalCode"`
	Address           string    `json:"address"`
	IsCompleteProfile bool      `json:"isCompleteProfile"`
	CompleteProfileAt time.Time `json:"completeProfileAt"`
	IsVerified        bool      `json:"isVerified"`
	VerifiedBy        *User     `json:"verifiedBy"`
	VerifiedDate      time.Time `json:"verified_date"`

	entity.BaseEntity
}

func NewUser(email string, password string, role UserRole) *User {
	return &User{
		Role:     role,
		Email:    email,
		Password: password,
	}
}
