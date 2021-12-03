load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "healthz_proto",
    srcs = ["healthz.proto"],
    visibility = ["//visibility:public"],
    deps = [
        "@com_google_protobuf//:empty_proto",
        "@go_googleapis//google/api:annotations_proto",
        "@go_googleapis//google/api:httpbody_proto",
    ],
)

# gazelle:go_grpc_compilers @io_bazel_rules_go//proto:go_grpc,@com_github_grpc_ecosystem_grpc_gateway_v2//protoc-gen-grpc-gateway:go_gen_grpc_gateway
go_proto_library(
    name = "healthz_go_proto",
    compilers = [
        "@io_bazel_rules_go//proto:go_grpc",
        "@com_github_grpc_ecosystem_grpc_gateway_v2//protoc-gen-grpc-gateway:go_gen_grpc_gateway",
    ],
    importpath = "github.com/cloneable/go-microservice-template/api/proto/healthz",
    proto = ":healthz_proto",
    visibility = ["//visibility:public"],
    deps = [
        "@go_googleapis//google/api:annotations_go_proto",
        "@go_googleapis//google/api:httpbody_go_proto",
    ],
)
