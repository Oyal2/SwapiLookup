BINARY_NAME=finder
MAIN_PACKAGE=./cmd/finder

.PHONY: all build clean test run

all: test build

build:
	go build -o $(BINARY_NAME) -v $(MAIN_PACKAGE)

clean:
	go clean
	rm -f $(BINARY_NAME)

test:
	go test -v ./test/...

run:
	go build -o $(BINARY_NAME) -v $(MAIN_PACKAGE)
	./$(BINARY_NAME)