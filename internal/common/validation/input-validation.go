package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
)

type InputValidator struct {
	validator *validator.Validate
}

func NewInputValidator() *InputValidator {
	return &InputValidator{
		validator: validator.New(),
	}
}

func (inputValidator *InputValidator) Validate(input any) error {
	errors := make(map[string]string)
	if err := inputValidator.validator.Struct(input); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errors[fieldErr.Field()] = fieldErr.Error()
		}
	}
	if len(errors) > 0 {
		return types.NewClientValidationError(errors)
	}
	return nil
}
