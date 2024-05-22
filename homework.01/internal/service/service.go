package service

import (
	"context"

	"github.com/olesonya/highload-architect-course/homework.01/internal/model"
)

type UserService interface {
	Register(ctx context.Context, info *model.User) (string, error)
	Get(ctx context.Context, uuid string) (*model.User, error)
	Login(ctx context.Context, uuid string, pswd string) (string, error)
}
