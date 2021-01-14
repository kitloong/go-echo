package controllers

import (
	"github.com/kitloong/go-echo/app/helpers/session"
	"github.com/kitloong/go-echo/app/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

// UserController define user handlers
type UserController struct {
	Controller
}

var u *UserController

func init() {
	u = &UserController{}
}

// User UserController struct
func User() *UserController {
	return u
}

// Index list records
func (u *UserController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!"+session.Get("Foo", c).(string))
}

// Show get single record
func (u *UserController) Show(c echo.Context) (err error) {
	session.Put("Foo", "Bar", c)

	id, _ := strconv.Atoi(c.Param("id"))
	user := new(models.User)
	user.ID = id
	user.Name = "Name"
	user.Email = "test@email.com"
	user.Now = time.Now()

	return c.JSON(http.StatusOK, user)
}

// Store create single record
func (u *UserController) Store(c echo.Context) (err error) {
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		return
	}
	return c.JSON(http.StatusCreated, user)
}

// Update single record
func (u *UserController) Update(c echo.Context) (err error) {
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		return
	}
	return c.JSON(http.StatusOK, user)
}

// Delete single record
func (u *UserController) Delete(c echo.Context) (err error) {
	//id, _ := strconv.Atoi(c.Param("id"))
	return c.NoContent(http.StatusNoContent)
}
