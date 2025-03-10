package helpers

import "github.com/go-playground/validator"

func Validate(v interface{}) error {
	validate := validator.New()

	return validate.Struct(v)
}
