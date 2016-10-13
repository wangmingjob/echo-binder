package binder

import (
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"io/ioutil"
)

type protobufBinder struct{}

func (protobufBinder) Bind(obj interface{}, c echo.Context) error {
	buf, err := ioutil.ReadAll(c.Request().Body())
	if err != nil {
		return err
	}

	if err = proto.Unmarshal(buf, obj.(proto.Message)); err != nil {
		return err
	}

	return nil
}
