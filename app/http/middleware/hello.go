package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type Hello struct {
}

// Handle session start
func (m *Hello) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Hello")
		return next(c)
	}
}
