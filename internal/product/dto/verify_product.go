package productdto

type VerifyProductReqBody struct {
	Fee float32 `json:"fee" validate:"required,gte=1"`
}
