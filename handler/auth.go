package handler

import (
	"context"
	"pet-pamper/model"
)

type IAuthHandler interface {
	OpenAuthorize(ctx context.Context, request model.OpenAuthorizeRequest) (string, error)
	DeleteRole(ctx context.Context, roleID int64) error
}

func NewAuthHandler() IAuthHandler {
	return &authRepo{}
}

type authRepo struct {
}

func (a *authRepo) OpenAuthorize(ctx context.Context, request model.OpenAuthorizeRequest) (string, error) {
	//request.
	return "Token", nil
}

func (a *authRepo) DeleteRole(ctx context.Context, roleID int64) error {
	return nil
}
