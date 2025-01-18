package pkgzarinpaldto

type ZarinpalSendRequestReqBody struct {
	MerchantID  string  `json:"merchant_id"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	CallbackURL string  `json:"callback_url"`
}

type ZarinpalSendRequestResponseDetail struct {
	Authority string `json:"authority"`
}

type ZarinpalSendRequestResponse struct {
	Data ZarinpalSendRequestResponseDetail `json:"data"`
}

func NewZarinpalSendRequestReqBody(
	merchantID string,
	amount float32,
	callbackURL string,
) ZarinpalSendRequestReqBody {
	return ZarinpalSendRequestReqBody{
		Amount:      amount,
		MerchantID:  merchantID,
		CallbackURL: callbackURL,
		Description: "ecommerce",
	}
}
