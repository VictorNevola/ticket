package customvalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/paemuri/brdoc"
)

type (
	CustomValidator struct {
		Validator *validator.Validate
	}
)

func NewCustomValidator() *CustomValidator {
	v := validator.New()
	v.RegisterValidation("isCNPJORCPF", isCNPJORCPF)

	return &CustomValidator{
		Validator: v,
	}
}

func isCNPJORCPF(fl validator.FieldLevel) bool {
	return brdoc.IsCPF(fl.Field().String()) || brdoc.IsCNPJ(fl.Field().String())
}
