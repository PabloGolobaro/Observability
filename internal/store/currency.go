package store

import (
	"context"
	"currency-api/internal/model"
	"currency-api/internal/tracing"
	"currency-api/pkg/mongo"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

//Currency store for valute
type Currency struct {
	coll  string
	store *mongo.DB
}

//Find
func (c Currency) Find(ctx context.Context, filter model.ValuteFilter) ([]model.Valute, error) {
	var span trace.Span
	ctx, span = otel.Tracer(tracing.Name).Start(ctx, "Currency.Find method")
	defer span.End()
	reqID, ok := ctx.Value("req-id").(string)
	if !ok {
		reqID = "none"
	}
	log.Debug("Find Currency", reqID)

	select {
	case <-ctx.Done():
		log.Debug("Find Currency cancwel", reqID)
		return []model.Valute{}, nil
	default:
		query, err := prepareQuery(filter)
		if err != nil {
			return nil, err
		}
		var currency []model.Valute
		err = c.store.FindWithQueryAll(c.coll, query, &currency)
		if err != nil {
			return nil, err
		}
		return currency, nil
	}
}

//Store
func (c Currency) Store(valute model.Valute) error {
	return nil
}

//prepareQuery
func prepareQuery(filter model.ValuteFilter) (bson.M, error) {
	if filter.Date != 0 && len(filter.CharCode) != 0 {
		charCodeQuery := bson.M{
			"char_code": filter.CharCode,
		}
		dateQuery := bson.M{
			"date": filter.Date,
		}
		query := bson.M{
			"$and": []bson.M{charCodeQuery, dateQuery},
		}
		return query, nil
	}
	if filter.Date != 0 {
		dateQuery := bson.M{
			"date": filter.Date,
		}
		return dateQuery, nil
	}
	if len(filter.CharCode) != 0 {
		charCodeQuery := bson.M{
			"char_code": filter.CharCode,
		}
		return charCodeQuery, nil
	}
	return bson.M{}, fmt.Errorf("all filter param empty")
}

//NewCurrency
func NewCurrency(db *mongo.DB) *Currency {
	return &Currency{store: db, coll: "valute"}
}
