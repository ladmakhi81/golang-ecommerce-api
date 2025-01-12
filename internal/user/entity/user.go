package userentity

import (
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
)

type User struct {
	Role UserRole `json:"role,omitempty"`

	Email    string `json:"email,omitempty"`
	Password string `json:"-"`

	// vendor
	FullName          string    `json:"fullName,omitempty"`
	NationalID        string    `json:"nationalId,omitempty"`
	PostalCode        string    `json:"postalCode,omitempty"`
	Address           string    `json:"address,omitempty"`
	IsCompleteProfile bool      `json:"isCompleteProfile,omitempty"`
	CompleteProfileAt time.Time `json:"completeProfileAt,omitempty"`
	IsVerified        bool      `json:"isVerified,omitempty"`
	VerifiedBy        *User     `json:"verifiedBy,omitempty"`
	VerifiedDate      time.Time `json:"verified_date,omitempty"`

	entity.BaseEntity
}

func NewUser(email string, password string, role UserRole) *User {
	return &User{
		Role:     role,
		Email:    email,
		Password: password,
	}
}
