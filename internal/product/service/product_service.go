package productservice

import (
	"mime/multipart"

	productdto "github.com/ladmakhi81/golang-ecommerce-api/internal/product/dto"
	productentity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
)

type IProductService interface {
	CreateProduct(reqBody productdto.CreateProductReqBody, vendorID uint) (*productentity.Product, error)
	ConfirmProductByAdmin(adminId uint, productId uint, fee float32) error
	FindProductById(id uint) (*productentity.Product, error)
	GetProductsPage(page, limit uint) ([]*productentity.Product, uint, error)
	DeleteProductById(productId, userId uint) error
	UploadProductImages(productId uint, ownerId uint, multipartForms *multipart.Form) ([]string, error)
}
