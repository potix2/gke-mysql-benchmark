load("@io_bazel_rules_docker//container:container.bzl", "container_image")

container_image(
    name = "devel",
    base = "@mysql//image:image.tar",
    cmd = ["/entrypoint.sh"],
    files = [
        "entrypoint.sh",
        "schema.sql",
        "data.sql",
    ],
)
