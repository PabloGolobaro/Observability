package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

//Config is a configuration for api
type Config struct {
	PORT string `envconfig:"PORT" default:":8080"`
	DB   string `envconfig:"DB" default:"mongodb://root:password123@mongo:27017"`
}

//GetConfig parse env and return Config
func GetConfig() Config {
	var env Config
	var err = envconfig.Process("", &env)
	if err != nil {
		log.Panicf("Bad Init Config with err : %s", err.Error())
		panic(err)
	}
	return env
}
