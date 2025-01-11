package pkgemaildto

type SendEmailDto struct {
	Receiver string
	Subject  string
	Body     string
}

func NewSendEmailDto(receiver, subject, body string) SendEmailDto {
	return SendEmailDto{
		Receiver: receiver,
		Subject:  subject,
		Body:     body,
	}
}
