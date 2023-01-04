BUILD_DIR ?= build
BUILD_PACKAGE ?= ./cmd/main.go

BINARY_NAME ?= doer-api
VERSION ?= $(shell git describe --tags --exact-match 2>/dev/null || git symbolic-ref -q --short HEAD)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)
LDFLAGS += -X internal.buildinfo.version=${VERSION} -X internal.buildinfo.commitHash=${COMMIT_HASH} -X internal.buildinfo.buildDate=${BUILD_DATE}

build:
	go build ${GOARGS} -tags "${GOTAGS}" -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/${BINARY_NAME} ${BUILD_PACKAGE}


.PHONY proto:
proto:
	protoc -I/usr/local/include -I. \
			-I${GOPATH}/src \
			-I./third_party/googleapis \
			-I./third_party \
			--grpc-gateway_out=logtostderr=true,paths=source_relative:.  \
			 --openapiv2_out . \
			--validate_out=lang=go,paths=source_relative:.\
			--go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			./api/v1/doer/*.proto
	$(MAKE) mocks


.PHONY: mocks
mocks:
	./mockgen.sh

.PHONY: generate
generate:
	go generate ./...

serve:
	go run ./cmd/main.go --config config.yaml

test: 
	go test -race -v ./...