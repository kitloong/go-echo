package bootstrap

import (
	. "github.com/kitloong/go-echo/app/providers"

	"github.com/kitloong/go-echo/app/helpers/app"
	"github.com/labstack/echo/v4"
)

// CreateApplication and load services
func CreateApplication() *echo.Echo {
	e := app.Make()

	// Place boot application on top
	new(AppServiceProvider).Boot(e)

	new(RouteServiceProvider).Boot(e)

	return e
}
