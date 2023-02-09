package filter

import (
	"github.com/go-playground/validator/v10"
)

func IsEmail(fl validator.FieldLevel) bool {
	// ...
	return true
}

func IsNumberString(fl validator.FieldLevel) bool {
	// ...
	return true
}

func IsPhoneNumber(fl validator.FieldLevel) bool {
	// ...
	return true
}
