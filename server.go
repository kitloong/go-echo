package main

import (
	"fmt"
	config2 "github.com/kitloong/go-echo/app/helpers/config"
	"github.com/kitloong/go-echo/bootstrap"
)

func main() {
	e := bootstrap.CreateApplication()

	fmt.Println(config2.Get("app.name"))
	fmt.Println(config2.Get("app.test"))
	fmt.Println(config2.Get("app.testsss"))

	e.Logger.Fatal(e.Start("localhost:1323"))
}
