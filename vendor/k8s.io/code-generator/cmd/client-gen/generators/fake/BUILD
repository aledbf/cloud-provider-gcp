load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "fake_client_generator.go",
        "generator_fake_for_clientset.go",
        "generator_fake_for_group.go",
        "generator_fake_for_type.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/code-generator/cmd/client-gen/generators/fake",
    importpath = "k8s.io/code-generator/cmd/client-gen/generators/fake",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/code-generator/cmd/client-gen/args:go_default_library",
        "//vendor/k8s.io/code-generator/cmd/client-gen/generators/scheme:go_default_library",
        "//vendor/k8s.io/code-generator/cmd/client-gen/generators/util:go_default_library",
        "//vendor/k8s.io/code-generator/cmd/client-gen/path:go_default_library",
        "//vendor/k8s.io/code-generator/cmd/client-gen/types:go_default_library",
        "//vendor/k8s.io/gengo/generator:go_default_library",
        "//vendor/k8s.io/gengo/namer:go_default_library",
        "//vendor/k8s.io/gengo/types:go_default_library",
    ],
)
