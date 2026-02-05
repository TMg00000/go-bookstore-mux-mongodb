package validation

import (
	"github.com/go-playground/validator/v10"
)

func ValidateRequestStruct(req interface{}) error {
	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		return err
	}

	return nil
}
