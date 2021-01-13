package routes

import (
	"github.com/kitloong/go-echo/app/http/controllers"
	"github.com/labstack/echo/v4"
)

// Register routes
func Register(e *echo.Echo) {
	// Define routes
	e.GET("/", controllers.User().Index)
	e.GET("users/:id", controllers.User().Show)
}
