package cartdto

type AddProductCartReqBody struct {
	ProductID   uint `json:"productId" validate:"required,numeric"`
	Quantity    uint `json:"quantity" validate:"required,gte=1"`
	PriceItemID uint `json:"priceItemId" validate:"required,gte=1"`
}
