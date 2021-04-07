load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "action",
    srcs = [
        "action.go",
    ],
    importpath = "github.com/Jooms/canI/pkg/cmd/action",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_glog//:go_default_library",
        "@com_github_spf13_cobra//:go_default_library",
    ],
)
