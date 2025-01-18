package pkgzarinpaldto

type VerifyPaymentReqBody struct {
	MerchantID string  `json:"merchant_id"`
	Amount     float32 `json:"amount"`
	Authority  string  `json:"authority"`
}

func NewVerifyPaymentReqBody(merchantID string, amount float32, authority string) VerifyPaymentReqBody {
	return VerifyPaymentReqBody{
		MerchantID: merchantID,
		Amount:     amount,
		Authority:  authority,
	}
}

type VerifyPaymentResponseDetail struct {
	RefID uint `json:"ref_id"`
}

type VerifyPaymentResponse struct {
	Data VerifyPaymentResponseDetail `json:"data"`
}
