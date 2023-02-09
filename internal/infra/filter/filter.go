package filter

import "github.com/go-playground/validator/v10"

type (
	Validator struct {
		v *validator.Validate
	}
)

func NewValidator() *Validator {
	return &Validator{
		v: validator.New(),
	}
}
