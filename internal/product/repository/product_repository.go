package productrepository

import productentity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"

type IProductRepository interface {
	CreateProduct(product *productentity.Product) error
	UpdateProductById(product *productentity.Product) error
	FindProductById(id uint) (*productentity.Product, error)
	FindProductsPage(page, limit uint) ([]*productentity.Product, error)
	GetProductsCount() (uint, error)
	DeleteProductById(id uint) error
}
