package validation

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors []FieldError

func (e ValidationErrors) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func extractErrors(err error) error {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}
	var errs ValidationErrors
	for _, e := range err.(validator.ValidationErrors) {
		errs = append(errs, FieldError{
			Field:   e.Field(),
			Message: e.Translate(trans),
		})
	}
	return errs
}
