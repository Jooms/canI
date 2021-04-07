load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/Jooms/canI/cmd",
    visibility = ["//visibility:public"],
    deps = ["//pkg/cmd"],
)

go_binary(
    name = "cmd",
    embed = [":go_default_library"],
    importpath = "github.com/Jooms/canI",
    visibility = ["//visibility:public"],
)
