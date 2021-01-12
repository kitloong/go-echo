package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// UserController Handlers for routes
type UserController struct {
	Controller
	test string
}

var u *UserController

func init() {
	u = &UserController{}
}

// User UserController struct
func User() *UserController {
	return u
}

// Index Listing page
func (u *UserController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!"+u.test)
}

// Show Get single record
func (u *UserController) Show(c echo.Context) error {
	id := c.Param("id")
	u.test = id
	return c.JSON(http.StatusOK, "Users: "+id)
}
