module github.com/tatskaari/please-codegen-example/apps/auth_service

go 1.14

replace github.com/tatskaari/please-codegen-example/common/go => ../../common/go

require (
	github.com/golang/protobuf v1.4.3
	github.com/tatskaari/please-codegen-example/common/go v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)
