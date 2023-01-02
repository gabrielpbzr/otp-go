OUTPUT=bin/otp

all: clean build

build: clean
	mkdir bin
	go build -o $(OUTPUT)

clean:
	rm -rf bin

run:
	./$(OUTPUT)

test:
	go test -v ./...
