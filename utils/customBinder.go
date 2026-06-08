package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomBinder struct {
	Validator *validator.Validate
}

func (cb *CustomBinder) Bind(i any, c echo.Context) error {
	db := &echo.DefaultBinder{}
	println("In CustomBinder", i)
	if err := db.Bind(i, c); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}
	if err := cb.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}
	return nil
}
