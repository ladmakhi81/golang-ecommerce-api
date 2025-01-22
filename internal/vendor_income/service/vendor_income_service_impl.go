package vendorincomeservice

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	orderservice "github.com/ladmakhi81/golang-ecommerce-api/internal/order/service"
	transactionentity "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/entity"
	transactionservice "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/service"
	vendorincomeentity "github.com/ladmakhi81/golang-ecommerce-api/internal/vendor_income/entity"
	vendorincomerepository "github.com/ladmakhi81/golang-ecommerce-api/internal/vendor_income/repository"
)

type VendorIncomeService struct {
	vendorIncomeRepo   vendorincomerepository.IVendorIncomeRepository
	orderService       orderservice.IOrderService
	transactionService transactionservice.ITransactionService
}

func NewVendorIncomeService(
	vendorIncomeRepo vendorincomerepository.IVendorIncomeRepository,
	orderService orderservice.IOrderService,
	transactionService transactionservice.ITransactionService,
) IVendorIncomeService {
	return VendorIncomeService{
		vendorIncomeRepo:   vendorIncomeRepo,
		orderService:       orderService,
		transactionService: transactionService,
	}
}

func (vendorIncomeService VendorIncomeService) CreateVendorIncome(transaction *transactionentity.Transaction) error {
	orderId, orderIdErr := vendorIncomeService.transactionService.GetOrderIdOfTransaction(transaction.ID)
	if orderIdErr != nil {
		return orderIdErr
	}
	items, itemsErr := vendorIncomeService.orderService.FindOrderItemsByOrderId(*orderId)
	if itemsErr != nil {
		return itemsErr
	}
	for _, item := range items {
		vendorIncome := vendorincomeentity.NewVendorIncome(
			item.Customer,
			item.Order.FinalPrice,
			item.Product.Fee,
			item.Order.FinalPrice-item.Product.Fee,
			item,
			transaction,
		)
		if err := vendorIncomeService.vendorIncomeRepo.CreateIncome(vendorIncome); err != nil {
			return types.NewServerError(
				"error in creating vendor income",
				"VendorIncomeService.CreateVendorIncome.CreateIncome",
				err,
			)
		}
	}
	return nil
}
