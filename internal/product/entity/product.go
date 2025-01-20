package productentity

import (
	"time"

	category_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/category/entity"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type Product struct {
	Name          string                    `json:"name,omitempty"`
	Description   string                    `json:"description,omitempty"`
	Category      *category_entity.Category `json:"category,omitempty"`
	Vendor        *user_entity.User         `json:"vendor,omitempty"`
	BasePrice     float32                   `json:"basePrice,omitempty"`
	Tags          []string                  `json:"tags,omitempty"`
	IsConfirmed   bool                      `json:"isConfirmed,omitempty"`
	ConfirmedBy   *user_entity.User         `json:"confirmedBy,omitempty"`
	ConfirmedAt   time.Time                 `json:"confirmedAt,omitempty"`
	ProductPrices []*ProductPrice           `json:"productPrices,omitempty"`
	Fee           float32                   `json:"fee,omitempty"`
	Images        []string                  `json:"images,omitempty"`

	entity.BaseEntity
}

func NewProduct(
	name,
	description string,
	category *category_entity.Category,
	vendor *user_entity.User,
	basePrice float32,
	tags []string,
) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Category:    category,
		Vendor:      vendor,
		BasePrice:   basePrice,
		Tags:        tags,
	}
}
