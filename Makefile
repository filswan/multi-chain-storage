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
	@echo "Done lint."

test: ## Run unittests
	@go test -short ${PKG_LIST}
	@echo "Done testing."

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
	@echo "Done dep."

ffi:
	./extern/filecoin-ffi/install-filcrypto
.PHONY: ffi

build: ## Build the binary file
	@go mod download
	@go mod tidy
	@go build -o build/payment-bridge main/payment-bridge.go
	@mkdir -p ./build/config ./build/config/bsc ./build/config/goerli ./build/config/nbai ./build/config/polygon
	@mkdir -p ./build/on-chain/contracts/abi
	@cp ./config/config.toml.example ./build/config/config.toml
	@cp ./config/bsc/config_bsc.toml.example ./build/config/bsc/config_bsc.toml
	@cp ./config/goerli/config_goerli.toml.example ./build/config/goerli/config_goerli.toml
	@cp ./config/nbai/config_nbai.toml.example ./build/config/nbai/config_nbai.toml
	@cp ./config/polygon/config_polygon.toml.example ./build/config/polygon/config_polygon.toml
	@cp ./on-chain/contracts/abi/SwanPayment.json ./build/on-chain/contracts/abi/SwanPayment.json
	@echo "Done building."

clean: ## Remove previous build
	@go clean
	@rm -rf $(shell pwd)/build
	@echo "Done cleaning."

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
