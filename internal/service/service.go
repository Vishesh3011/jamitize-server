package service

import (
	"context"
	"example/internal/core/application"
)

type Service interface {
	User() UserService
}

type service struct {
	userSvc UserService
}

func (s service) User() UserService {
	return s.userSvc
}

func NewService(ctx context.Context, app application.Application) Service {
	return &service{
		userSvc: NewUserService(ctx, app.Logger(), app.DB(), app.Config().JwtConfig().Secret()),
	}
}
