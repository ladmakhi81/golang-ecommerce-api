package productservice

import (
	productdto "github.com/ladmakhi81/golang-ecommerce-api/internal/product/dto"
	productentity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
)

type IProductService interface {
	CreateProduct(reqBody productdto.CreateProductReqBody, vendorID uint) (*productentity.Product, error)
	ConfirmProductByAdmin(adminId uint, productId uint) error
	FindProductById(id uint) (*productentity.Product, error)
	GetProductsPage(page, limit uint) ([]*productentity.Product, error)
	DeleteProductById(productId, userId uint) error
}
