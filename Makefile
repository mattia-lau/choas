.PHONY: test
test:
	go test ./test -v

.PHONY: start-server
start-server:
	go run main.go