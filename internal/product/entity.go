package product

import (
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/category"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/user"
)

type Product struct {
	Name          string
	Description   string
	Category      *category.Category
	Vendor        *user.User
	BasePrice     float32
	Tags          []string
	IsConfirmed   bool
	ConfirmedBy   *user.User
	ConfirmedAt   time.Time
	ProductPrices []*ProductPrice

	entity.BaseEntity
}

type ProductPrice struct {
	Key        string
	Value      string
	ExtraPrice float32

	entity.BaseEntity
}
