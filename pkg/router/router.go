package router

import "github.com/labstack/echo/v4"

//Route endpoints and controller api
type Route struct {
	Controller echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
	Path       string
	Method     string
}

//ConfigureServer return echo server
func ConfigureServer(routers []Route) *echo.Echo {
	server := echo.New()
	for _, route := range routers {
		server.Add(route.Method, route.Path, route.Controller, route.Middleware...)
	}
	server.Debug = false
	return server
}
