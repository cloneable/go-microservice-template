load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "healthz",
    srcs = ["healthz.go"],
    importpath = "github.com/cloneable/go-microservice-template/internal/handler/healthz",
    visibility = ["//:__subpackages__"],
    deps = [
        "@go_googleapis//google/api:httpbody_go_proto",
        "@org_golang_google_protobuf//types/known/emptypb",
    ],
)
