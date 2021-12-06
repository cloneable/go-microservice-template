# go-microservice-template

Go + gRPC + gRPC-Gateway + Bazel + Tiltfile

## Setup

```shell
go install github.com/bazelbuild/bazelisk@latest

brew install kubectl
brew install ctlptl
brew install kind
brew install tilt
brew install jsonnet
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

```shell
jsonnetfmt --string-style d -i *.jsonnet
```

## To-Do's

- [ ] helm charts + jsonnet
- [ ] Example unit/integration/e2e tests
- [ ] Profiling
- [ ] Coverage
- [x] OpenTelemetry?
- [ ] Auth
- [x] GitHub Actions
- [ ] Add Docker commands to Makefile?
- [x] Tracing/correlation ID plumbing
- [ ] ORM: Ent?
- [ ] Use upx for the binary?
