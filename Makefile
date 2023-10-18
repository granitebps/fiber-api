APP_NAME = api
BUILD_DIR = $(PWD)/build

swag:
	swag init

air:
	air

build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) .

lint:
	golangci-lint run ./...

security:
	gosec ./...

critic:
	gocritic check -enableAll ./...