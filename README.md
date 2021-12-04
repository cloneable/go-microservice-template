# go-microservice-template

## Setup

```shell
go install github.com/bazelbuild/bazelisk@latest

brew install kubectl
brew install ctlptl
brew install kind
brew install tilt
```

```shell
ctlptl create registry ctlptl-registry --port=5005
ctlptl create cluster kind --registry=ctlptl-registry
```

## Commands

```shell
bazelisk build //...
bazelisk test //...
```

```shell
tilt up
^C

tilt down
```

```shell
bazelisk build -c opt --stamp --strip=always --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //container:image.tar
docker load -i bazel-bin/container/image.tar
```

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
