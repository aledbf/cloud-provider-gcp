load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/sigs.k8s.io/controller-tools/cmd/controller-gen",
    importpath = "sigs.k8s.io/controller-tools/cmd/controller-gen",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/sigs.k8s.io/controller-tools/pkg/crd:go_default_library",
        "//vendor/sigs.k8s.io/controller-tools/pkg/deepcopy:go_default_library",
        "//vendor/sigs.k8s.io/controller-tools/pkg/genall:go_default_library",
        "//vendor/sigs.k8s.io/controller-tools/pkg/genall/help:go_default_library",
        "//vendor/sigs.k8s.io/controller-tools/pkg/genall/help/pretty:go_default_library",
        "//vendor/sigs.k8s.io/controller-tools/pkg/markers:go_default_library",
        "//vendor/sigs.k8s.io/controller-tools/pkg/rbac:go_default_library",
        "//vendor/sigs.k8s.io/controller-tools/pkg/schemapatcher:go_default_library",
        "//vendor/sigs.k8s.io/controller-tools/pkg/version:go_default_library",
        "//vendor/sigs.k8s.io/controller-tools/pkg/webhook:go_default_library",
    ],
)

go_binary(
    name = "controller-gen",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
