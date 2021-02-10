package main

import (
	"context"
	"errors"

	"github.com/tatskaari/please-codegen-example/apps/auth_service/proto"
	"github.com/tatskaari/please-codegen-example/common/go/auth"
)

type authService struct {

}

func (a *authService) Authenticate(c context.Context, request *proto.AuthenticateRequest) (*proto.AuthenticateResponse, error) {
	if valid := auth.Authenticate(request.Username, request.Password); !valid {
		return nil, errors.New("failed to authenticate")
	}
	return &proto.AuthenticateResponse{
		Token: "asdf",
	}, nil
}

func (a *authService) ValidateToken(c context.Context, request *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	panic("implement me")
}

func main() {

}