package binder

import (
	"encoding/xml"
	"github.com/labstack/echo"
)

type xmlBinder struct{}

func (xmlBinder) Bind(obj interface{}, c echo.Context) error {
	decoder := xml.NewDecoder(c.Request().Body())
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	xssFilter(obj)
	return validate(obj)
}
