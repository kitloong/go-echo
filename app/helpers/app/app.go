package app

import (
	"github.com/labstack/echo/v4"
	"sync"
)

var (
	once sync.Once
	e    *echo.Echo
)

// Get echo instance
func Get() *echo.Echo {
	return e
}

// Make a new echo instance
func Make() *echo.Echo {
	once.Do(func() {
		e = echo.New()
	})
	return e
}
