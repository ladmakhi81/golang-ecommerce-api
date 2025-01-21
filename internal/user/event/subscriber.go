package userevent

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
	pkgemaildto "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/dto"
	pkgemail "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/service"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
)

type UserEventsSubscriber struct {
	emailService pkgemail.IEmailService
	translation  translations.ITranslation
}

func NewUserEventsSubscriber(
	emailService pkgemail.IEmailService,
	translation translations.ITranslation,
) UserEventsSubscriber {
	return UserEventsSubscriber{
		emailService: emailService,
		translation:  translation,
	}
}

func (subscriber UserEventsSubscriber) SubscribeUserRegistered(event events.Event) {
	eventBody := event.Payload.(events.UserRegisteredEventBody)
	go subscriber.emailService.SendEmail(
		pkgemaildto.NewSendEmailDto(
			eventBody.Email,
			subscriber.translation.Message("auth.signup_subject_email"),
			subscriber.translation.Message("auth.signup_body_email"),
		),
	)
}

func (subscriber UserEventsSubscriber) SubscribeCompleteProfile(event events.Event) {
	eventBody := event.Payload.(events.UserCompleteProfileEventBody)
	go subscriber.emailService.SendEmail(
		pkgemaildto.NewSendEmailDto(
			eventBody.Email,
			subscriber.translation.Message("user.complete_profile_subject_email"),
			subscriber.translation.Message("user.complete_profile_body_email"),
		),
	)
}

func (subscriber UserEventsSubscriber) SubscribeUserVerification(event events.Event) {
	eventBody := event.Payload.(events.UserVerificationEventBody)
	go subscriber.emailService.SendEmail(
		pkgemaildto.NewSendEmailDto(
			eventBody.AdminEmail,
			subscriber.translation.Message("user.admin_verify_account_subject_email"),
			subscriber.translation.MessageWithArgs("user.admin_verify_account_body_email", map[string]any{
				"Name": eventBody.VendorFullName,
				"Date": eventBody.Date.Format("2006-01-02 15:04:05"),
			}),
		),
	)
	go subscriber.emailService.SendEmail(
		pkgemaildto.NewSendEmailDto(
			eventBody.VendorEmail,
			subscriber.translation.Message("user.user_verify_account_subject_email"),
			subscriber.translation.Message("user.user_verify_account_body_email"),
		),
	)
}
