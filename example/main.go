package main

import (
	"github.com/cnjack/echo-binder"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
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
		return c.String(http.StatusOK, "Hello, "+u.Name)
	})
	e.GET("/", func(c echo.Context) error {
		var u User
		if err := c.Bind(&u); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, "Hello, "+u.Name)
	})
	e.Run(standard.New(":1314"))
}
