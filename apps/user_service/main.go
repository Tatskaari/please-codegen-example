package main

import (
	"context"

	authproto "github.com/tatskaari/please-codegen-example/apps/auth_service/proto"
	userproto "github.com/tatskaari/please-codegen-example/apps/user_service/proto"
)

type userService struct {
	authService authproto.AuthServiceClient
}

func (u *userService) CreateUser(ctx context.Context, request *userproto.CreateUserRequest) (*userproto.User, error) {
	_, err := u.authService.ValidateToken(ctx, &authproto.ValidateTokenRequest{
		RequestId: request.RequestId,
		Token:     request.AuthToken,
	})

	/**
	 * Imagine some code here to create the user in the DB
	 */

	return request.User, err
}

func main() {
	panic("this is a stub")
}