package main

import (
	"context"

	authproto "apps/auth_service/proto"
	userproto "apps/user_service/proto"
)

type DAO interface {
	StoreUser(string, string) error
}

type userService struct {
	authService authproto.AuthServiceClient

	dao DAO
}

func (u *userService) CreateUser(ctx context.Context, request *userproto.CreateUserRequest) (*userproto.User, error) {
	_, err := u.authService.ValidateToken(ctx, &authproto.ValidateTokenRequest{
		RequestId: request.RequestId,
		Token:     request.AuthToken,
	})

	if err := u.dao.StoreUser(request.User.Username, request.User.Email); err != nil {
		return nil, err
	}

	return request.User, err
}

func main() {
	println("Success! User service compiled and ran!")
}