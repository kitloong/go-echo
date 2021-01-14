package routes

import (
	"github.com/kitloong/go-echo/app/http/controllers"
	"github.com/labstack/echo/v4"
)

// API routes
func API(api *echo.Group) {
	// Define api routes
	api.GET("/users", controllers.User().Index)
	api.GET("/users/:id", controllers.User().Show)
}
