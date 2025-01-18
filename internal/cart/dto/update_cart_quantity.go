package cartdto

type UpdateCartQuantityReqBody struct {
	Quantity uint `json:"quantity" validate:"required,gte=1"`
}
