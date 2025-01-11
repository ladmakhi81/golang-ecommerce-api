package types

type ServerError struct {
	Message  string `json:"message"`
	Location string `json:"location"`
	Detail   error  `json:"detail"`
}

func (serverError ServerError) Error() string {
	return serverError.Message + " - " + serverError.Location + " - " + serverError.Detail.Error()
}

func NewServerError(message string, location string, detail error) ServerError {
	return ServerError{
		Message:  message,
		Location: location,
		Detail:   detail,
	}
}
