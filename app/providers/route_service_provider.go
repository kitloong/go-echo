package providers

import (
	sess "github.com/kitloong/go-echo/app/helpers/session"
	"github.com/kitloong/go-echo/routes"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// RouteServiceProvider register routes with middleware
type RouteServiceProvider struct {
}

// Boot service
func (p *RouteServiceProvider) Boot(e *echo.Echo) {
	g := e.Group("/api")
	g.Use(session.Middleware(sess.Store()))

	routes.API(g)
}
