load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "common.go",
        "doc.go",
        "interfaces.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/kube-openapi/pkg/common",
    importpath = "k8s.io/kube-openapi/pkg/common",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/emicklei/go-restful:go_default_library",
        "//vendor/k8s.io/kube-openapi/pkg/openapiconv:go_default_library",
        "//vendor/k8s.io/kube-openapi/pkg/spec3:go_default_library",
        "//vendor/k8s.io/kube-openapi/pkg/validation/spec:go_default_library",
    ],
)
