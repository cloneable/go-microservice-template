load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/cloneable/go-microservice-template
# gazelle:build_file_name BUILD
# gazelle:proto package

gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=go_dependencies.bzl%go_dependencies",
        "-prune",
        # "-build_file_proto_mode=disable_global",
    ],
    command = "update-repos",
)