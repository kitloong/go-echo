package bootstrap

import (
	"github.com/kitloong/go-echo/routes"
	"github.com/labstack/echo/v4"
)

// CreateApplication Register and create application
func CreateApplication(e *echo.Echo) {
	routes.Register(e)
}
