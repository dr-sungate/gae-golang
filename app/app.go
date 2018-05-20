// +build appengine

package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/appengine"
	"gaeapp/gae-golang/api/handler"
)

func main() {

	e := echo.New()
	e.Debug = true
	g := e.Group("/users")
	g.Use(middleware.CORS())

	g.POST("", handler.Users{}.CreateUser)
	g.GET("", handler.Users{}.GetUsers)
	g.GET("/:id", handler.Users{}.GetUser)
	
	http.Handle("/", e)
	appengine.Main()
}