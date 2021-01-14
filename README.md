# Setup

```bash
go version      
go version go1.14.13 darwin/amd64

go build
go run server.go

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.1.17
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323

```

or with air (Hot reload)

```bash
air -c .air.toml

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.1.17
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323

```

## Structure

```
- app
    - helpers
        - config
            config.go
        - session
            session.go
    - http
        - controllers
            - user.go
        - middleware
    - models
        - user.go
    - providers
        - route_service_provider.go
    - repositories
    - services
- bootstrap
    - app.go
- config
    - app.yaml
    - database.yaml
    - ...
- resources
    - lang
        - en
        - ja
- routes
    - api.go
.env
go.mod
server.go
```

# References

## Hot reload

1. https://github.com/cosmtrek/air

## Code practice

1. https://github.com/golang/lint
1. https://talks.golang.org/2013/bestpractices.slide#1

## Go mod

1. https://blog.golang.org/using-go-modules
1. https://github.com/labstack/echo/issues/1374

## Struct

1. https://medium.com/rungo/structures-in-go-76377cc106a2
1. https://medium.com/rungo/anatomy-of-methods-in-go-f552aaa8ac4a

## Interface

1. https://golang.org/doc/faq#Is_Go_an_object-oriented_language
1. https://medium.com/rungo/interfaces-in-go-ab1601159b3a
1. https://medium.com/@dotronglong/interface-naming-convention-in-golang-f53d9f471593

## Packages

1. https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc

## Pointer

1. http://piotrzurek.net/2013/09/20/pointers-in-go.html#:~:text=*%20in%20front%20of%20a%20type,speak%20this%20is%20called%20dereferencing.

## ORM

1. https://gorm.io/index.html

## Configuration

1. https://github.com/spf13/viper

## String, rune, byte

1. https://blog.golang.org/strings

## Type assertion

1. https://tour.golang.org/methods/15