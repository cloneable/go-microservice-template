build:
	go build ./cmd/server

test:
	go test ./...

clean:
	go clean

tidy:
	go mod tidy

install-tools:
	go list --tags tools -f '{{range .Imports}}{{.|println}}{{end}}' ./tools | xargs go install

generate-protobufs:
	buf generate

.PHONY: build test clean tidy install-tools generate-protobufs
