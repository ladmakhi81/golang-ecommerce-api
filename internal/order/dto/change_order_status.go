package orderdto

import (
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
)

type ChangeOrderStatusReqBody struct {
	Status orderentity.OrderStatus `json:"status" validate:"required"`
}
