load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "factory.go",
        "factoryinterface.go",
        "generic.go",
        "groupinterface.go",
        "informer.go",
        "packages.go",
        "types.go",
        "versioninterface.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/code-generator/cmd/informer-gen/generators",
    importpath = "k8s.io/code-generator/cmd/informer-gen/generators",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/code-generator/cmd/client-gen/generators/util:go_default_library",
        "//vendor/k8s.io/code-generator/cmd/client-gen/types:go_default_library",
        "//vendor/k8s.io/code-generator/cmd/informer-gen/args:go_default_library",
        "//vendor/k8s.io/code-generator/pkg/namer:go_default_library",
        "//vendor/k8s.io/code-generator/pkg/util:go_default_library",
        "//vendor/k8s.io/gengo/args:go_default_library",
        "//vendor/k8s.io/gengo/generator:go_default_library",
        "//vendor/k8s.io/gengo/namer:go_default_library",
        "//vendor/k8s.io/gengo/types:go_default_library",
        "//vendor/k8s.io/klog/v2:go_default_library",
    ],
)
