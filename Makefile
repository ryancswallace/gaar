NAME = gaar
GO = go
BIN = ./bin

.PHONY: all
all: format test build install

.PHONY: build
build:
	mkdir -p $(BIN) \
	&& $(GO) build -o $(BIN)/$(NAME) main.go

.PHONY: install
install:
	$(GO) install

.PHONY: format
format:
	$(GO) fmt

.PHONY: test
test:
	$(GO) test -v ./cmd