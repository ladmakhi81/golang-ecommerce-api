package orderservice

import (
	"net/http"

	cartservice "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	orderdto "github.com/ladmakhi81/golang-ecommerce-api/internal/order/dto"
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	orderrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/order/repository"
	paymentservice "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/service"
	productservice "github.com/ladmakhi81/golang-ecommerce-api/internal/product/service"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
)

type OrderService struct {
	userService    userservice.IUserService
	orderRepo      orderrepository.IOrderRepository
	cartService    cartservice.ICartService
	productService productservice.IProductService
	paymentService paymentservice.IPaymentService
}

func NewOrderService(
	userService userservice.IUserService,
	orderRepo orderrepository.IOrderRepository,
	cartService cartservice.ICartService,
	productService productservice.IProductService,
	paymentService paymentservice.IPaymentService,
) OrderService {
	return OrderService{
		userService:    userService,
		orderRepo:      orderRepo,
		cartService:    cartService,
		productService: productService,
		paymentService: paymentService,
	}
}

func (orderService OrderService) SubmitOrder(customerId uint, reqBody orderdto.CreateOrderReqBody) (*orderdto.CreateOrderResponse, error) {
	customer, customerErr := orderService.userService.FindBasicUserInfoById(customerId)
	if customerErr != nil {
		return nil, customerErr
	}
	carts, cartsErr := orderService.cartService.FindCartsByIds(reqBody.CartIds)
	if cartsErr != nil {
		return nil, cartsErr
	}
	for _, cart := range carts {
		if cart.Customer.ID != customerId {
			return nil, types.NewClientError("only the owner of the cart can buy this cart", http.StatusForbidden)
		}
	}
	finalPrice := orderService.cartService.CalculateFinalPriceOfCarts(carts)
	order := orderentity.NewOrder(customer, finalPrice)
	order.Items = []*orderentity.OrderItem{}
	orderErr := orderService.orderRepo.CreateOrder(order)
	if orderErr != nil {
		return nil, types.NewServerError(
			"error in creating order",
			"OrderService.SubmitOrder.CreateOrder",
			orderErr,
		)
	}
	for _, cart := range carts {
		product, productErr := orderService.productService.FindProductById(cart.Product.ID)
		if productErr != nil {
			return nil, productErr
		}
		orderItem := orderentity.NewOrderItem(
			cart.Product,
			cart.PriceItem,
			product.Vendor,
			cart.Customer,
			order,
			cart.Quantity,
		)
		if orderItemErr := orderService.orderRepo.CreateOrderItem(orderItem); orderItemErr != nil {
			return nil, types.NewServerError(
				"error in creating order item",
				"OrderService.SubmitOrder.OrderRepo.CreateOrderItem",
				orderItemErr,
			)
		}
		order.Items = append(order.Items, orderItem)
	}
	payment, paymentErr := orderService.paymentService.CreatePayment(order)
	if paymentErr != nil {
		return nil, paymentErr
	}
	payLink := orderService.paymentService.GetPayLink(payment)
	res := orderdto.NewCreateOrderResponse(
		order.ID,
		payLink,
	)
	return res, nil
}
func (orderService OrderService) FindOrderItemsByOrderId(orderId uint) ([]*orderentity.OrderItem, error) {
	orderItems, orderItemsErr := orderService.orderRepo.FindOrderItemsByOrderId(orderId)
	if orderItemsErr != nil {
		return nil, types.NewServerError("error in finding order items", "OrderService.FindOrderItemsByOrderId", orderItemsErr)
	}
	return orderItems, nil
}
