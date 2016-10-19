# echo-binder
echo-binder 一个提供echo中数据binder和validator功能的middleware

##TODO
 - 添加gin的标签  
 - 完善注入方式

## Update

- 20161018 使用[bluemonday](github.com/microcosm-cc/bluemonday)添加xss过滤,使用方式详见test TestXssBinder_Bind


## Quick Start

### Installation
```
$ go get -u github.com/cnjack/echo-binder
```
### Hello, World!
```
package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/cnjack/echo-binder"
	"github.com/labstack/echo/engine/standard"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" binding:"required"`
	Age   int    `json:"age" xml:"age" form:"age" binding:"gte=0,lte=130"`
	Email string `json:"email" xml:"email" form:"email" binding:"required,email"`
}

func main() {
	e := echo.New()
	e.Use(binder.BindBinder(e))
	e.POST("/", func(c echo.Context) error {
		var u User
		if err := c.Bind(&u); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, "Hello, " + u.Name)
	})
	e.Run(standard.New(":1314"))
}
```

## Thx.
[echo](https://github.com/labstack/echo) Fast and unfancy HTTP server framework for Go (Golang)  
[assert](https://github.com/stretchr/testify) A sacred extension to the standard go testing package  
[validator.v9](https://gopkg.in/go-playground/validator.v9) Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving  
[bluemonday](https://github.com/microcosm-cc/bluemonday) a fast golang HTML sanitizer (inspired by the OWASP Java HTML Sanitizer) to scrub user generated content of XSS  
