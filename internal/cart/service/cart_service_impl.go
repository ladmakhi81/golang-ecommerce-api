package cartservice

import (
	"net/http"

	cartdto "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/dto"
	cartentity "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/entity"
	cartrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/repository"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	productservice "github.com/ladmakhi81/golang-ecommerce-api/internal/product/service"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
)

type CartService struct {
	cartRepo            cartrepository.ICartRepository
	productService      productservice.IProductService
	productPriceService productservice.IProductPriceService
	userService         userservice.IUserService
	translation         translations.ITranslation
}

func NewCartService(
	cartRepo cartrepository.ICartRepository,
	productService productservice.IProductService,
	productPriceService productservice.IProductPriceService,
	userService userservice.IUserService,
	translation translations.ITranslation,
) CartService {
	return CartService{
		cartRepo:            cartRepo,
		productService:      productService,
		productPriceService: productPriceService,
		userService:         userService,
		translation:         translation,
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
		return nil, types.NewClientError(
			cartService.translation.Message("product.price_item_not_found_id"),
			http.StatusNotFound,
		)
	}
	duplicatedCart, duplicatedCartErr := cartService.FindCustomerCartByProductIdAndPriceId(customerID, reqBody.ProductID, reqBody.PriceItemID)
	if duplicatedCartErr != nil {
		return nil, duplicatedCartErr
	}
	if duplicatedCart != nil {
		return nil, types.NewClientError(
			cartService.translation.Message("cart.duplicate_cart"),
			http.StatusConflict,
		)
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
func (cartService CartService) FindCustomerCartByProductIdAndPriceId(customerID, productID uint, priceID uint) (*cartentity.Cart, error) {
	cart, cartErr := cartService.cartRepo.FindCustomerCartByProductIdAndPriceId(customerID, productID, priceID)
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
			cartService.translation.Message("cart.owner_cart_update_err"),
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
			cartService.translation.Message("cart.not_found_id"),
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
			cartService.translation.Message("cart.owner_cart_delete_err"),
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
func (cartService CartService) FindCartsByIds(ids []uint) ([]*cartentity.Cart, error) {
	carts, cartsErr := cartService.cartRepo.FindCartsByIds(ids)
	if cartsErr != nil {
		return nil, types.NewServerError(
			"error in finding carts by ids",
			"CartService.FindCartsByIds",
			cartsErr,
		)
	}
	return carts, nil
}
func (cartService CartService) CalculateFinalPriceOfCarts(carts []*cartentity.Cart) float32 {
	finalPrice := float32(0)
	for _, cart := range carts {
		finalPrice += cart.Product.BasePrice + cart.PriceItem.ExtraPrice
	}
	return finalPrice
}
func (cartService CartService) DeleteCartsByIds(ids []uint) error {
	deleteErr := cartService.cartRepo.DeleteCartsByIds(ids)
	if deleteErr != nil {
		return types.NewServerError(
			"error in deleting carts by ids",
			"CartService.DeleteCartsByIds",
			deleteErr,
		)
	}
	return nil
}
