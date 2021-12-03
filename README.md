# go-microservice-template

## Build

```shell
go install github.com/bazelbuild/bazelisk@latest
```

```shell
bazelisk build //...
```

## Development Setup

You need the `go` tool infrastructure for development and building.
`protoc` if you want to make changes to `.proto` files and regenerate the proto code.

Other project-specific tools written in Go can be installed with `make install-tools`.
Make sure `$GOPATH/bin` is in your `$PATH`.

Optional tools useful during development:

- `dlv`, the Delve debugger
- `shadow`, a `go vet` tool detecting shadowing or variables
- `gopls`, the Go language server used by IDEs
- `gotip`, the bleeding-edge version of the Go tools

## To-Do's

- Example unit/integration/e2e tests
- Profiling
- Coverage
- OpenTelemetry?
- Auth
- GitHub Actions
- Add Docker commands to Makefile?
- Tracing/correlation ID plumbing
- Redis
- RabbitMQ
- ORM
- Use upx for the binary?
  - `RUN apt-get update && apt-get -y install upx`
  - `RUN upx -q -9 /build/out/server`
