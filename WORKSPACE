workspace(name = "github_com_eikemeier_hello_go")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "9fb16af4d4836c8222142e54c9efa0bb5fc562ffc893ce2abeac3e25daead144",
    urls = [
        "https://github.com/bazelbuild/rules_go/releases/download/0.19.0/rules_go-0.19.0.tar.gz",
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.19.0/rules_go-0.19.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "3556d4972571f288f8c43378295d84ed64fef5b1a875211ee1046f9f6b4258fa",
    strip_prefix = "rules_docker-0.8.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.8.0.tar.gz"],
)

load("@io_bazel_rules_docker//go:image.bzl", go_image_repositories = "repositories")
load("@io_bazel_rules_docker//container:container.bzl", "container_pull")

go_image_repositories()

# ---

# http://gcr.io/distroless/static:nonroot
container_pull(
    name = "base",
    architecture = "amd64",
    digest = "sha256:06a3a00a6b52b85b5ec325308c921edbf8befd6cb3025b4c02bca3e616caf76e",
    os = "linux",
    registry = "gcr.io",
    repository = "distroless/static",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    tag = "v0.0.1",
)

go_repository(
    name = "com_github_datadog_zstd",
    importpath = "github.com/DataDog/zstd",
    tag = "1.x",
)
