package mongo

import "time"

const (
	connectionTimeout = 10 * time.Second
	errIsNotConnected = "Mongo is not Connected"
)

//IsNotConnected is err when DB is not connected to mongo
var IsNotConnected = Error{Code: "1", Message: errIsNotConnected}
