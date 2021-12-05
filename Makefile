export GO111MODULE=on
export PATH:=$(shell go env GOPATH)/bin:$(PATH)

.PHONY: tidy
tidy:
	go mod tidy
	bazel run //:gazelle-update-repos
	bazel run //:gazelle
