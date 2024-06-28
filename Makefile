PROJECT=go-biwywfok
VERSION=0.0.2

.PHONY: all build test vendor

build:
	go build -o builds/ -mod=vendor ./cmd/go-biwywfok 

vendor: 
	go mod tidy && go mod vendor

up: build
	./builds/go-biwywfok

docker-lint:
	docker run --name $(PROJECT)-lint --rm -i -v "$(PWD):/src" -w /src golangci/golangci-lint:v1.59 golangci-lint run ./... -E gofmt -E revive -E dupl -E gocritic -E nestif -E errorlint -E bodyclose -E gosec -E goconst -E unparam -E unconvert -E asciicheck -E exportloopref -E nilerr --timeout=10m
