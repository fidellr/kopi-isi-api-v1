package validator

import (
	"gopkg.in/go-playground/validator.v9"
)

func Validate(data interface{}) error {
	vld := validator.New()
	return vld.Struct(data)
}
