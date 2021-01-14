package providers

import (
	"github.com/kitloong/go-echo/app/helpers/config"
	"github.com/labstack/echo/v4"
	"time"
)

// AppServiceProvider load configuration and setup application setting
type AppServiceProvider struct {
}

// Boot application
func (p *AppServiceProvider) Boot(e *echo.Echo) {
	// First of all, load config
	config.Load()

	// Set default timezone
	loc, _ := time.LoadLocation(config.Get("app.timezone").(string))
	time.Local = loc
}
