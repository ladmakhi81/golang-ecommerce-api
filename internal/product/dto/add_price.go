package productdto

type AddPriceToProductsPriceListReqBody struct {
	Key        string  `json:"key" validate:"required"`
	Value      string  `json:"value" validate:"required"`
	ExtraPrice float32 `json:"extraPrice" validate:"required,gte=1"`
}
