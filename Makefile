BUILDPATH=$(CURDIR)
API_NAME=api-go-test

build:
	@echo "Creando Binario ..."
	@go build -mod=vendor -ldflags '-s -w' -o $(BUILDPATH)/build/bin/${API_NAME} cmd/main.go
	@echo "Binario generado en build/bin/${API_NAME}"

test:
	@echo "Ejecutando tests..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

coverage:
	@echo "Coverfile..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -func coverfile_out | grep total | awk '{print substr($$3, 1, length($$3)-1)}' > coverage.txt

runDev:
	@docker-compose up -d
	@go run cmd/main.go

.PHONY: test build