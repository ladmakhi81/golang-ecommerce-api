package order

import (
	"net/http"

	"github.com/labstack/echo/v4"
	responsehandling "github.com/ladmakhi81/golang-ecommerce-api/internal/common/response_handling"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	orderdto "github.com/ladmakhi81/golang-ecommerce-api/internal/order/dto"
	orderservice "github.com/ladmakhi81/golang-ecommerce-api/internal/order/service"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
)

type OrderHandler struct {
	orderService orderservice.IOrderService
	util         utils.Util
	translation  translations.ITranslation
}

func NewOrderHandler(
	orderService orderservice.IOrderService,
	translation translations.ITranslation,
) OrderHandler {
	return OrderHandler{
		orderService: orderService,
		util:         utils.NewUtil(),
		translation:  translation,
	}
}

func (orderHandler OrderHandler) CreateOrder(c echo.Context) error {
	var reqBody orderdto.CreateOrderReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError(
			orderHandler.translation.Message("errors.invalid_request_body"),
			http.StatusBadRequest,
		)
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
	orderId, parsedOrderIdErr := orderHandler.util.NumericParamConvertor(
		c.Param("orderId"),
		orderHandler.translation.Message("order.invalid_id"),
	)
	if parsedOrderIdErr != nil {
		return parsedOrderIdErr
	}
	var reqBody orderdto.ChangeOrderStatusReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError(
			orderHandler.translation.Message("errors.invalid_request_body"),
			http.StatusBadRequest,
		)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	if !reqBody.Status.IsValid() {
		return types.NewClientValidationError(map[string]string{
			"Status": orderHandler.translation.Message("order.invalid_order_status_err"),
		})
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
