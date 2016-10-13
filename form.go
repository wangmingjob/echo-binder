package binder

import "github.com/labstack/echo"

type formBinder struct{}
type formPostBinder struct{}
type formMultipartBinder struct{}

func (formBinder) Bind(obj interface{}, c echo.Context) error {
	return nil
}

func (formPostBinder) Bind(obj interface{}, c echo.Context) error {
	return nil
}

func (formMultipartBinder) Bind(obj interface{}, c echo.Context) error {
	return nil
}
