package userdto

type AssignActiveAddressUserReqBody struct {
	AddressId uint `json:"addressId" validate:"required,gte=1"`
}
