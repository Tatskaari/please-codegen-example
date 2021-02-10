subinclude("//build/defs:gitignore")

go_binary(
    name = "auth_service",
    srcs = ["main.go"],
    deps = [
        "//apps/auth_service/proto",
        "//common/go/auth",
    ],
)

update_gitignore(name = "update_gitignore")
