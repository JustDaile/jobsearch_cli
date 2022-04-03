.DEFAULT_GOAL := run
APPNAME ?= $(shell cat .build | grep -oP "name=\K.+")
SITENAME ?= default
ENV ?= dev
VERSION := $(shell cat version.txt)
HOSTOS := $(shell go env | grep -oP "GOHOSTOS=\K\w+")
HOSTARCH := $(shell go env | grep -oP "GOHOSTARCH=\K\w+")
PACKAGE_PATH := $(shell cat go.mod | grep -oP "module \K.*")
TIMESTAMP := $(shell date +"%Y%m%d-%H%M%S")
LDFLAGS := -ldflags='-X $(PACKAGE_PATH)/internal.AppName=$(APPNAME) -X $(PACKAGE_PATH)/internal.SiteName=$(SITENAME) -X $(PACKAGE_PATH)/internal.Environment=$(ENV) -X $(PACKAGE_PATH)/internal.Version=$(VERSION) -X $(PACKAGE_PATH)/internal.Timestamp=$(TIMESTAMP)'

.PHONY: all

# List prints all targets in this makefile.
list:
	@make -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

# Install all golang dependencies for this project.
deps:
	go get github.com/PuerkitoBio/goquery

# Clean binary - destroys all compiled binaries.
clean:
	@rm -rf bin
	@go mod tidy

# Compile for windows amd64.
compile.windows64: deps
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o bin/$(APPNAME)-$(VERSION)-windows-arm64 main.go

# Compile for linux amd64.
compile.linux64: deps
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o bin/$(APPNAME)-$(VERSION)-linux-arm64 main.go

# Compile for linux 386.
compile.linux386: deps
	@GOOS=linux GOARCH=386 CGO_ENABLED=0 go build $(LDFLAGS) -o bin/$(APPNAME)-$(VERSION)-linux-arm64 main.go

# Run the main.go without building it.
# trap make process to ensure its killed on ctrl+c
run: deps
	@bash -c "trap 'ps | grep make | grep -oP \"\d+\" | head -1' EXIT; \
		DB_HOST=$(DB_HOST) DB_PASSWORD=$(DB_PASSWORD) DB_NAME=$(DB_NAME) DB_USER=$(DB_USER) go run $(LDFLAGS) main.go $(ARGS)"

# Run all go tests
test: deps
	@go test ./... $(ARGS)