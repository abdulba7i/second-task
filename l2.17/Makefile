BINARY_NAME=telnet-client

.PHONY: build run clean help

build:
	go build -o $(BINARY_NAME) main.go

clean:
	rm -f $(BINARY_NAME)
	go clean