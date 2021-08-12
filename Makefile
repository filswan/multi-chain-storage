PROJECT_NAME=payment-bridge
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
BINARY_NAME=payment-bridge

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOBIN=$(shell pwd)/build/bin

.PHONY: all dep build clean test coverage coverhtml lint

all: build

lint: ## Lint the files
	@golangci-lint run --timeout 150m

test: ## Run unittests
	@go test -short ${PKG_LIST}

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

msan: dep ## Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

coverage: ## Generate global code coverage report
	./tools/coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	./tools/coverage.sh html;

dep: ## Get the dependencies
ifeq ($(shell command -v dep 2> /dev/null),)
	$(GOGET) -u -v github.com/golang/dep/cmd/dep
endif
ifeq ($(shell command -v govendor 2> /dev/null),)
	$(GOGET) -u -v github.com/kardianos/govendor
endif
	@dep ensure
	@govendor init
	@govendor add +e
	@rm -rf  ./vendor/github.com/nebulaai/nbai-node/crypto/secp256k1/
	@rm -rf ./vendor/github.com/karalabe/hid
	@govendor fetch -tree  github.com/nebulaai/nbai-node/crypto/secp256k1
	@govendor fetch -tree  github.com/karalabe/hid


build: ## Build the binary file
	@go mod download
	@go mod tidy
	@go build -o off-chain/build/bin/payment-bridge main/main.go
	@mkdir -p ./off-chain/build/bin/off-chain/config
	@mkdir -p ./off-chain/build/bin/on-chain/contracts/abi
	@cp ./off-chain/config/config.toml ./off-chain/build/bin/off-chain/config/config.toml
	@cp ./on-chain/contracts/abi/SwanPayment.json ./off-chain/build/bin/on-chain/contracts/abi/SwanPayment.json

clean: ## Remove previous build
	@go clean
	@rm -rf $(shell pwd)/off-chain/build

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'