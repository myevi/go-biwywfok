.PHONY: all build test vendor

build:
	go build -o builds/ -mod=vendor ./cmd/go-biwywfok 

vendor: 
	go mod tidy && go mod vendor

up: build
	./builds/go-biwywfok