default: help

## translate lib
##   AUTHOR: karazhanov
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

## make lint - check rules
lint:
	golangci-lint run

## make test - run all tests
test: lint
	go test ./...

## make run - run app
run: test
	go run main.go