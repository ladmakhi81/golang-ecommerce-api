package product_entity

import (
	"time"

	category_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/category/entity"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type Product struct {
	Name          string
	Description   string
	Category      *category_entity.Category
	Vendor        *user_entity.User
	BasePrice     float32
	Tags          []string
	IsConfirmed   bool
	ConfirmedBy   *user_entity.User
	ConfirmedAt   time.Time
	ProductPrices []*ProductPrice

	entity.BaseEntity
}
