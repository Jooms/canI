load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cmd",
    srcs = [
        "root.go",
    ],
    importpath = "github.com/Jooms/canI/pkg/cmd",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/cmd/action",
        "@com_github_golang_glog//:go_default_library",
        "@com_github_spf13_cobra//:go_default_library",
    ],
)
