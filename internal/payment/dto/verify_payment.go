package paymentdto

type VerifyPaymentReqBody struct {
	Authority string `json:"authority" validate:"required"`
	Status    string `json:"status" validate:"required"`
}
