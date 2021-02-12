go_binary(
    name = "auth_service",
    srcs = ["main.go"],
    deps = [
        "//apps/auth_service/proto",
        "//common/go/auth",
    ],
)

pkg = package_name()
sh_cmd(
    name = "update_gitignore",
    cmd = f"plz generate --update_gitignore={pkg}/.gitignore",
)