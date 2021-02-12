module apps/user_service

go 1.14

replace (
    apps/auth_service => ../auth_service
    common/go => ../../common/go
)

require (
	github.com/golang/protobuf v1.4.3
	apps/auth_service v0.0.0
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)
