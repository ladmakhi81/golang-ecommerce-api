package orderevent

import (
	"fmt"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
	orderdto "github.com/ladmakhi81/golang-ecommerce-api/internal/order/dto"
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	orderservice "github.com/ladmakhi81/golang-ecommerce-api/internal/order/service"
)

type OrderEventsSubscriber struct {
	orderService orderservice.IOrderService
}

func NewOrderEventsSubscriber(
	orderService orderservice.IOrderService,
) OrderEventsSubscriber {
	return OrderEventsSubscriber{
		orderService: orderService,
	}
}

func (subscriber OrderEventsSubscriber) SubscribeChangeStatusOfOrderToPayed(event events.Event) {
	orderId := event.Payload.(events.PayedOrderEventBody).OrderId
	reqBody := orderdto.ChangeOrderStatusReqBody{Status: orderentity.OrderStatusPayed}
	err := subscriber.orderService.ChangeOrderStatus(orderId, reqBody)
	if err != nil {
		fmt.Println("SubscribeChangeStatusOfOrderToPayed Error : ", err)
	}
}
