package productentity

import "github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"

type ProductPrice struct {
	Key        string  `json:"key,omitempty"`
	Value      string  `json:"value,omitempty"`
	ExtraPrice float32 `json:"extraPrice,omitempty"`
	ProductID  uint    `json:"productId,omitempty"`

	entity.BaseEntity
}

func NewProductPrice(key, value string, extraPrice float32, productId uint) *ProductPrice {
	return &ProductPrice{
		Key:        key,
		Value:      value,
		ExtraPrice: extraPrice,
		ProductID:  productId,
	}
}
