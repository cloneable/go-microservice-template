export GO111MODULE=on
export PATH:=$(shell go env GOPATH)/bin:$(PATH)

.PHONY: tidy
tidy:
	go mod tidy
	bazelisk run //:gazelle-update-repos
	bazelisk run //:gazelle
