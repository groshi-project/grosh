BIN_DEST := /usr/bin/grosh

.PHONY: all
all: build

.PHONY: build
build:
	mkdir -p ./build/
	go build -o ./build/grosh ./main.go

.PHONY: install
install:
	install -m 755 ./build/grosh $(BIN_DEST)

.PHONY: uninstall
uninstall:
	rm -f $(BIN_DEST)