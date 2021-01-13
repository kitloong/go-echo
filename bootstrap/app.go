package bootstrap

import (
	"github.com/kitloong/go-echo/app/providers"
	"github.com/labstack/echo/v4"
)

// CreateApplication Register and create application
func CreateApplication() *echo.Echo {
	e := echo.New()

	// Boot config provider at top
	new(providers.ConfigServiceProvider).Boot(e)

	new(providers.RouteServiceProvider).Boot(e)

	return e
}
