package main

import (
	"github.com/kitloong/go-echo/bootstrap"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	bootstrap.CreateApplication(e)

	e.Logger.Fatal(e.Start("localhost:1323"))
}
