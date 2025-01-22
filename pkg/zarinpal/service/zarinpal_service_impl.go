package pkgzarinpalservice

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	pkgzarinpaldto "github.com/ladmakhi81/golang-ecommerce-api/pkg/zarinpal/dto"
)

type ZarinpalService struct {
	config config.MainConfig
}

func NewZarinpalService(
	config config.MainConfig,
) IZarinpalService {
	return ZarinpalService{
		config: config,
	}
}

func (zarinpalService ZarinpalService) SendRequest(amount float32) (*pkgzarinpaldto.ZarinpalSendRequestResponse, error) {
	reqBody := pkgzarinpaldto.NewZarinpalSendRequestReqBody(
		zarinpalService.GetMerchantID(),
		amount,
		zarinpalService.config.ZarinpalCallbackURL,
	)
	parsedReqBody, parseErr := json.Marshal(reqBody)
	if parseErr != nil {
		return nil, parseErr
	}
	response, reqErr := http.Post(zarinpalService.config.ZarinpalRequestURL, "application/json", bytes.NewBuffer(parsedReqBody))
	if reqErr != nil {
		return nil, reqErr
	}
	defer response.Body.Close()
	responseBody, responseBodyErr := io.ReadAll(response.Body)
	if responseBodyErr != nil {
		return nil, responseBodyErr
	}
	parsedResponseBody := new(pkgzarinpaldto.ZarinpalSendRequestResponse)
	parsedResponseBodyErr := json.Unmarshal(responseBody, parsedResponseBody)
	if parsedResponseBodyErr != nil {
		return nil, parsedResponseBodyErr
	}
	return parsedResponseBody, nil
}
func (zarinpalService ZarinpalService) GetPayLink(authority string) string {
	return zarinpalService.config.ZarinpalPayURL + authority
}
func (zarinpalService ZarinpalService) GetMerchantID() string {
	return zarinpalService.config.ZarinpalMerchantID
}
func (zarinpalService ZarinpalService) VerifyPayment(amount float32, authority string) (*pkgzarinpaldto.VerifyPaymentResponse, error) {
	reqBody := pkgzarinpaldto.NewVerifyPaymentReqBody(
		zarinpalService.GetMerchantID(),
		amount,
		authority,
	)
	parsedReqBody, parseErr := json.Marshal(reqBody)
	if parseErr != nil {
		return nil, parseErr
	}
	response, reqErr := http.Post(zarinpalService.config.ZarinpalVerifyURL, "application/json", bytes.NewBuffer(parsedReqBody))
	if reqErr != nil {
		return nil, reqErr
	}
	defer response.Body.Close()
	responseBody, responseBodyErr := io.ReadAll(response.Body)
	if responseBodyErr != nil {
		return nil, responseBodyErr
	}
	parsedResponseBody := new(pkgzarinpaldto.VerifyPaymentResponse)
	parsedResponseBodyErr := json.Unmarshal(responseBody, parsedResponseBody)
	if parsedResponseBodyErr != nil {
		return nil, parsedResponseBodyErr
	}
	return parsedResponseBody, nil
}
