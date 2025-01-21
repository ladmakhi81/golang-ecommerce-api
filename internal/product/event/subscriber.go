package productevent

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
	pkgemaildto "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/dto"
	pkgemail "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/service"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
)

type ProductEventsSubscriber struct {
	emailService pkgemail.IEmailService
	translation  translations.ITranslation
}

func NewProductEventsSubscriber(
	emailService pkgemail.IEmailService,
	translation translations.ITranslation,
) ProductEventsSubscriber {
	return ProductEventsSubscriber{
		emailService: emailService,
		translation:  translation,
	}
}

func (subscriber ProductEventsSubscriber) SubscribeProductCreated(event events.Event) {
	eventBody := event.Payload.(events.ProductCreatedEventBody)
	product := eventBody.Product
	go subscriber.emailService.SendEmail(
		pkgemaildto.NewSendEmailDto(
			product.Vendor.Email,
			subscriber.translation.MessageWithArgs(
				"product.vendor_product_create_subject_email",
				map[string]any{"ProductName": product.Name},
			),
			subscriber.translation.MessageWithArgs(
				"product.vendor_product_create_body_email",
				map[string]any{
					"Name": product.Name,
					"ID":   product.ID,
					"Date": product.CreatedAt.Format("2006-01-02 15:04:05"),
				},
			),
		),
	)
}

func (subscriber ProductEventsSubscriber) SubscribeProductVerification(event events.Event) {
	eventBody := event.Payload.(events.ProductVerifiedEventBody)
	product := eventBody.Product
	go subscriber.emailService.SendEmail(
		pkgemaildto.NewSendEmailDto(
			product.Vendor.Email,
			subscriber.translation.Message("product.verify_product_subject_email"),
			subscriber.translation.MessageWithArgs(
				"product.verify_product_body_email",
				map[string]any{
					"Name": product.Name,
					"ID":   product.ID,
					"Date": product.ConfirmedAt.Format("2006-01-02 15:04:05"),
				},
			),
		),
	)
}
