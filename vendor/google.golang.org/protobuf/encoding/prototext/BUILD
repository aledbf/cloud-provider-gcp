load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "decode.go",
        "doc.go",
        "encode.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/google.golang.org/protobuf/encoding/prototext",
    importpath = "google.golang.org/protobuf/encoding/prototext",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/google.golang.org/protobuf/encoding/protowire:go_default_library",
        "//vendor/google.golang.org/protobuf/internal/encoding/messageset:go_default_library",
        "//vendor/google.golang.org/protobuf/internal/encoding/text:go_default_library",
        "//vendor/google.golang.org/protobuf/internal/errors:go_default_library",
        "//vendor/google.golang.org/protobuf/internal/flags:go_default_library",
        "//vendor/google.golang.org/protobuf/internal/genid:go_default_library",
        "//vendor/google.golang.org/protobuf/internal/order:go_default_library",
        "//vendor/google.golang.org/protobuf/internal/pragma:go_default_library",
        "//vendor/google.golang.org/protobuf/internal/set:go_default_library",
        "//vendor/google.golang.org/protobuf/internal/strs:go_default_library",
        "//vendor/google.golang.org/protobuf/proto:go_default_library",
        "//vendor/google.golang.org/protobuf/reflect/protoreflect:go_default_library",
        "//vendor/google.golang.org/protobuf/reflect/protoregistry:go_default_library",
    ],
)
