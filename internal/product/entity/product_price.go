package productentity

import "github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"

type ProductPrice struct {
	Key        string  `json:"key"`
	Value      string  `json:"value"`
	ExtraPrice float32 `json:"extraPrice"`

	entity.BaseEntity
}
