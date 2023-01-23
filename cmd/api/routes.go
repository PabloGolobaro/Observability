// currency-api
//
// Сервис для получения курса валют.
//
//
//
//     Schemes: http
//     Host: localhost
//     Version: 0.0.1
//     Contact: nazemnov.g.a@gmail.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
// swagger:meta
package main

import (
	"currency-api/internal/controllers"
	"currency-api/internal/middleware"
	"currency-api/pkg/router"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

//Routes is routes app
var Routes = []router.Route{
	{
		Path:       "/currency",
		Controller: controllers.GetCurrency,
		Method:     http.MethodGet,
		Middleware: []echo.MiddlewareFunc{middleware.GetCurrencyReqID},
	},
	{
		Path:       "/prometheus",
		Controller: echo.WrapHandler(promhttp.Handler()),
		Method:     http.MethodGet,
	},
}
