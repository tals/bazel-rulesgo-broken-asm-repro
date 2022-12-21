load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_docker//container:container.bzl", "container_image")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

# gazelle:prefix github.com/eikemeier/hello-go
# gazelle:build_file_name BUILD,BUILD.bazel
# gazelle:exclude vendor
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/eikemeier/hello-go",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_datadog_zstd//:go_default_library",
        "@com_github_golang_snappy//:go_default_library",
    ],
)

go_binary(
    name = "hello-go",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

container_image(
    name = "base_image",
    base = "@base//image",
)

go_image(
    name = "hello_app",
    base = ":base_image",
    embed = [":go_default_library"],
    goarch = "amd64",
    goos = "linux",
    pure = "on",
    static = "on",
    visibility = ["//visibility:public"],
)
