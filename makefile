.DEFAULT_GOAL := help

BIN_NAME := "pocketshorten"
MAIN_BRANCH := main
HEAD_BRANCH := HEAD
ifeq ($(strip $(VERSION_HASH)),)

# hash of current commit
VERSION_HASH := $(shell git rev-parse --short HEAD)

# tag matching current commit or empty
HEAD_TAG := $(shell git tag --points-at HEAD)

#name of branch
BRANCH_NAME := $(shell git rev-parse --abbrev-ref HEAD)
endif

VERSION_STRING := $(BRANCH_NAME)

#if we are on HEAD and there is a tag pointing at head, use that for version else use branch name as version
ifeq ($(BRANCH_NAME),$(HEAD_BRANCH))

$(info match head)
ifneq ($(strip $(HEAD_TAG)),)
VERSION_STRING := $(HEAD_TAG)
$(info    $(version_string))
endif
endif

BINDIR    := $(CURDIR)/bin
PLATFORMS := linux/amd64/rk-Linux-x86_64 darwin/amd64/rk-Darwin-x86_64 windows/amd64/rk.exe linux/arm64/rk-Linux-arm64 darwin/arm64/rk-Darwin-arm64
# dlv exec ./bin/${BIN_NAME} --headless --listen=:2345 --log --api-version=2 -- testserver --loglevel=debug
BUILDCOMMANDDEBUG := go build -gcflags "all=-N -l" -tags "osusergo netgo static_build"
BUILDCOMMAND := go build -trimpath -ldflags "-s -w -X github.com/clarkezone/${BIN_NAME}/pkg/config.VersionHash=${VERSION_HASH} -X github.com/clarkezone/pocketshorten/pkg/config.VersionString=${VERSION_STRING}" -tags "osusergo netgo static_build"
temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))
label = $(word 3, $(temp))

UNAME := $(shell uname)
ifeq ($(UNAME), Darwin)
SHACOMMAND := shasum -a 256
else
SHACOMMAND := sha256sum
endif

### -------------------------------------
### Help
### -------------------------------------

##@ Help:

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n    ", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: install-binaries

### -------------------------------------
### Install
### -------------------------------------

##@ Install:

install-protoc-go: ## install protoc compiler
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

install-tools:     ## install tooling
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
	go install github.com/uw-labs/strongbox@latesta && \
	go install github.com/mgechev/revive@latest

.PHONY: test
test:
	go test -p 4 -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: dep
dep:
	go mod tidy

.PHONY: latest
latest:
	echo ${VERSION_STRING} > bin/latest

.PHONY: lint
lint:
	revive $(shell go list ./...)
	go vet $(shell go list ./...)
	golangci-lint run

### -------------------------------------
### Scan Manifests
### -------------------------------------

##@ Kubescore:
.PHONY: kubescore
kubescore:## run manifest scans for kubernetes manifests
	kubectl kustomize k8s/layered_viper/overlay/staging > /tmp/pocketshorten_layered_staging_merged.yaml
	kubectl score /tmp/pocketshorten_layered_staging_merged.yaml

##@ Config validation
.PHONY: validateconfig
cuevalidate: ##Run cue against configuration to validate
	cue vet testfiles/.pocketshorten.json testfiles/schema.cue

.PHONY: precommit-installhooks
precommit:
	pre-commit install

.PHONY: precommit
precommit:
	pre-commit run --all-files

.PHONY: build
build:
	$(BUILDCOMMAND) -o ${BINDIR}/${BIN_NAME}

.PHONY: buildproto
buildproto:
	mkdir -p pkg/greetingservice && protoc -I ./proto/greeter --go_out=pkg/greetingservice --go_opt=paths=source_relative \
          --go-grpc_out=pkg/greetingservice --go-grpc_opt=paths=source_relative \
          ./proto/greeter/*.proto
	mkdir -p pkg/cacheloaderservice && protoc -I ./proto/cacheloader --go_out=pkg/cacheloaderservice --go_opt=paths=source_relative \
          --go-grpc_out=pkg/cacheloaderservice --go-grpc_opt=paths=source_relative \
          ./proto/cacheloader/*.proto

.PHONY: buildimage
buildimage:
	$(eval IMG := "pocketshorten")
	$(eval VERSION := "latest")

	@echo ${IMG}
	@echo ${VERSION}

	-podman manifest exists localhost/$(IMG):latest && podman manifest rm localhost/$(IMG):latest

	podman build --arch=amd64 --build-arg BUILD_HEADTAG="${HEAD_TAG}" --build-arg BUILD_HASH="${VERSION_HASH}" --build-arg BUILD_BRANCH="${HEAD_BRANCH}" -t ${IMG}:${VERSION}.amd64 -f Dockerfile
	podman build --arch=arm64 --build-arg BUILD_VERSION="${HEAD_TAG}" --build-arg BUILD_HASH="${VERSION_HASH}" --build-arg BUILD_BRANCH="${HEAD_BRANCH}" -t ${IMG}:${VERSION}.arm64 -f Dockerfile

	podman manifest create ${IMG}:${VERSION}
	podman manifest add ${IMG}:${VERSION} containers-storage:localhost/${IMG}:${VERSION}.amd64
	podman manifest add ${IMG}:${VERSION} containers-storage:localhost/${IMG}:${VERSION}.arm64

.PHONY: builddlv
builddlv:
	$(BUILDCOMMANDDEBUG) -o ${BINDIR}/{BIN_NAME}

.PHONY: release
build-all: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) CGO_ENABLED=0 $(BUILDCOMMAND) -o "bin/$(label)"
	$(SHACOMMAND) "bin/$(label)" > "bin/$(label).sha256"

### -------------------------------------
### Run shortener locally
### -------------------------------------

##@ Run locally:
.PHONY: kubescore
runlocal:## Run pocketshorten locally using the test configuration
	`./bin/pocketshorten servefrontend --config testfiles/.pocketshorten.json --loglevel=debug`
