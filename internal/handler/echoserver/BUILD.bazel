load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "echoserver",
    srcs = ["echoserver.go"],
    importpath = "github.com/cloneable/go-microservice-template/internal/handler/echoserver",
    visibility = ["//:__subpackages__"],
    deps = [
        "//api/proto/server",
        "@io_opentelemetry_go_otel_trace//:trace",
        "@org_uber_go_zap//:zap",
    ],
)

go_test(
    name = "echoserver_test",
    size = "small",
    srcs = ["echoserver_test.go"],
    embed = [":echoserver"],
)
