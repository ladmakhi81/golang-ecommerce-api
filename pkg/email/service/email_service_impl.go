package pkgemail

import (
	"fmt"
	"net/smtp"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	pkgemaildto "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/dto"
)

type EmailService struct {
	config config.MainConfig
}

func NewEmailService(config config.MainConfig) EmailService {
	return EmailService{config}
}

func (emailService EmailService) SendEmail(dto pkgemaildto.SendEmailDto) {
	mailUser := emailService.config.MailUser
	mailPassword := emailService.config.MailPassword
	mailHost := emailService.config.MailHost
	mailPort := emailService.config.MailPort

	receivers := []string{dto.Receiver}
	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", receivers[0], dto.Subject, dto.Body))

	auth := smtp.PlainAuth(
		"",
		mailUser,
		mailPassword,
		mailHost,
	)
	addr := fmt.Sprintf("%s:%d", mailHost, mailPort)
	err := smtp.SendMail(addr, auth, mailUser, receivers, []byte(message))
	if err != nil {
		fmt.Println("Email not send for this address", dto.Receiver, err)
	}
}
