package orderevent

import (
	"fmt"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
	orderdto "github.com/ladmakhi81/golang-ecommerce-api/internal/order/dto"
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	orderservice "github.com/ladmakhi81/golang-ecommerce-api/internal/order/service"
	pkgemaildto "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/dto"
	pkgemail "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/service"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
)

type OrderEventsSubscriber struct {
	orderService orderservice.IOrderService
	emailService pkgemail.IEmailService
	translation  translations.ITranslation
}

func NewOrderEventsSubscriber(
	orderService orderservice.IOrderService,
	emailService pkgemail.IEmailService,
	translation translations.ITranslation,
) OrderEventsSubscriber {
	return OrderEventsSubscriber{
		orderService: orderService,
		emailService: emailService,
		translation:  translation,
	}
}

func (subscriber OrderEventsSubscriber) SubscribeChangeOrderStatus(event events.Event) {
	eventBody := event.Payload.(events.ChangeOrderStatusEventBody)
	customer := eventBody.Customer
	order := eventBody.Order

	// UPDATE ORDER STATUS TO PAYER
	if eventBody.Order.Status == orderentity.OrderStatusPayed {
		reqBody := orderdto.ChangeOrderStatusReqBody{Status: orderentity.OrderStatusPayed}
		err := subscriber.orderService.ChangeOrderStatus(order.ID, reqBody)
		if err != nil {
			fmt.Println("SubscribeChangeStatusOfOrderToPayed Error : ", err)
		}
		return
	}
	// NOTIFY USER ORDER STATUS CHANGED
	go subscriber.emailService.SendEmail(
		pkgemaildto.NewSendEmailDto(
			customer.Email,
			subscriber.translation.Message("order.order_status_update_subject_email"),
			subscriber.translation.MessageWithArgs(
				"order.order_status_update_body_email",
				map[string]any{
					"ID":     order.ID,
					"Status": order.Status,
					"Date":   order.StatusChangedAt.Format("2006-01-02 15:04:05"),
				},
			),
		),
	)
}
