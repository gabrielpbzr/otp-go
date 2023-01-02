OUTPUT=bin/otp

all: clean build test

build: clean
	mkdir bin
	go build -o $(OUTPUT)

clean:
	rm -rf bin

run:
	./$(OUTPUT)

test:
	go test -v ./...
