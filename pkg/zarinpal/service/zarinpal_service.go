package pkgzarinpalservice

import pkgzarinpaldto "github.com/ladmakhi81/golang-ecommerce-api/pkg/zarinpal/dto"

type IZarinpalService interface {
	VerifyPayment(amount float32, authority string) (*pkgzarinpaldto.VerifyPaymentResponse, error)
	SendRequest(amount float32) (*pkgzarinpaldto.ZarinpalSendRequestResponse, error)
	GetPayLink(authority string) string
	GetMerchantID() string
}
