package vendorincomerepository

import vendorincomeentity "github.com/ladmakhi81/golang-ecommerce-api/internal/vendor_income/entity"

type IVendorIncomeRepository interface {
	CreateIncome(vendorIncome *vendorincomeentity.VendorIncome) error
}
