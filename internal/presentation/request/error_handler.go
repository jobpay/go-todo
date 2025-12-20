package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ParseValidationErrors(err error) []ValidationError {
	var errors []ValidationError

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			errors = append(errors, ValidationError{
				Field:   fieldError.Field(),
				Message: getErrorMessage(fieldError),
			})
		}
	}

	return errors
}

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%sは必須です", fe.Field())
	case "min":
		return fmt.Sprintf("%sは%s文字以上で入力してください", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%sは%s文字以内で入力してください", fe.Field(), fe.Param())
	case "email":
		return fmt.Sprintf("%sの形式が正しくありません", fe.Field())
	default:
		return fmt.Sprintf("%sの入力値が正しくありません", fe.Field())
	}
}
