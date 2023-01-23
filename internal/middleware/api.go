package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	reqGot = promauto.NewCounter(prometheus.CounterOpts{
		Name: "requests_got",
		Help: "The total number of requests sent to server",
	})
)

//GetCurrencyReqID generate req id
func GetCurrencyReqID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqID := uuid.New().String()
		c.Set("req-id", reqID)
		reqGot.Inc()
		return next(c)
	}
}
