load("@io_bazel_rules_k8s//k8s:objects.bzl", "k8s_objects")
load("@k8s_deployment//:defaults.bzl", "k8s_deployment")

k8s_objects(
    name = "cluster",
    objects = [
        ":deployment",
    ],
)

k8s_deployment(
    name = "deployment",
    # images = {
    #     "cloneable/go-microservice-template:dev": "//container:image",
    # },
    template = ":deployment.yaml",
)
