.PHONY: build run

build:
	go build -o build/school_management


run: build
	./build/school_management
