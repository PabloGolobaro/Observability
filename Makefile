export PORT := 8080
export DB_PORT := 27017
export MONGO_INITDB_ROOT_USERNAME := root
export MONGO_INITDB_ROOT_PASSWORD := password123
export USERNAME := $(MONGO_INITDB_ROOT_USERNAME)
export PASSWORD := $(MONGO_INITDB_ROOT_PASSWORD)
export DB := mongodb://$(USERNAME):$(PASSWORD)@mongo:$(DB_PORT)
export LOADER_DB := mongodb://localhost:$(DB_PORT)
export GRAYLOG_HOST := graylog:12201
export GRAYLOG_PASSWORD_SECRET := somepasswordpepper
export GRAYLOG_ROOT_PASSWORD_SHA2 := 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918

run:
	@go run ./cmd/api
app:
	@docker-compose up -d
load:
	@go run ./cmd/loader
