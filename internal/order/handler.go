package order

import (
	"net/http"

	"github.com/labstack/echo/v4"
	responsehandling "github.com/ladmakhi81/golang-ecommerce-api/internal/common/response_handling"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	orderdto "github.com/ladmakhi81/golang-ecommerce-api/internal/order/dto"
	orderservice "github.com/ladmakhi81/golang-ecommerce-api/internal/order/service"
)

type OrderHandler struct {
	orderService orderservice.IOrderService
	util         utils.Util
}

func NewOrderHandler(
	orderService orderservice.IOrderService,
) OrderHandler {
	return OrderHandler{
		orderService: orderService,
		util:         utils.NewUtil(),
	}
}

func (orderHandler OrderHandler) CreateOrder(c echo.Context) error {
	var reqBody orderdto.CreateOrderReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError("invalid request body", http.StatusBadRequest)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	customerId := c.Get("AuthClaim").(*types.AuthClaim).ID
	createdOrder, orderErr := orderHandler.orderService.SubmitOrder(customerId, reqBody)
	if orderErr != nil {
		return orderErr
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusCreated,
		createdOrder,
	)
	return nil
}
func (orderHandler OrderHandler) UpdateOrderStatus(c echo.Context) error {
	orderId, parsedOrderIdErr := orderHandler.util.NumericParamConvertor(c.Param("orderId"), "invalid order id")
	if parsedOrderIdErr != nil {
		return parsedOrderIdErr
	}
	var reqBody orderdto.ChangeOrderStatusReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError("invalid request body", http.StatusBadRequest)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	if !reqBody.Status.IsValid() {
		return types.NewClientValidationError(map[string]string{"Status": "Status must be between Pending | Payed | Preparation | Delivery | Done"})
	}
	if err := orderHandler.orderService.ChangeOrderStatus(orderId, reqBody); err != nil {
		return err
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		nil,
	)
	return nil
}
func (orderHandler OrderHandler) FindOrdersPage(c echo.Context) error {
	pagination := orderHandler.util.PaginationExtractor(c)
	orders, ordersCount, ordersErr := orderHandler.orderService.FindOrdersPage(pagination.Page, pagination.Limit)
	if ordersErr != nil {
		return ordersErr
	}
	paginatedResponse := types.NewPaginationResponse(
		ordersCount,
		pagination,
		orders,
	)
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		paginatedResponse,
	)
	return nil
}
