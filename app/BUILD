load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//container:container.bzl", "container_image")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/potix2/gke-mysql-benchmark/app",
    visibility = ["//visibility:private"],
    deps = ["@com_github_go_sql_driver_mysql//:go_default_library"],
)

go_binary(
    name = "app",
    importpath = "github.com/potix2/gke-mysql-benchmark/app",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

container_image(
    name = "devel",
    base = "@ubuntu//image:image.tar",
    cmd = ["/app"],
    files = [
        ":app",
    ],
)
