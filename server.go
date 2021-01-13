package main

import (
	"github.com/kitloong/go-echo/bootstrap"
)

func main() {
	e := bootstrap.CreateApplication()

	e.Logger.Fatal(e.Start("localhost:1323"))
}
