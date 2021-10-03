RELEASE_FLAGS=-trimpath -ldflags '-s -w -extldflags "-static"'
BUILD_OUTPUT=build

export PATH:=$(shell go env GOPATH)/bin:$(PATH)

# TODO: different output dirs for debug/release?

export GO111MODULE=on

# Default rule creates regular binaries with some debug info.
.PHONY: debug
debug:
	@mkdir -p $(BUILD_OUTPUT)
	go build -o $(BUILD_OUTPUT) ./...

# Run all tests.
.PHONY: test
test:
	go test ./...

# Make release binaries.
# Used in Dockerfile.
.PHONY: release
release: export CGO_ENABLED=0
release:
	@mkdir -p $(BUILD_OUTPUT)
	go build $(RELEASE_FLAGS) -o $(BUILD_OUTPUT) ./...

# Run all tests with same flags as release build.
# Used in Dockerfile.
.PHONY: release-test
release-test: export CGO_ENABLED=0
release-test:
	go test $(RELEASE_FLAGS) ./...

# Add/remove modules to/from go.mod/go.sum and download direct deps.
.PHONY: tidy
tidy:
	go mod tidy

# Install tools needed by project during code generate phase.
.PHONY: install-tools
install-tools:
	go list --tags tools -f '{{range .Imports}}{{.|println}}{{end}}' ./tools | xargs go install

# Re-generate the source files after changes to .proto files.
.PHONY: generate-protobufs
generate-protobufs:
	buf generate
