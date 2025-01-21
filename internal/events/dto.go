package events

import (
	"time"

	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	productentity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
	transactionentity "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/entity"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type CalculateVendorIncomeEventBody struct {
	Transaction *transactionentity.Transaction
}

func NewCalculateVendorIncomeEventBody(
	transaction *transactionentity.Transaction,
) CalculateVendorIncomeEventBody {
	return CalculateVendorIncomeEventBody{
		Transaction: transaction,
	}
}

type UserRegisteredEventBody struct {
	Email string
}

func NewUserRegisteredEventBody(email string) UserRegisteredEventBody {
	return UserRegisteredEventBody{Email: email}
}

type ChangeOrderStatusEventBody struct {
	Order    *orderentity.Order
	Customer *userentity.User
}

func NewChangeOrderStatusEventBody(order *orderentity.Order, customer *userentity.User) ChangeOrderStatusEventBody {
	return ChangeOrderStatusEventBody{
		Order:    order,
		Customer: customer,
	}
}

type ProductCreatedEventBody struct {
	Product *productentity.Product
}

func NewProductCreatedEventBody(
	product *productentity.Product,
) ProductCreatedEventBody {
	return ProductCreatedEventBody{
		Product: product,
	}
}

type ProductVerifiedEventBody struct {
	Product *productentity.Product
}

func NewProductVerifiedEventBody(
	product *productentity.Product,
) ProductVerifiedEventBody {
	return ProductVerifiedEventBody{
		Product: product,
	}
}

type UserCompleteProfileEventBody struct {
	Email string
}

func NewUserCompleteProfileEventBody(email string) UserCompleteProfileEventBody {
	return UserCompleteProfileEventBody{Email: email}
}

type UserVerificationEventBody struct {
	AdminEmail     string
	VendorEmail    string
	VendorFullName string
	Date           time.Time
}

func NewUserVerificationEventBody(
	adminEmail,
	vendorEmail,
	vendorFullName string,
	date time.Time,
) UserVerificationEventBody {
	return UserVerificationEventBody{
		AdminEmail:  adminEmail,
		VendorEmail: vendorEmail,
		Date:        date,
	}
}
