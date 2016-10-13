package binder

import "github.com/labstack/echo"

type Binder interface {
	Bind(interface{}, echo.Context) error
}

type StructValidator interface {
	ValidateStruct(interface{}) error
}

var (
	JSON          = jsonBinder{}
	XML           = xmlBinder{}
	Form          = formBinder{}
	FormPost      = formPostBinder{}
	FormMultipart = formMultipartBinder{}
	ProtoBuf      = protobufBinder{}
)

func NewBinder(c echo.Context) Binder {
	if c.Request().Method() == "GET" {
		return Form
	} else {
		switch c.Request().Header().Get("Content-Type") {
		case echo.MIMEApplicationJSON, echo.MIMEApplicationJSONCharsetUTF8:
			return JSON
		case echo.MIMEApplicationXML, echo.MIMEApplicationXMLCharsetUTF8:
			return XML
		case echo.MIMEApplicationProtobuf:
			return ProtoBuf
		case echo.MIMEApplicationForm:
			return FormPost
		case echo.MIMEMultipartForm:
			return FormMultipart
		default:
			return Form
		}
	}
}
