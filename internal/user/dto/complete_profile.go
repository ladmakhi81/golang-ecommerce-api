package userdto

type CompleteProfileReqBody struct {
	FullName   string `json:"fullName" validate:"required,min=3"`
	NationalID string `json:"nationalID" validate:"required,min=11"`
	PostalCode string `json:"postalCode" validate:"required"`
	Address    string `json:"address" validate:"required"`
}
