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

go_mod_download(
    name = "protobuf_download",
    module = "github.com/golang/protobuf",
    version = "v1.4.3",
)

go_module(
    name = "protobuf",
    download = ":protobuf_download",
    exported_deps = [
        ":protobuf-go",
    ],
    install = ["..."],
    module = "github.com/golang/protobuf",
    strip = [
        "proto/proto3_proto",
        "conformance",
    ],
)

go_module(
    name = "cloud-metadata",
    install = ["compute/metadata"],
    module = "cloud.google.com/go",
    version = "v0.40.0",
)

go_mod_download(
    name = "genproto_download",
    module = "google.golang.org/genproto",
    version = "2b5a72b8730b0b16380010cfe5286c42108d88e7",
)

go_module(
    name = "net",
    install = ["..."],
    module = "golang.org/x/net",
    version = "136a25c244d3019482a795d728110278d6ba09a4",
    deps = [
        ":text",
        ":xcrypto",
    ],
)

go_module(
    name = "xcrypto",
    install = [
        "ssh/terminal",
        "cast5",

        # TODO(jpoole): consider if requiring this to be explicit is a good idea or if we should do the
        # `openpgp/openpgp.a` "hack" here
        "openpgp/...",
        "openpgp",
    ],
    module = "golang.org/x/crypto",
    version = "7b85b097bf7527677d54d3220065e966a0e3b613",
)

go_module(
    name = "tools",
    install = [
        "cover",
        "go/loader",
        "go/buildutil",
        "go/ast/astutil",
    ],
    module = "golang.org/x/tools",
    version = "2ae76fd1560b622911f444c1e66b70a857e1de67",
)

go_module(
    name = "text",
    install = [
        "secure/...",
        "unicode/...",
        "transform",
        "encoding",
        "encoding/...",
    ],
    module = "golang.org/x/text",
    version = "4e4a3210bb54bb31f6ab2cdca2edcc0b50c420c1",
    deps = [
        ":tools",
    ],
)

go_module(
    name = "grpc",
    exported_deps = [
        ":net",
        ":protobuf-go",
    ],
    install = [
        ".",
        "attributes",
        "backoff",
        "balancer",
        "balancer/base",
        "balancer/grpclb/...",
        "balancer/grpclb",
        "balancer/roundrobin",
        "binarylog/...",
        "codes",
        "connectivity",
        "credentials",
        "credentials/oauth",
        "encoding",
        "encoding/gzip",
        "encoding/proto",
        "grpclog",
        "health",
        "health/grpc_health_v1",
        "internal/...",
        "internal",
        "keepalive",
        "metadata",
        "peer",
        "resolver",
        "resolver/dns",
        "resolver/passthrough",
        "serviceconfig",
        "stats",
        "status",
        "tap",
    ],
    module = "google.golang.org/grpc",
    version = "v1.34.0",
    deps = [
        ":cmp",
        ":genproto_rpc",
        ":oauth2",
        ":protobuf",
        ":uuid",
        ":xsys",
    ],
)

# genproto stuff for grpc to resolve cyclic deps
go_module(
    name = "genproto_rpc",
    download = ":genproto_download",
    install = [
        "googleapis/rpc/...",
    ],
    module = "google.golang.org/genproto",
    deps = [
        ":protobuf",
    ],
)

go_module(
    name = "uuid",
    module = "github.com/google/uuid",
    version = "v1.1.1",
)

go_module(
    name = "xsys",
    install = ["..."],
    module = "golang.org/x/sys",
    version = "765f4ea38db36397e827c4153018aa272eed7835",
)

go_module(
    name = "oauth2",
    install = [
        ".",
        "google",
        "internal",
        "jws",
        "jwt",
    ],
    module = "golang.org/x/oauth2",
    version = "d668ce993890a79bda886613ee587a69dd5da7a6",
    visibility = ["//third_party/go/..."],
    deps = [
        ":cloud-metadata",
        "//third_party/go:net",
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
