.PHONY: build

build:
	go build ./cmd/bruter

.DEFAULT_GOAL:=build