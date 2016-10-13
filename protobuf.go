package binder

import "github.com/labstack/echo"

type protobufBinder struct{}

func (protobufBinder) Bind(obj interface{}, c echo.Context) error {
	return nil
}