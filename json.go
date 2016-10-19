package binder

import (
	"encoding/json"
	"github.com/labstack/echo"
)

type jsonBinder struct{}

func (jsonBinder) Bind(obj interface{}, c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body())

	if err := decoder.Decode(obj); err != nil {
		return err
	}
	xssFilter(obj)
	return validate(obj)
}
