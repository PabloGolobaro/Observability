package controllers

import (
	"context"
	"currency-api/internal/model"
	"currency-api/internal/services"
	"currency-api/internal/store"
	"currency-api/internal/tracing"
	"currency-api/pkg/mongo"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"net/http"
	"strconv"
	"time"
)

//GetCurrency get currency controller
// swagger:operation GET /currency controllers GetCurrency
// ---
// summary: Возвращает курс валют в зависимости от параметров с параметром
// description: От параметров зависит фильтрация ответов, можно получить как все курсы волют по дате, таккурс валют конткретной валюты.
// produces:
//	- application/json
// parameters:
// - name: date
//   in: query
//   description: дата в формате 2021/01/01
//   type: string
//   required: true
// - name: char
//   in: query
//   description: короткое описание валюты(три буквы)
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": GetCurrencyAnswer
//   "400":
//	   "$ref": GetCurrencyBadAnswer
//   "500":
//	   "$ref": GetCurrencyBadAnswer
func GetCurrency(c echo.Context) error {

	serviceContext, span := otel.Tracer(tracing.Name).Start(c.Request().Context(), "/currency")
	defer span.End()
	var response GetCurrencyResponse
	reqID, ok := c.Get("req-id").(string)
	if !ok {
		reqID = "none"
	}
	serviceContext = context.WithValue(serviceContext, "req-id", reqID)
	parsedDate, err := time.Parse("2006/01/02", c.QueryParam("date"))
	if err != nil {
		log.WithFields(log.Fields{
			"req-id":     reqID,
			"controller": "GetCurrency",
			"Error":      err.Error(),
		}).Debug("False bind response")
		return c.JSON(http.StatusInternalServerError, GetCurrencyBadAnswer{Message: err.Error()})
	}
	response.CharCode = c.QueryParam("char_code")
	response.Date = parsedDate.Unix()
	span.SetAttributes(
		attribute.String("Charcode", response.CharCode),
		attribute.String("Date", strconv.Itoa(int(response.Date))),
	)
	log.WithFields(log.Fields{
		"Get Currency": reqID,
		"Response":     response,
	}).Debug("Success bind response")
	currencyStore := store.NewCurrency(mongo.GetDB())

	valutes, err := services.GetCurrency(serviceContext, model.ValuteFilter{CharCode: response.CharCode, Date: response.Date}, currencyStore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GetCurrencyBadAnswer{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, GetCurrencyAnswer{valutes})
}

//GetCurrencyAnswer response for GetCurrency
// swagger:response GetCurrencyAnswer
type GetCurrencyAnswer struct {
	// in: body
	Currency []model.Valute `json:"currency"`
}

//GetCurrencyBadAnswer response for GetCurrency
// swagger:response GetCurrencyBadAnswer
type GetCurrencyBadAnswer struct {
	// in: body
	Message string `json:"message"`
}

//GetCurrencyResponse response
type GetCurrencyResponse struct {
	Date     int64  `json:"date" query:"date"`
	CharCode string `json:"char_code" query:"char_code"`
}
