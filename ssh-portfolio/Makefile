GO ?= go
.PHONY: fmt test build run check
fmt:
	$(GO) fmt ./...
test:
	$(GO) test -race ./...
build:
	mkdir -p bin
	$(GO) build -buildvcs=false -trimpath -o bin/ssh-portfolio .
run:
	$(GO) run -buildvcs=false .
check: fmt
	$(GO) vet ./...
	$(GO) test -race ./...
	$(MAKE) build
