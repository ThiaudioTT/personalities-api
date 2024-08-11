package models

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func ValidateModel(model interface{}) error {
	if validate == nil {
		validate = validator.New()
	}

	err := validate.Struct(model)
	if err != nil {
		return err
	}

	return nil
}
