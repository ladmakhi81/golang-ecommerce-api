package productservice

import (
	"net/http"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	productdto "github.com/ladmakhi81/golang-ecommerce-api/internal/product/dto"
	productentity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
	productrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/product/repository"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
)

type ProductPriceService struct {
	productService   ProductService
	productPriceRepo productrepository.IProductPriceRepository
	translation      translations.ITranslation
}

func NewProductPriceService(
	productService ProductService,
	productPriceRepo productrepository.IProductPriceRepository,
	translation translations.ITranslation,
) ProductPriceService {
	return ProductPriceService{
		productService:   productService,
		productPriceRepo: productPriceRepo,
		translation:      translation,
	}
}

func (productPriceService ProductPriceService) AddPriceToProductsPriceList(reqBody productdto.AddPriceToProductsPriceListReqBody, productId uint, assignerId uint) (*productentity.ProductPrice, error) {
	product, productErr := productPriceService.productService.FindProductById(productId)
	if productErr != nil {
		return nil, productErr
	}
	if product.Vendor.ID != assignerId {
		return nil, types.NewClientError(
			productPriceService.translation.Message("product.price_item_owner_add"),
			http.StatusForbidden,
		)
	}
	priceItem := productentity.NewProductPrice(reqBody.Key, reqBody.Value, reqBody.ExtraPrice, product.ID)
	if createErr := productPriceService.productPriceRepo.CreateProductPrice(priceItem); createErr != nil {
		return nil, types.NewServerError(
			"error in create price item",
			"ProductPriceService.AddPriceToProductPriceList.CreateProductPrice",
			createErr,
		)
	}
	return priceItem, nil
}
func (productPriceService ProductPriceService) RemovePriceItemFromProductList(priceItem uint) error {
	isExist, existenceErr := productPriceService.IsProductPriceItemExist(priceItem)
	if existenceErr != nil {
		return existenceErr
	}
	if !isExist {
		return types.NewClientError(
			productPriceService.translation.Message("product.price_item_not_found_id"),
			http.StatusNotFound,
		)
	}
	deleteErr := productPriceService.productPriceRepo.DeleteProductPriceById(priceItem)
	if deleteErr != nil {
		return types.NewServerError(
			"error in deleting price",
			"ProductPriceService.RemovePriceItemFromProductList",
			deleteErr,
		)
	}
	return nil
}
func (productPriceService ProductPriceService) IsProductPriceItemExist(priceItem uint) (bool, error) {
	isExist, existenceErr := productPriceService.productPriceRepo.IsProductPriceItemExist(priceItem)
	if existenceErr != nil {
		return false, types.NewServerError(
			"error in finding product price item",
			"ProductPriceService.IsProductPriceItemExist",
			existenceErr,
		)
	}
	return isExist, nil
}
func (productPriceService ProductPriceService) FindPriceItemById(priceItemID uint) (*productentity.ProductPrice, error) {
	priceItem, priceItemErr := productPriceService.productPriceRepo.FindPriceItemById(priceItemID)
	if priceItemErr != nil {
		return nil, types.NewServerError(
			"error in finding price item",
			"ProductPriceService.FindPriceItemById",
			priceItemErr,
		)
	}
	if priceItem == nil {
		return nil, types.NewClientError(
			productPriceService.translation.Message("product.price_item_not_found_id"),
			http.StatusNotFound,
		)
	}
	return priceItem, nil
}
func (productPriceService ProductPriceService) FindPricesByProductId(productId uint) (*[]productentity.ProductPrice, error) {
	prices, priceErr := productPriceService.productPriceRepo.FindPricesByProductId(productId)
	if priceErr != nil {
		return nil, types.NewServerError(
			"error in finding prices based on product id",
			"ProductPriceService.productPriceRepo.FindPricesByProductId",
			priceErr,
		)
	}
	return prices, nil
}
