.PHONY: all build test clean

all: build

build:
	go build -o proxy_pool ./concretepool
	go build -o webserver main.go
test:
	go test ./...

clean:
	rm -f proxy_pool
	rm -f webserver