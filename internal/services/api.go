package services

import (
	"context"
	"currency-api/internal/model"
	"currency-api/internal/tracing"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

//GetCurrency get currency service
func GetCurrency(ctx context.Context, filer model.ValuteFilter, currency model.ValuteRepository) ([]model.Valute, error) {
	var span trace.Span
	ctx, span = otel.Tracer(tracing.Name).Start(ctx, "GetCurrency service")
	defer span.End()
	reqID, ok := ctx.Value("req-id").(string)
	if !ok {
		reqID = "none"
	}
	log.WithField("GetCurrency Service ", reqID).Debug()

	select {
	case <-ctx.Done():
		log.WithField("GetCurrency service cancel", reqID).Debug()
		return []model.Valute{}, nil
	default:
		valute, err := currency.Find(ctx, filer)
		if err != nil {
			log.WithFields(log.Fields{
				"req-id":  reqID,
				"service": "GetCurrency",
				"Error":   err.Error(),
			}).Debug("Error")
			return nil, err
		}
		return valute, nil
	}
}
