package productdto

type CreateProductReqBody struct {
	Name        string   `json:"name" validate:"required,min=3"`
	Description string   `json:"description" validate:"required,min=3"`
	CategoryID  uint     `json:"categoryId" validate:"required,gte=1"`
	BasePrice   float32  `json:"basePrice" validate:"required,gte=1"`
	Tags        []string `json:"tags"`
}
