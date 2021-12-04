load("@io_bazel_rules_k8s//k8s:objects.bzl", "k8s_objects")
load("@k8s_deployment//:defaults.bzl", "k8s_deployment")
load("@io_bazel_rules_jsonnet//jsonnet:jsonnet.bzl", "jsonnet_to_json")

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
    template = ":deployment_yaml",
)

jsonnet_to_json(
    name = "deployment_yaml",
    src = "deployment.jsonnet",
    outs = ["deployment.yaml"],
    yaml_stream = 1,
)
