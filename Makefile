UNAME := $(shell uname)
CWD := ${CURDIR}

BINDATA := $(shell command -v go-bindata 2> /dev/null)
VERSION := $(shell cat Version)
PROTOC_GEN_GO := $(shell command -v protoc-gen-go 2> /dev/null)
BITBUCKET_COMMIT ?= $(shell git rev-parse --verify HEAD)
BITBUCKET_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
BRANCH ?= ${BITBUCKET_BRANCH}
HASH ?= ${BITBUCKET_COMMIT}
STAGE := $(shell cli/deploy -e ${BRANCH} -v 2> /dev/null)

bindata:
ifndef BINDATA
	go get -u github.com/jteeuwen/go-bindata/...
endif
	@echo ${VERSION}-b${HASH} > resources/VERSION
	@go-bindata  -pkg utils -o lib/utils/resources.go resources/sql/...

migrate-up:
	@go run cli/stovia_service.go migrate up

build:
	@dep ensure
	@go get -v $(shell go list ./... | grep -v /vendor/)
	@go build -race -v -o resources/deploy/${ENV}/stovia-service cli/stovia_service.go
