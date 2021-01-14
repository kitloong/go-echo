package main

import (
	"fmt"
	"github.com/kitloong/go-echo/bootstrap"
	"github.com/kitloong/go-echo/config"
)

func main() {
	e := bootstrap.CreateApplication()

	fmt.Println(config.Get("app.name"))
	fmt.Println(config.Get("app.test"))
	fmt.Println(config.Get("app.testsss"))

	e.Logger.Fatal(e.Start("localhost:1323"))
}
