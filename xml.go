package binder

import "github.com/labstack/echo"

type xmlBinder struct{}

func (xmlBinder) Bind(obj interface{}, c echo.Context) error {
	return nil
}
