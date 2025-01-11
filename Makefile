export DB_USERNAME=$(shell cat .env | grep ECOMMERCE_DB_USERNAME | cut -d '=' -f2)
export DB_PASSWORD=$(shell cat .env | grep ECOMMERCE_DB_PASSWORD | cut -d '=' -f2)
export DB_HOST=$(shell cat .env | grep ECOMMERCE_DB_HOST | cut -d '=' -f2)
export DB_PORT=$(shell cat .env | grep ECOMMERCE_DB_PORT | cut -d '=' -f2)
export DB_NAME=$(shell cat .env | grep ECOMMERCE_DB_NAME | cut -d '=' -f2)

build:
	@go build -o ./bin/main ./cmd/api/main.go
run:build
	@./bin/main

create-migration:
	@migrate create -ext=sql -dir=./migrations -seq $(name)
run-migrations:
	@migrate -path ./migrations -database postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable up

rollback-migrations:
	@migrate -path ./migrations -database postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable down