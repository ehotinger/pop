GO := go
PREFIX?=$(shell pwd)
PACKAGES=$(shell go list ./... | grep -v /vendor/)
GO_GCFLAGS=$(shell				\
	set -- ${GOPATHS};			\
	echo "-gcflags=-trimpath=$${1}/src";	\
	)
GO_BUILD_FLAGS=
GO_EXTRA_FLAGS=
GO_TAGS=
VERSION=$(shell git describe --match 'v[0-9]*' --dirty='.m' --always)
GITCOMMIT=$(shell git rev-parse HEAD)$(shell if ! git diff --no-ext-diff --quiet --exit-code; then echo .m; fi)
PKG=github.com/ehotinger/pop
GO_LDFLAGS=-ldflags '-s -w -X $(PKG)/version.Version=$(VERSION) -X $(PKG)/version.Revision=$(GITCOMMIT)'
COMMANDS=pop
BINARIES=$(addprefix bin/,$(COMMANDS))
INSTALLDIR=/usr/local

.PHONY: binaries
binaries: $(BINARIES) ## Build the binaries
	@echo "+ $@"

FORCE:
bin/%: cmd/% FORCE
	@echo "+ $@${BINARY_SUFFIX}"
	@CGO_ENABLED=0 go build ${GO_GCFLAGS} ${GO_BUILD_FLAGS} -o $@${BINARY_SUFFIX} ${GO_LDFLAGS} ${GO_TAGS} ./$<


.PHONY: build
build: ## Build the Go packages
	@echo "+ $@"
	@go build ${GO_GCFLAGS} ${GO_BUILD_FLAGS} ${GO_EXTRA_FLAGS} ${GO_LDFLAGS} ${PACKAGES}

.PHONY: clean
clean: ## Clean up all binaries
	@echo "+ $@"
	@rm -f $(BINARIES)

.PHONY: lint
lint: ## Run all linters
	@echo "+ $@"
	golangci-lint run


.PHONY: test
test: ## Runs the Go tests
	@echo "+ $@"
	@go test -v -tags "$(BUILDTAGS) cgo" $(shell go list ./... | grep -v vendor)

.PHONY: help
help: ## Prints this help menu
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort