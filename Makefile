.PHONY: all test clean


test:
	go test ./...

clean:
	rm -f ip_proxy
	rm -f pool/pool