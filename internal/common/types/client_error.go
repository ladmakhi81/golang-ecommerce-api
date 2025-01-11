package types

type ClientError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func (clientError ClientError) Error() string {
	return clientError.Message
}

func NewClientError(message string, statusCode int) ClientError {
	return ClientError{Message: message, StatusCode: statusCode}
}
