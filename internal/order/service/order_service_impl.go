package orderservice

import (
	"fmt"
	"net/http"
	"time"

	cartservice "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	orderdto "github.com/ladmakhi81/golang-ecommerce-api/internal/order/dto"
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	orderrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/order/repository"
	paymentservice "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/service"
	productservice "github.com/ladmakhi81/golang-ecommerce-api/internal/product/service"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
	pkgemaildto "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/dto"
	pkgemail "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/service"
)

type OrderService struct {
	userService        userservice.IUserService
	orderRepo          orderrepository.IOrderRepository
	cartService        cartservice.ICartService
	productService     productservice.IProductService
	paymentService     paymentservice.IPaymentService
	emailService       pkgemail.IEmailService
	userAddressService userservice.IUserAddressService
}

func NewOrderService(
	userService userservice.IUserService,
	orderRepo orderrepository.IOrderRepository,
	cartService cartservice.ICartService,
	productService productservice.IProductService,
	paymentService paymentservice.IPaymentService,
	emailService pkgemail.IEmailService,
	userAddressService userservice.IUserAddressService,
) OrderService {
	return OrderService{
		userService:        userService,
		orderRepo:          orderRepo,
		cartService:        cartService,
		productService:     productService,
		paymentService:     paymentService,
		emailService:       emailService,
		userAddressService: userAddressService,
	}
}

func (orderService OrderService) SubmitOrder(customerId uint, reqBody orderdto.CreateOrderReqBody) (*orderdto.CreateOrderResponse, error) {
	customer, customerErr := orderService.userService.FindBasicUserInfoById(customerId)
	if customerErr != nil {
		return nil, customerErr
	}
	// user don't have any address
	if customer.ActiveAddress.ID == 0 && reqBody.AddressId == 0 {
		return nil, types.NewClientError("user don't have any address", http.StatusBadRequest)
	}
	carts, cartsErr := orderService.cartService.FindCartsByIds(reqBody.CartIds)
	if cartsErr != nil {
		return nil, cartsErr
	}
	if len(carts) != len(reqBody.CartIds) {
		return nil, types.NewClientError("carts is not found", http.StatusNotFound)
	}
	for _, cart := range carts {
		if cart.Customer.ID != customerId {
			return nil, types.NewClientError("only the owner of the cart can buy this cart", http.StatusForbidden)
		}
	}
	finalPrice := orderService.cartService.CalculateFinalPriceOfCarts(carts)
	order := orderentity.NewOrder(customer, finalPrice)
	if reqBody.AddressId == 0 {
		order.Address = customer.ActiveAddress
	} else {
		address, addressErr := orderService.userAddressService.FindAddressById(reqBody.AddressId)
		if addressErr != nil {
			return nil, addressErr
		}
		order.Address = address
	}
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
	if deleteCartsErr := orderService.cartService.DeleteCartsByIds(reqBody.CartIds); deleteCartsErr != nil {
		return nil, deleteCartsErr
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
func (orderService OrderService) FindOrderById(id uint) (*orderentity.Order, error) {
	order, orderErr := orderService.orderRepo.FindOrderById(id)
	if orderErr != nil {
		return nil, types.NewServerError(
			"error in finding order by id",
			"OrderService.FindOrderById",
			orderErr,
		)
	}
	if order == nil {
		return nil, types.NewClientError("order not found", http.StatusNotFound)
	}
	return order, nil
}
func (orderService OrderService) ChangeOrderStatus(orderId uint, reqBody orderdto.ChangeOrderStatusReqBody) error {
	order, orderErr := orderService.FindOrderById(orderId)
	if orderErr != nil {
		return orderErr
	}
	order.Status = reqBody.Status
	order.StatusChangedAt = time.Now()
	if updateErr := orderService.orderRepo.ChanegOrderStatus(order); updateErr != nil {
		return types.NewServerError(
			"error in updating order status",
			"orderService.ChangeOrderStatus",
			updateErr,
		)
	}
	orderService.emailService.SendEmail(
		pkgemaildto.NewSendEmailDto(
			order.Customer.Email,
			"Order Updated",
			fmt.Sprintf("Your Order With ID %d Updated To %s In %s",
				order.ID,
				order.Status,
				order.StatusChangedAt.Format("2006-01-02 15:04:05"),
			),
		),
	)
	return nil
}
