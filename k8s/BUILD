load("@io_bazel_rules_k8s//k8s:object.bzl", "k8s_object")
load("@io_bazel_rules_k8s//k8s:objects.bzl", "k8s_objects")

k8s_objects(
    name = "cluster",
    objects = [
        ":deployment",
    ],
)

k8s_object(
    name = "deployment",
    # images = {
    #     "cloneable/go-microservice-template:dev": "//container:image",
    # },
    kind = "deployment",
    template = ":deployment.yaml",
)
