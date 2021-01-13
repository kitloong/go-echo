package providers

import (
	"github.com/kitloong/go-echo/routes"
	"github.com/labstack/echo/v4"
)

// RouteServiceProvider register routes with middleware
type RouteServiceProvider struct {
}

// Boot service
func (p *RouteServiceProvider) Boot(e *echo.Echo) {
	routes.Register(e)
}
