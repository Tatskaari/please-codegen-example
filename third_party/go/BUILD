package(default_visibility = ["PUBLIC"])

go_module(
    name = "protoc-gen-go",
    binary = True,
    install = ["protoc-gen-go"],
    module = "github.com/golang/protobuf",
    version = "v1.4.3",
    deps = [":protobuf-go"],
)

go_module(
    name = "protobuf-go",
    install = ["..."],
    module = "google.golang.org/protobuf",
    version = "v1.25.0",
    deps = [
        ":cmp",
    ],
)

go_module(
    name = "cmp",
    install = [
        "cmp",
        "cmp/...",
    ],
    module = "github.com/google/go-cmp",
    version = "b1c9c4891a6525d98001fea424c8926c6d77bb56",
)
