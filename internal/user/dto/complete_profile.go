package userdto

type CompleteProfileReqBody struct {
	FullName   string `json:"fullName" validation:"required,min=3"`
	NationalID string `json:"nationalID" validation:"required,min=11"`
	PostalCode string `json:"postalCode" validation:"required"`
	Address    string `json:"address" validation:"required"`
}
