package cartservice

import (
	"net/http"

	cartdto "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/dto"
	cartentity "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/entity"
	cartrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/repository"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	productservice "github.com/ladmakhi81/golang-ecommerce-api/internal/product/service"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
)

type CartService struct {
	cartRepo            cartrepository.ICartRepository
	productService      productservice.IProductService
	productPriceService productservice.IProductPriceService
	userService         userservice.IUserService
}

func NewCartService(
	cartRepo cartrepository.ICartRepository,
	productService productservice.IProductService,
	productPriceService productservice.IProductPriceService,
	userService userservice.IUserService,
) CartService {
	return CartService{
		cartRepo:            cartRepo,
		productService:      productService,
		productPriceService: productPriceService,
		userService:         userService,
	}
}

func (cartService CartService) AddProductToCart(customerID uint, reqBody cartdto.AddProductCartReqBody) (*cartentity.Cart, error) {
	product, productErr := cartService.productService.FindProductById(reqBody.ProductID)
	if productErr != nil {
		return nil, productErr
	}
	priceItem, priceItemErr := cartService.productPriceService.FindPriceItemById(reqBody.PriceItemID)
	if priceItemErr != nil {
		return nil, priceItemErr
	}
	if priceItem == nil {
		return nil, types.NewClientError("price item not found", http.StatusNotFound)
	}
	duplicatedCart, duplicatedCartErr := cartService.cartRepo.FindCustomerCartByProductId(customerID, reqBody.ProductID)
	if duplicatedCartErr != nil {
		return nil, types.NewServerError(
			"error in finding customer cart by product id",
			"CartService.AddProductToCart.FindCustomerCartByProductId",
			duplicatedCartErr,
		)
	}
	if duplicatedCart != nil {
		return nil, types.NewClientError("this cart created before", http.StatusConflict)
	}
	customer, customerErr := cartService.userService.FindBasicUserInfoById(customerID)
	if customerErr != nil {
		return nil, customerErr
	}
	createdCart := cartentity.NewCart(product, customer, priceItem, reqBody.Quantity)
	if createErr := cartService.cartRepo.CreateProductCart(createdCart); createErr != nil {
		return nil, types.NewServerError(
			"error in creating cart",
			"CartService.AddProductToCart.CreateProductCart",
			createErr,
		)
	}
	return createdCart, nil
}
func (cartService CartService) FindCustomerCartByProductId(customerID, productID uint) (*cartentity.Cart, error) {
	cart, cartErr := cartService.cartRepo.FindCustomerCartByProductId(customerID, productID)
	if cartErr != nil {
		return nil, types.NewServerError(
			"error in finding carts of customer",
			"CartService.FindCustomerCartByProductId",
			cartErr,
		)
	}
	return cart, nil
}
func (cartService CartService) UpdateCartQuantityById(customerID, cartId uint, reqBody cartdto.UpdateCartQuantityReqBody) error {
	cart, cartErr := cartService.FindCartById(cartId)
	if cartErr != nil {
		return cartErr
	}
	if cart.Customer.ID != customerID {
		return types.NewClientError(
			"the owner of cart can update",
			http.StatusForbidden,
		)
	}
	cart.Quantity = reqBody.Quantity
	if updateErr := cartService.cartRepo.UpdateCartById(cart); updateErr != nil {
		return types.NewServerError(
			"error in updating cart",
			"CartService.UpdateCartQuantityById.UpdateCartById",
			updateErr,
		)
	}
	return nil
}
func (cartService CartService) FindCartById(cartId uint) (*cartentity.Cart, error) {
	cart, cartErr := cartService.cartRepo.FindCartById(cartId)
	if cartErr != nil {
		return nil, types.NewServerError(
			"error in finding cart by id",
			"CartService.FindCartById",
			cartErr,
		)
	}
	if cart == nil {
		return nil, types.NewClientError(
			"cart not found",
			http.StatusNotFound,
		)
	}
	return cart, nil
}
func (cartService CartService) DeleteCartById(customerID, cartId uint) error {
	cart, cartErr := cartService.FindCartById(cartId)
	if cartErr != nil {
		return cartErr
	}
	if cart.Customer.ID != customerID {
		return types.NewClientError(
			"the owner of cart can delete",
			http.StatusForbidden,
		)
	}
	if deleteErr := cartService.cartRepo.DeleteCartById(cart.ID); deleteErr != nil {
		return types.NewServerError(
			"error in delete cart",
			"CartService.DeleteCartById",
			deleteErr,
		)
	}
	return nil
}
func (cartService CartService) FindCustomerCart(customerId uint) ([]*cartentity.Cart, error) {
	carts, cartsErr := cartService.cartRepo.FindCustomerCart(customerId)
	if cartsErr != nil {
		return nil, types.NewServerError(
			"error in finding carts based on customer id",
			"CartService.FindCustomerCart",
			cartsErr,
		)
	}
	return carts, nil
}
