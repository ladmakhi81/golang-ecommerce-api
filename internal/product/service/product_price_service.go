package productservice

import (
	productdto "github.com/ladmakhi81/golang-ecommerce-api/internal/product/dto"
	productentity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
)

type IProductPriceService interface {
	AddPriceToProductsPriceList(reqBody productdto.AddPriceToProductsPriceListReqBody, productId uint, assignerId uint) (*productentity.ProductPrice, error)
	RemovePriceItemFromProductList(priceItem uint) error
	IsProductPriceItemExist(priceItem uint) (bool, error)
	FindPriceItemById(priceItemID uint) (*productentity.ProductPrice, error)
	FindPricesByProductId(productId uint) (*[]productentity.ProductPrice, error)
}
