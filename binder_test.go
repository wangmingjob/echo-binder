package binder_test

import (
	"github.com/cnjack/echo-binder"
	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" binding:"required"`
	Age   int    `json:"age" xml:"age" form:"age" binding:"gte=0,lte=130"`
	Email string `json:"email" xml:"email" form:"email" binding:"required,email"`
}

var (
	json = `{"name": "jack","age": 25,"email": "h_7357@qq.com"}`
)

func TestFormBinder_Bind(t *testing.T) {
	e := echo.New()
	rec := test.NewResponseRecorder()
	req := test.NewRequest("POST", "/", strings.NewReader(json))
	c := e.NewContext(req, rec)
	req.Header().Set(echo.HeaderContentType, "application/json")
	b := binder.NewBinder(c)
	var user User
	err := b.Bind(&user, c)
	if assert.NoError(t, err) {
		assert.Equal(t, "jack", user.Name)
		assert.Equal(t, 25, user.Age)
		assert.Equal(t, "h_7357@qq.com", user.Email)
	}
}
