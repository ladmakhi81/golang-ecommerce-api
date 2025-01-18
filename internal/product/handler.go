package product

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	productdto "github.com/ladmakhi81/golang-ecommerce-api/internal/product/dto"
	productservice "github.com/ladmakhi81/golang-ecommerce-api/internal/product/service"
)

type ProductHandler struct {
	productService      productservice.IProductService
	productPriceService productservice.IProductPriceService
	util                utils.Util
}

func NewProductHandler(
	productService productservice.IProductService,
	productPriceService productservice.IProductPriceService,
) ProductHandler {
	return ProductHandler{
		productService:      productService,
		productPriceService: productPriceService,
		util:                utils.NewUtil(),
	}
}

func (productHandler ProductHandler) CreateProduct(c echo.Context) error {
	vendorId := c.Get("AuthClaim").(*types.AuthClaim).ID
	var reqBody productdto.CreateProductReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError(
			"invalid request body",
			http.StatusBadRequest,
		)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	product, productErr := productHandler.productService.CreateProduct(reqBody, vendorId)
	if productErr != nil {
		return productErr
	}
	c.JSON(http.StatusCreated, map[string]any{
		"product":   product,
		"isSuccess": true,
	})
	return nil
}
func (productHandler ProductHandler) ConfirmProductByAdmin(c echo.Context) error {
	adminId := c.Get("AuthClaim").(*types.AuthClaim).ID
	productId, parseProductIdErr := productHandler.util.NumericParamConvertor(c.Param("id"), "invalid product id")
	if parseProductIdErr != nil {
		return parseProductIdErr
	}
	if confirmErr := productHandler.productService.ConfirmProductByAdmin(adminId, productId); confirmErr != nil {
		return confirmErr
	}
	return nil
}
func (productHandler ProductHandler) FindProductDetailById(c echo.Context) error {
	productId, parsedProductIdErr := productHandler.util.NumericParamConvertor(c.Param("id"), "invalid product id")
	if parsedProductIdErr != nil {
		return parsedProductIdErr
	}
	product, productErr := productHandler.productService.FindProductById(productId)
	if productErr != nil {
		return productErr
	}
	c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"data":    product,
	})
	return nil
}
func (productHandler ProductHandler) GetProductsPage(c echo.Context) error {
	pagination := productHandler.util.PaginationExtractor(c)
	products, productsErr := productHandler.productService.GetProductsPage(pagination.Page, pagination.Limit)
	if productsErr != nil {
		return productsErr
	}
	c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"data":    products,
	})
	return nil
}
func (productHandler ProductHandler) DeleteProductById(c echo.Context) error {
	vendorId := c.Get("AuthClaim").(*types.AuthClaim).ID
	productId, parsedProductIdErr := productHandler.util.NumericParamConvertor(c.Param("id"), "invalid product id")
	if parsedProductIdErr != nil {
		return parsedProductIdErr
	}
	deleteErr := productHandler.productService.DeleteProductById(productId, vendorId)
	if deleteErr != nil {
		return deleteErr
	}
	c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"message": "delete successfully ...",
	})
	return nil
}
func (productHandler ProductHandler) AddPriceToProductPriceList(c echo.Context) error {
	var reqBody productdto.AddPriceToProductsPriceListReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError("invalid request body", http.StatusBadRequest)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	vendorId := c.Get("AuthClaim").(*types.AuthClaim).ID
	productId, parsedProductIdErr := productHandler.util.NumericParamConvertor(c.Param("product_id"), "invalid product id")
	if parsedProductIdErr != nil {
		return parsedProductIdErr
	}
	priceItem, priceItemErr := productHandler.productPriceService.AddPriceToProductsPriceList(reqBody, productId, vendorId)
	if priceItemErr != nil {
		return priceItemErr
	}
	c.JSON(http.StatusCreated, map[string]any{
		"success": true,
		"data":    priceItem,
	})
	return nil
}
func (productHandler ProductHandler) DeletePriceItem(c echo.Context) error {
	itemId, parsedItemIdErr := productHandler.util.NumericParamConvertor(
		c.Param("id"),
		"invalid price item id",
	)
	if parsedItemIdErr != nil {
		return parsedItemIdErr
	}
	if removeErr := productHandler.productPriceService.RemovePriceItemFromProductList(itemId); removeErr != nil {
		return removeErr
	}
	return nil
}
func (productHandler ProductHandler) GetPricesOfProduct(c echo.Context) error {
	productId, parseErr := productHandler.util.NumericParamConvertor(c.Param("productId"), "invalid product")
	if parseErr != nil {
		return parseErr
	}
	prices, pricesErr := productHandler.productPriceService.FindPricesByProductId(productId)
	if pricesErr != nil {
		return pricesErr
	}
	c.JSON(http.StatusOK, map[string]any{"prices": prices})
	return nil
}
