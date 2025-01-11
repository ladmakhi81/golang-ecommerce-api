package categorydto

type CreateCategoryReqBody struct {
	Name             string `json:"name" validate:"required,min=3"`
	Icon             string `json:"icon" validate:"required,min=3"`
	ParentCategoryId uint   `json:"parentCategoryId"`
}
