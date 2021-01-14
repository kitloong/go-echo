package bootstrap

import (
	"github.com/kitloong/go-echo/app/helpers/app"
	config2 "github.com/kitloong/go-echo/app/helpers/config"
	"github.com/kitloong/go-echo/app/providers"
	"github.com/labstack/echo/v4"
	"time"
)

// CreateApplication and load services
func CreateApplication() *echo.Echo {
	e := app.Make()

	// Boot config provider at top
	new(providers.ConfigServiceProvider).Boot(e)

	loc, _ := time.LoadLocation(config2.Get("app.timezone").(string))
	time.Local = loc

	new(providers.RouteServiceProvider).Boot(e)

	return e
}
