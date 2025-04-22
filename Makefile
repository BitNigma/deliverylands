.PHONE: build 
.PHONY: gen
.PHONE: prepare

gen :
	templ generate

prepare:
	export PATH="$PATH:$HOME/go/bin"

build : gen
		go build -o ./delivery cmd/main.go

.PHONE : test
test :
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build