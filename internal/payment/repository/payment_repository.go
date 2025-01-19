package paymentrepository

import paymententity "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/entity"

type IPaymentRepository interface {
	CreatePayment(payment *paymententity.Payment) error
	FindPaymentByAuthority(authority string) (*paymententity.Payment, error)
	UpdatePaymentStatus(payment *paymententity.Payment) error
	GetPaymentsPage(page, limit uint) ([]*paymententity.Payment, error)
}
