package productrepository

import productentity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"

type IProductPriceRepository interface {
	CreateProductPrice(productPrice *productentity.ProductPrice) error
	DeleteProductPriceById(id uint) error
	IsProductPriceItemExist(id uint) (bool, error)
}
