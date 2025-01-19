package vendorincomeevent

import (
	"fmt"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
	vendorincomeservice "github.com/ladmakhi81/golang-ecommerce-api/internal/vendor_income/service"
)

type VendorIncomeEventsSubscriber struct {
	vendorIncomeService vendorincomeservice.IVendorIncomeService
}

func NewVendorIncomeEventsSubscriber(
	vendorIncomeService vendorincomeservice.IVendorIncomeService,
) VendorIncomeEventsSubscriber {
	return VendorIncomeEventsSubscriber{
		vendorIncomeService: vendorIncomeService,
	}
}

func (subscriber VendorIncomeEventsSubscriber) SubscribeCalculateVendorIncome(event events.Event) {
	payload := event.Payload.(events.CalculateVendorIncomeEventBody)
	err := subscriber.vendorIncomeService.CreateVendorIncome(
		payload.Transaction,
	)
	if err != nil {
		fmt.Println("vendor income fails", err)
	}
}
