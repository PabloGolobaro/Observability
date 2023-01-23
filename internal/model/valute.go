package model

import "context"

//Valute currency
// swagger:response Valute
type Valute struct {
	// in: body
	// id currency
	ID       string  `json:"id" bson:"_id"`
	Name     string  `json:"name" bson:"name"`
	CharCode string  `json:"char_code" bson:"char_code"`
	Nominal  float64 `json:"nominal" bson:"nominal"`
	Date     int64   `json:"date" bson:"date"`
	Value    float64 `json:"value" bson:"value"`
}

//ValuteRepository repo for currency
type ValuteRepository interface {
	Find(ctx context.Context, filter ValuteFilter) ([]Valute, error)
	Store(Valute) error
}

//ValuteFilter param for search currency
type ValuteFilter struct {
	Date     int64
	CharCode string
}
