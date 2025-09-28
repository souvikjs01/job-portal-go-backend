package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()
		var msg string

		// Customize messages
		switch tag {
		case "required":
			msg = "is required"
		case "email":
			msg = "must be a valid email"
		case "min":
			msg = "value is too short"
		case "max":
			msg = "value is too long"
		default:
			msg = "is invalid"
		}

		errors[field] = msg
	}

	return errors
}
