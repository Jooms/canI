load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_binary")


# gazelle:prefix github.com/Jooms/canI
gazelle(name = "gazelle")

go_binary(
    name = "canI",
    embed = ["//cmd:go_default_library"],
    importpath = "github.com/Jooms/canI",
    visibility = ["//visibility:public"],
)
