package binder

import (
	"reflect"
	"github.com/microcosm-cc/bluemonday"
)

var p = bluemonday.UGCPolicy().AllowElements("font").AllowAttrs("style", "align", "color", "size").Globally()

func xssFilter (ptr interface{}) {
	if kindOfData(ptr) == reflect.Struct {
		typ := reflect.TypeOf(ptr).Elem()
		val := reflect.ValueOf(ptr).Elem()
		for i := 0; i < typ.NumField(); i++ {
			typeField := typ.Field(i)
			structField := val.Field(i)
			if !structField.CanSet() {
				continue
			}
			structFieldKind := structField.Kind()
			xss := typeField.Tag.Get("xss")
			if structFieldKind == reflect.Struct {
				xssFilter(structField.Addr().Interface())
			}
			if xss == "true" && structFieldKind == reflect.String {
				structField.SetString(p.Sanitize(structField.String()))
			}
		}
	}
}