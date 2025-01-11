package types

type ClientValidationError struct {
	Errors map[string]string `json:"errors,omitempty"`
}

func (validationError ClientValidationError) Error() string {
	return "input validation failed ..."
}

func NewClientValidationError(errors map[string]string) ClientValidationError {
	return ClientValidationError{Errors: errors}
}
