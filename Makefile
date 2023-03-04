.PHONY: build

run:
	go run .

build:
	go build -o build/ . && \
	cp -r ./templates ./build