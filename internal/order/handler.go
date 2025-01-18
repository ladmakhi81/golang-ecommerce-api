package order

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	orderdto "github.com/ladmakhi81/golang-ecommerce-api/internal/order/dto"
	orderservice "github.com/ladmakhi81/golang-ecommerce-api/internal/order/service"
)

type OrderHandler struct {
	orderService orderservice.IOrderService
}

func NewOrderHandler(
	orderService orderservice.IOrderService,
) OrderHandler {
	return OrderHandler{
		orderService: orderService,
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
	orderRes, orderErr := orderHandler.orderService.SubmitOrder(customerId, reqBody)
	if orderErr != nil {
		return orderErr
	}
	c.JSON(http.StatusCreated, map[string]any{"data": orderRes})
	return nil
}
