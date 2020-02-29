.PHONY: all build test clean

all: build

build:
	go build -o proxy_pool ./concretepool
test:
	go test ./...

clean:
	rm -f ip_proxy
	rm -f proxy_pool