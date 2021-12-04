BAZEL_RUN = "bazelisk run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64"

def bazel_k8s(target):
  return local("{bazel_run} {target}".format(bazel_run=BAZEL_RUN, target=target))

def bazel_build(image, target):
  dest = target.replace('//', 'bazel/')
  custom_build(
    image,
    "{bazel_run} {target} -- --norun && docker tag {dest} $EXPECTED_REF".format(bazel_run=BAZEL_RUN, target=target, dest=dest),
    [],
    tag="dev",
  )

k8s_kind("Deployment", image_json_path="{.spec.template.spec.containers[0].image}")

k8s_yaml(bazel_k8s("//k8s:deployment"))

bazel_build("cloneable/go-microservice-template", "//container:image")

k8s_resource("go-microservice-template", port_forwards=[8080,9090,12345])
