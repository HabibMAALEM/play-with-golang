SHELL := /bin/bash

GO ?= go
GO_TEST=$(GO) test

.PHONY: test
test:
	$(GO_TEST) ./...
