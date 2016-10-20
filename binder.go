package binder

import (
	"github.com/labstack/echo"
	"strings"
)

type Binder interface {
	Bind(interface{}, echo.Context) error
}

type StructValidator interface {
	ValidateStruct(interface{}) error
}

var Validator StructValidator = &defaultValidator{}

var (
	JSON     = jsonBinder{}
	XML      = xmlBinder{}
	Form     = formBinder{}
	FormPost = formPostBinder{}
	ProtoBuf = protobufBinder{}
)

func BindBinder(e *echo.Echo) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			b := NewBinder(c)
			e.SetBinder(b)
			return next(c)
		}
	}
}

func NewBinder(c echo.Context) Binder {

	if c.Request().Method() == echo.GET {
		return Form
	} else {
		ctype := c.Request().Header().Get(echo.HeaderContentType)
		switch {
		case strings.HasPrefix(ctype, echo.MIMEApplicationJSON):
			return JSON
		case strings.HasPrefix(ctype, echo.MIMEApplicationXML):
			return XML
		case strings.HasPrefix(ctype, echo.MIMEApplicationProtobuf):
			return ProtoBuf
		case strings.HasPrefix(ctype, echo.MIMEApplicationForm), strings.HasPrefix(ctype, echo.MIMEMultipartForm):
			return FormPost
		default:
			return Form
		}
	}

}

func validate(obj interface{}) error {
	if Validator == nil {
		return nil
	}
	return Validator.ValidateStruct(obj)
}
