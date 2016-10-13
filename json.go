package binder

import "github.com/labstack/echo"

type jsonBinder struct{}

func (jsonBinder) Bind(obj interface{}, c echo.Context) error {
	return nil
}
