BIN="./bin"
SRC=$(shell find . -name "*.go")

.PHONY: test clean

default: all

all: test_go test_api test_cli

test_go: test
	$(info ******************** running go tests ********************)
	go build -v ./...
	go test -v ./...

test_api: test
	$(info ******************** running api test ********************)
	./mon --version
	

test_cli: test
	$(info ******************** running cli test ********************)
	./mon sensors list
	./mon notify list

clean:
	rm -rf $(BIN)