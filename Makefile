export GO111MODULE=on
export PATH:=$(shell go env GOPATH)/bin:$(PATH)

.PHONY: tidy
tidy:
	go mod tidy
	bazel run //:gazelle-update-repos
	bazel run //:gazelle

.PHONY: jsonnetfmt
jsonnetfmt:
	find -E . -type f -regex ".*[.](j|lib)sonnet" -print0  | xargs -0 jsonnetfmt --string-style d --max-blank-lines 1 -i
