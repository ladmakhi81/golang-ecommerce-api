package paymentservice

import (
	"net/http"
	"strings"
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	paymentdto "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/dto"
	paymententity "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/entity"
	paymentrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/repository"
	transactionservice "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/service"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
	pkgzarinpalservice "github.com/ladmakhi81/golang-ecommerce-api/pkg/zarinpal/service"
)

type PaymentService struct {
	paymentRepo        paymentrepository.IPaymentRepository
	zarinpalService    pkgzarinpalservice.IZarinpalService
	transactionService transactionservice.ITransactionService
	eventsContainer    *events.EventsContainer
	translation        translations.ITranslation
}

func NewPaymentService(
	paymentRepo paymentrepository.IPaymentRepository,
	zarinpalService pkgzarinpalservice.IZarinpalService,
	transactionService transactionservice.ITransactionService,
	eventsContainer *events.EventsContainer,
	translation translations.ITranslation,
) IPaymentService {
	return PaymentService{
		paymentRepo:        paymentRepo,
		zarinpalService:    zarinpalService,
		transactionService: transactionService,
		eventsContainer:    eventsContainer,
		translation:        translation,
	}
}

func (paymentService PaymentService) CreatePayment(order *orderentity.Order) (*paymententity.Payment, error) {
	zarinpalRes, zarinpalErr := paymentService.zarinpalService.SendRequest(order.FinalPrice)
	if zarinpalErr != nil {
		return nil, types.NewServerError(
			"error in sending request for zarinpal",
			"PaymentService.ZarinpalService.SendRequest",
			zarinpalErr,
		)
	}
	payment := paymententity.NewPayment(
		order,
		zarinpalRes.Data.Authority,
		paymentService.zarinpalService.GetMerchantID(),
	)
	paymentErr := paymentService.paymentRepo.CreatePayment(payment)
	if paymentErr != nil {
		return nil, types.NewServerError(
			"error in creating payment",
			"PaymentService.CreatePayment",
			paymentErr,
		)
	}
	return payment, nil
}
func (paymentService PaymentService) GetPayLink(payment *paymententity.Payment) string {
	return paymentService.zarinpalService.GetPayLink(payment.Authority)
}
func (paymentService PaymentService) VerifyPayment(customerId uint, reqBody paymentdto.VerifyPaymentReqBody) error {
	payment, paymentErr := paymentService.paymentRepo.FindPaymentByAuthority(reqBody.Authority)
	if paymentErr != nil {
		return types.NewServerError(
			"error in finding payments by authority",
			"PaymentService.VerifyPayment.FindPaymentByAuthority",
			paymentErr,
		)
	}
	if payment == nil {
		return types.NewClientError(
			paymentService.translation.Message("payment.not_found_authority"),
			http.StatusNotFound,
		)
	}
	if payment.Status != paymententity.PaymentStatusPending {
		return types.NewClientError(
			paymentService.translation.Message("payment.verified_before"),
			http.StatusBadRequest,
		)
	}
	if payment.Customer.ID != customerId {
		return types.NewClientError(
			paymentService.translation.Message("payment.owner_verified_err"),
			http.StatusForbidden,
		)
	}
	if strings.ToLower(reqBody.Status) != "ok" {
		payment.Status = paymententity.PaymentStatusFailed
		payment.StatusChangedAt = time.Now()
	} else {
		verifyRes, verifyErr := paymentService.zarinpalService.VerifyPayment(payment.Amount, payment.Authority)
		if verifyErr != nil || verifyRes == nil {
			return types.NewServerError(
				"error in verifying the payment",
				"PaymentService.VerifyPayment",
				verifyErr,
			)
		}
		refId := verifyRes.Data.RefID
		payment.Status = paymententity.PaymentStatusSuccess
		payment.StatusChangedAt = time.Now()
		customerTransaction, customerTransactionErr := paymentService.transactionService.CreateTransaction(
			payment,
			refId,
			payment.Customer,
		)
		if customerTransactionErr != nil {
			return customerTransactionErr
		}
		paymentService.eventsContainer.PublishEvent(
			events.NewEvent(
				events.CALCULATE_VENDOR_INCOME_EVENT,
				events.NewCalculateVendorIncomeEventBody(customerTransaction),
			),
		)
		paymentService.eventsContainer.PublishEvent(
			events.NewEvent(
				events.CHANGED_ORDER_STATUS_EVENT,
				events.NewChangeOrderStatusEventBody(payment.Order, payment.Customer),
			),
		)

	}
	if updateErr := paymentService.paymentRepo.UpdatePaymentStatus(payment); updateErr != nil {
		return types.NewServerError(
			"error in update payment status",
			"PaymentService.VerifyPayment.UpdatePaymentStatus",
			updateErr,
		)
	}
	return nil
}
func (paymentService PaymentService) GetPaymentsPage(page, limit uint) ([]*paymententity.Payment, uint, error) {
	payments, paymentsErr := paymentService.paymentRepo.GetPaymentsPage(page, limit)
	if paymentsErr != nil {
		return nil, 0, types.NewServerError(
			"error in finding payments page",
			"PaymentService.GetPaymentsPage",
			paymentsErr,
		)
	}
	paymentsCount, paymentsCountErr := paymentService.paymentRepo.GetPaymentsCount()
	if paymentsCountErr != nil {
		return nil, 0, types.NewServerError(
			"error in get count of payments",
			"PaymentService.GetPaymentsPage.GetPaymentsCount",
			paymentsCountErr,
		)
	}
	return payments, paymentsCount, nil
}
