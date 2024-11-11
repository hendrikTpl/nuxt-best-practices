package validation

import (
	"fmt"
	"taskmanager/models"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateTask(task models.Task) map[string]string {
	err := validate.Struct(task)

	if err == nil {
		return nil
	}

	errors := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		var errorMessage string
		switch err.Tag() {
		case "required":
			errorMessage = fmt.Sprintf("%s is required", err.Field())
		case "min":
			errorMessage = fmt.Sprintf("%s must be atleast %s characters long", err.Field(), err.Param())
		case "max":
			errorMessage = fmt.Sprintf("%s cannot be longer than %s characters", err.Field(), err.Param())
		default:
			errorMessage = fmt.Sprintf("%s is invalid", err.Field())
		}
		errors[err.Field()] = errorMessage
	}

	return errors
}
