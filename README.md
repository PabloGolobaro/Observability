# Observability

## Simple REST API with fully instrumented ***observability*** services:
- Prometheus to get metrics
- Jaeger to get Opentelemetry traces
- Grafana to visualize metrics and traces
- Graylog to collect logs from distributed services
- MongoDB as data storage

### This project can be used as an **example** of how to initially set transparency tools in your project
and call it from your code.

---

### Go libraries used:
- github.com/labstack/echo/v4 **router**
- github.com/sirupsen/logrus **logger**
- go.mongodb.org/mongo-driver **mongo driver**
- go.opentelemetry.io/otel/exporters/jaeger
- github.com/prometheus/client_golang

---

## How to start up project:
- Start application with all services:
``make app``
- Start loader to collect data into MongoDB:
``make load``
#### ALl environment variables can be set in Makefile
    ``export PORT := 8080
    export DB_PORT := 27017
    export MONGO_INITDB_ROOT_USERNAME := root
    export MONGO_INITDB_ROOT_PASSWORD := password
    export USERNAME := $(MONGO_INITDB_ROOT_USERNAME)
    export PASSWORD := $(MONGO_INITDB_ROOT_PASSWORD)
    export DB := mongodb://$(USERNAME):$(PASSWORD)@mongo:$(DB_PORT)
    export LOADER_DB := mongodb://localhost:$(DB_PORT)
    export GRAYLOG_HOST := graylog:12201
    export GRAYLOG_PASSWORD_SECRET := somepasswordpepper
    export GRAYLOG_ROOT_PASSWORD_SHA2 := 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918``
#### API endpoint to use:
    /currency&date="2021/01/01"&char=JPY
#### Services:
- Grafana - localhost:3000
- Prometheus - localhost:9090
- Jaeger - localhost:16686
- Jaeger - localhost:16686
- Graylog - localhost:9000
