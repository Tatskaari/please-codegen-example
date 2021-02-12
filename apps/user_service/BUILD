go_binary(
    name = "user_service",
    srcs = ["main.go"],
    deps = [
        "//apps/auth_service/proto",
        "//apps/user_service/proto",
    ],
)

pkg = package_name()
sh_cmd(
    name = "update_gitignore",
    cmd = f"plz generate --update_gitignore={pkg}/.gitignore",
)