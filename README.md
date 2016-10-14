# echo-binder
echo-binder 一个提供echo中数据binder和validator功能的middleware

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

## Thx
[echo](https://github.com/labstack/echo)
[assert](github.com/stretchr/testify/assert)
[validator.v9](gopkg.in/go-playground/validator.v9)
