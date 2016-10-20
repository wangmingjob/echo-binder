package binder_test

import (
	"github.com/cnjack/echo-binder"
	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"fmt"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" binding:"required"`
	Age   int    `json:"age" xml:"age" form:"age" binding:"gte=0,lte=130"`
	Email string `json:"email" xml:"email" form:"email" binding:"required,email"`
}

type Xss struct {
	Data string `json:"data" xss:"true"`
	Image string `json:"image" xss:"true"`
}

var (
	json = `{"name": "jack","age": 25,"email": "h_7357@qq.com"}`
	xml = `<xml><name>jack</name><age>25</age><email>h_7357@qq.com</email></xml>`
	form = `name=jack&age=25&email=h_7357@qq.com`
	xss = `{"data":"<a onblur='alert(secret)' href='http://www.google.com'>Google</a>", "image":"<img src='https://ssl.gstatic.com/accounts/ui/logo_2x.png'/>"}`
)

func TestFormBinder_Bind(t *testing.T) {
	e := echo.New()
	rec := test.NewResponseRecorder()
	req := test.NewRequest("GET", "/?" + form, strings.NewReader(""))
	c := e.NewContext(req, rec)
	b := binder.NewBinder(c)
	var user User
	err := b.Bind(&user, c)
	if assert.NoError(t, err) {
		assert.Equal(t, "jack", user.Name)
		assert.Equal(t, 25, user.Age)
		assert.Equal(t, "h_7357@qq.com", user.Email)
	}
}

func TestFormPostBinder_Bind(t *testing.T) {
	e := echo.New()
	rec := test.NewResponseRecorder()
	req := test.NewRequest("POST", "/", strings.NewReader(form))
	c := e.NewContext(req, rec)
	req.Header().Set(echo.HeaderContentType, "application/x-www-form-urlencoded")
	b := binder.NewBinder(c)
	var user User
	err := b.Bind(&user, c)
	if assert.NoError(t, err) {
		assert.Equal(t, "jack", user.Name)
		assert.Equal(t, 25, user.Age)
		assert.Equal(t, "h_7357@qq.com", user.Email)
	}
}

func TestXmlBinder_Bind(t *testing.T) {
	e := echo.New()
	rec := test.NewResponseRecorder()
	req := test.NewRequest("POST", "/", strings.NewReader(xml))
	c := e.NewContext(req, rec)
	req.Header().Set(echo.HeaderContentType, "application/xml")
	b := binder.NewBinder(c)
	var user User
	err := b.Bind(&user, c)
	if assert.NoError(t, err) {
		assert.Equal(t, "jack", user.Name)
		assert.Equal(t, 25, user.Age)
		assert.Equal(t, "h_7357@qq.com", user.Email)
	}
}

func TestJsonBinder_Bind(t *testing.T) {
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

func TestXssBinder_Bind(t *testing.T) {
	e := echo.New()
	rec := test.NewResponseRecorder()
	req := test.NewRequest("POST", "/", strings.NewReader(xss))
	c := e.NewContext(req, rec)
	req.Header().Set(echo.HeaderContentType, "application/json")
	b := binder.NewBinder(c)
	var x Xss
	err := b.Bind(&x, c)
	fmt.Println(x.Image)
	if assert.NoError(t, err) {
		assert.Equal(t, "<a href=\"http://www.google.com\" rel=\"nofollow\">Google</a>", x.Data)
	}
}