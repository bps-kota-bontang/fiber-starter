.PHONY: wire run build

wire:
	wire gen ./di

run:
	go run cmd/main.go

build:
	go build -o myapp
