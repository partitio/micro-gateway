load("@bazel_gazelle//:def.bzl", "gazelle")
load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")

buildifier(
    name = "buildifier",
)

buildifier(
    name = "buildifier_check",
    mode = "check",
)

# gazelle:exclude third_party
# gazelle:exclude vendor

gazelle(
    name = "gazelle_diff",
    mode = "diff",
    prefix = "github.com/partitio/micro-gateway",
)

gazelle(
    name = "gazelle_fix",
    mode = "fix",
    prefix = "github.com/partitio/micro-gateway",
)

package_group(
    name = "generators",
    packages = [
        "//protoc-gen-micro-gateway/...",
        "//protoc-gen-swagger/...",
    ],
)
