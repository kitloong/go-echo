package session

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/kitloong/go-echo/config"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Store of session
func Store() *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(config.Get("session.secret").(string)))
}

// Get value by key
func Get(key string, c echo.Context) interface{} {
	sess, err := session.Get("session", c)
	if err != nil {
		return nil
	}
	return sess.Values[key]
}

// Put value into key
func Put(key string, value interface{}, c echo.Context) {
	sess, _ := session.Get("session", c)
	sess.Values[key] = value
	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		panic(fmt.Sprintf("Failed to set session key: %s value: %s", key, value))
	}
}
