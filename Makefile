.PHONY: build
build:
	mkdir -p ./build/
	go build -o ./build/grosh ./main.go
