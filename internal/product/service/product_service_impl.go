package productservice

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	categoryservice "github.com/ladmakhi81/golang-ecommerce-api/internal/category/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	productdto "github.com/ladmakhi81/golang-ecommerce-api/internal/product/dto"
	productentity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
	productrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/product/repository"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
	pkgemaildto "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/dto"
	pkgemail "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/service"
)

type ProductService struct {
	userService     userservice.IUserService
	categoryService categoryservice.ICategoryService
	productRepo     productrepository.IProductRepository
	emailService    pkgemail.IEmailService
}

func NewProductService(
	userService userservice.IUserService,
	categoryService categoryservice.ICategoryService,
	productRepo productrepository.IProductRepository,
	emailService pkgemail.IEmailService,
) ProductService {
	return ProductService{
		userService:     userService,
		categoryService: categoryService,
		productRepo:     productRepo,
		emailService:    emailService,
	}
}

func (productService ProductService) CreateProduct(reqBody productdto.CreateProductReqBody, vendorID uint) (*productentity.Product, error) {
	vendor, vendorErr := productService.userService.FindBasicUserInfoById(vendorID)
	if vendorErr != nil {
		return nil, vendorErr
	}
	// TODO: move this verification logic into separate middleware
	if !vendor.IsVerified {
		return nil, types.NewClientError(
			"you must verify your account",
			http.StatusForbidden,
		)
	}

	category, categoryErr := productService.categoryService.FindCategoryById(reqBody.CategoryID)
	if categoryErr != nil {
		return nil, categoryErr
	}
	product := productentity.NewProduct(
		reqBody.Name,
		reqBody.Description,
		category,
		vendor,
		reqBody.BasePrice,
		reqBody.Tags,
	)

	createProductErr := productService.productRepo.CreateProduct(product)
	if createProductErr != nil {
		return nil, types.NewServerError(
			"error in creating product",
			"ProductService.CreateProduct",
			createProductErr,
		)
	}
	productService.emailService.SendEmail(
		pkgemaildto.NewSendEmailDto(
			product.Vendor.Email,
			fmt.Sprintf("Product %s Created", product.Name),
			fmt.Sprintf("You Create Product %s With ID %d At %v, Wait Until Admins Verify This Product",
				product.Name, product.ID, product.CreatedAt.Format("2006-01-02 15:04:05")),
		),
	)
	return product, nil
}
func (productService ProductService) ConfirmProductByAdmin(adminId uint, productId uint, fee float32) error {
	product, productErr := productService.FindProductById(productId)
	if productErr != nil {
		return productErr
	}
	admin, adminErr := productService.userService.FindBasicUserInfoById(adminId)
	if adminErr != nil {
		return adminErr
	}
	if fee > product.BasePrice {
		return types.NewClientError("fee must be less than base price of products", http.StatusBadRequest)
	}
	if product.IsConfirmed {
		return types.NewClientError("product verified before", http.StatusBadRequest)
	}
	product.ConfirmedBy = admin
	product.IsConfirmed = true
	product.Fee = fee
	product.ConfirmedAt = time.Now()
	if updateErr := productService.productRepo.UpdateProductById(product); updateErr != nil {
		return types.NewServerError(
			"error in updating product",
			"ProductService.ConfirmProductByAdmin.UpdateProductById",
			updateErr,
		)
	}
	productService.emailService.SendEmail(
		pkgemaildto.NewSendEmailDto(
			product.Vendor.Email,
			"Your Product Verified",
			fmt.Sprintf("Product (%s) With ID %d Verified At %v",
				product.Name,
				product.ID,
				product.ConfirmedAt.Format("2006-01-02 15:04:05"),
			),
		),
	)
	return nil
}
func (productService ProductService) FindProductById(id uint) (*productentity.Product, error) {
	product, productErr := productService.productRepo.FindProductById(id)
	if productErr != nil {
		return nil, types.NewServerError(
			"error in finding product by id",
			"ProductService.FindProductById",
			productErr,
		)
	}
	if product == nil {
		return nil, types.NewClientError(
			"product is not found",
			http.StatusNotFound,
		)
	}
	return product, nil
}
func (productService ProductService) GetProductsPage(page, limit uint) ([]*productentity.Product, error) {
	products, productsErr := productService.productRepo.FindProductsPage(page, limit)
	if productsErr != nil {
		return nil, types.NewServerError(
			"error in returning products",
			"ProductService.GetProductsPage.FindProductsPage",
			productsErr,
		)
	}
	return products, nil
}
func (productService ProductService) DeleteProductById(productId, userId uint) error {
	product, productErr := productService.FindProductById(productId)
	if productErr != nil {
		return productErr
	}
	if product.Vendor.ID != userId {
		return types.NewClientError("only creator of product can delete it", http.StatusForbidden)
	}
	if deleteErr := productService.productRepo.DeleteProductById(product.ID); deleteErr != nil {
		return types.NewServerError(
			"error in deleting product",
			"ProductService.DeleteProductById",
			deleteErr,
		)
	}
	return nil
}
func (productService ProductService) UploadProductImages(productId, ownerId uint, multipartForms *multipart.Form) ([]string, error) {
	product, productErr := productService.FindProductById(productId)
	if productErr != nil {
		return nil, productErr
	}
	if product.Vendor.ID != ownerId {
		return nil, types.NewClientError("only the owner of product can upload images", http.StatusForbidden)
	}

	outputFilenames := []string{}
	for _, fileHeader := range multipartForms.File["images"] {
		inputFile, inputFileErr := fileHeader.Open()
		if inputFileErr != nil {
			return nil, types.NewServerError(
				"error in reading input file",
				"ProductService.UploadProductImages.Open",
				inputFileErr,
			)
		}
		fileExtname := filepath.Ext(fileHeader.Filename)
		filename := fmt.Sprintf("%d-%d%s", rand.Intn(10000000000), time.Now().Unix(), fileExtname)
		outputDestination := path.Join("./uploads/", filename)
		outputFile, outputFileErr := os.Create(outputDestination)
		if outputFileErr != nil {
			return nil, types.NewServerError(
				"error in creating output file",
				"ProductService.UploadProductImages.Create",
				outputFileErr,
			)
		}
		if _, copyErr := io.Copy(outputFile, inputFile); copyErr != nil {
			return nil, types.NewServerError(
				"error in copy the input file into output file",
				"ProductService.UploadProductImages.Copy",
				copyErr,
			)
		}
		outputFilenames = append(outputFilenames, filename)
		inputFile.Close()
	}
	product.Images = outputFilenames
	if updateErr := productService.productRepo.UpdateProductById(product); updateErr != nil {
		return nil, types.NewServerError(
			"error in updating product",
			"ProductService.ProductRepo.UpdateProductById",
			updateErr,
		)
	}
	return outputFilenames, nil
}
