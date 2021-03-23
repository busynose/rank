SHELL := /bin/bash
BASEDIR = $(shell pwd)

export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export GOPRIVATE=*.gitlab.com
export GOSUMDB=off

# params pass from cmd
APP_BRANCH = "master"
APP_ENV = "dev"
APP_KUBE_CONFIG = "mb-sz-test"

# params stable
APP_NAME=`cat package.json | grep name | head -1 | awk -F: '{ print $$2 }' | sed 's/[\",]//g' | tr -d '[[:space:]]'`
APP_VERSION=`cat package.json | grep version | head -1 | awk -F: '{ print $$2 }' | sed 's/[\",]//g' | tr -d '[[:space:]]'`
COMMIT_ID=`git rev-parse HEAD`
IMAGE_PREFIX="registry.cn-hangzhou.aliyuncs.com/makeblock/${APP_NAME}:v${APP_VERSION}"

fmt:
	gofmt -w .
mod:
	go mod tidy
lint:
	golangci-lint run -c .golangci.yml internal/...
.PHONY: test
test: mod
	go test -gcflags=-l -coverpkg=./... -coverprofile=coverage.data ./...
.PHONY: build
build:
	IMAGE_NAME="${IMAGE_PREFIX}-${APP_BRANCH}"; \
	sh build/package/build.sh ${COMMIT_ID} $$IMAGE_NAME
build-master:
	make build APP_BRANCH=master
build-release:
	make build APP_BRANCH=release
cleanup:
	sh scripts/cleanup.sh
deploy-preview:
	NEW_IMAGE="${APP_NAME}=${IMAGE_PREFIX}-${APP_BRANCH}"; \
	sh deploy/kubernetes/deploy-preview.sh $$NEW_IMAGE ${COMMIT_ID} ${APP_ENV}
.PHONY: deploy
help:
	@echo "fmt - format the source code"
	@echo "mod - go mod tidy"
	@echo "lint - run golangci-lint"
	@echo "test - unit test"
	@echo "build - build docker image"
	@echo "build-master - build docker image for master branch"
	@echo "build-release - build docker image for release branch"