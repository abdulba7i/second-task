BINARY=bin/minishell

all: build

build:
	mkdir -p bin
	go build -o $(BINARY) ./cmd/minishell.go

run: build
	./$(BINARY)

clean:
	rm -rf bin 