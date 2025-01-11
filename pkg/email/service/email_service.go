package pkgemail

import pkgemaildto "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/dto"

type IEmailService interface {
	SendEmail(dto pkgemaildto.SendEmailDto)
}
