load("@rules_pkg//:pkg.bzl", "pkg_tar")
load("@io_bazel_rules_docker//container:image.bzl", "container_image")
load("@io_bazel_rules_docker//contrib:passwd.bzl", "passwd_entry", "passwd_file")
load("@io_bazel_rules_docker//contrib:group.bzl", "group_entry", "group_file")

container_image(
    name = "image",
    base = "@alpine//image",
    cmd = [
        "/app/server",
        "-port",
        "12345",
        "-rest_port",
        "8080",
        "-grpc_port",
        "9090",
    ],
    empty_dirs = ["/tmp"],
    env = {"TMPDIR": "/tmp"},
    ports = [
        "12345",
        "8080",
        "9090",
    ],
    tars = [
        ":passwd_tar",
        ":app_tar",
    ],
    user = "app:app",
    workdir = "/app",
)

group_entry(
    name = "root_group",
    gid = 0,
    groupname = "root",
)

group_entry(
    name = "app_group",
    gid = 0,
    groupname = "app",
)

group_file(
    name = "group",
    entries = [
        ":root_group",
        ":app_group",
    ],
)

passwd_entry(
    name = "root_user",
    create_home = False,
    gid = 0,
    home = "/root",
    info = "root",
    shell = "/bin/bash",
    uid = 0,
    username = "root",
)

passwd_entry(
    name = "app_user",
    create_home = False,
    gid = 10000,
    home = "/app",
    info = "app",
    shell = "/sbin/nologin",
    uid = 10000,
    username = "app",
)

passwd_file(
    name = "passwd",
    entries = [
        ":root_user",
        ":app_user",
    ],
)

pkg_tar(
    name = "passwd_tar",
    srcs = [
        ":group",
        ":passwd",
    ],
    mode = "0o644",
    package_dir = "etc",
)

pkg_tar(
    name = "app_tar",
    srcs = [
        "//cmd/server",
    ],
    mode = "0o755",
    package_dir = "app",
)
