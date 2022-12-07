.PHONY: test
test:
	go test -v ./...

.PHONY: start-server
start-server:
	go run main.go