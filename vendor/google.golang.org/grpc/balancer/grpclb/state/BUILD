load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["state.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/google.golang.org/grpc/balancer/grpclb/state",
    importpath = "google.golang.org/grpc/balancer/grpclb/state",
    visibility = ["//visibility:public"],
    deps = ["//vendor/google.golang.org/grpc/resolver:go_default_library"],
)
