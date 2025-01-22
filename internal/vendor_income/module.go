package vendorincome

import (
	"fmt"

	"github.com/labstack/echo/v4"
	vendorincomeevent "github.com/ladmakhi81/golang-ecommerce-api/internal/vendor_income/event"
	vendorincomerepository "github.com/ladmakhi81/golang-ecommerce-api/internal/vendor_income/repository"
	vendorincomeservice "github.com/ladmakhi81/golang-ecommerce-api/internal/vendor_income/service"
	"go.uber.org/dig"
)

type VendorIncomeModule struct {
	diContainer *dig.Container
	baseApi     *echo.Group
}

func NewVendorIncomeModule(
	diContainer *dig.Container,
	baseApi *echo.Group,
) VendorIncomeModule {
	return VendorIncomeModule{
		diContainer: diContainer,
		baseApi:     baseApi,
	}
}

func (vendorIncomeModule VendorIncomeModule) LoadModule() {
	vendorIncomeModule.diContainer.Provide(vendorincomeservice.NewVendorIncomeService)
	vendorIncomeModule.diContainer.Provide(vendorincomerepository.NewVendorIncomeRepository)
	vendorIncomeModule.diContainer.Provide(vendorincomeevent.NewVendorIncomeEventsSubscriber)
	vendorIncomeModule.diContainer.Provide(vendorincomeevent.NewVendorIncomeEventsContainer)
}

func (vendorIncomeModule VendorIncomeModule) Run() {
	vendorIncomeEventErr := vendorIncomeModule.diContainer.Invoke(func(vendorIncomeEventContainer vendorincomeevent.VendorIncomeEventsContainer) {
		vendorIncomeEventContainer.RegisterEvents()
	})

	if vendorIncomeEventErr == nil {
		fmt.Println("VendorIncomeModule Loaded Successfully")
	} else {
		fmt.Println("VendorIncomeModule Not Load", vendorIncomeEventErr)
	}
}
