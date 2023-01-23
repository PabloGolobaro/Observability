package main

import (
	"context"
	"currency-api/internal/config"
	"currency-api/internal/logger"
	"currency-api/internal/tracing"
	"currency-api/pkg/mongo"
	"currency-api/pkg/router"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"time"
)

//https://github.com/teris-io/shortid

func init() {
	logger.InitializeLogger()

}

func main() {
	log.WithFields(log.Fields{
		"Start API At": time.Now(),
	}).Info("Start API")

	//configuration
	configuration := config.GetConfig()
	log.WithFields(log.Fields{
		"Config": configuration,
	}).Debug("Config")

	mongo.InitMongo(configuration.DB)
	log.Info("Init MongoDB Success")

	tp, err := tracing.TracerProvider()
	if err != nil {
		log.Fatal("Failed to initialize TracerProvicer: ", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatal("Failed to shutdown TracerProvider: ", err)
		}
	}()
	otel.SetTracerProvider(tp)
	server := router.ConfigureServer(Routes)
	log.Info("Configure Service Success")
	log.Fatal(server.Start(":" + configuration.PORT))
}
