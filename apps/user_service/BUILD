subinclude("//build/defs:gitignore")

go_binary(
    name = "user_service",
    srcs = ["main.go"],
    deps = [
        "//apps/auth_service/proto",
    ],
)

update_gitignore(name = "update_gitignore")
