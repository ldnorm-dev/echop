package echop

import (
	va "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Validator struct {
	validator *va.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: va.New(),
	}
}

func (v *Validator) Validate(i any) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
