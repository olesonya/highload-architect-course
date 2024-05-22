package user

import (
	"context"

	pbUser "github.com/olesonya/highload-architect-course/homework.01/pkg/grpc/user/v1"
)

func (s *Instance) Login(ctx context.Context, req *pbUser.LoginRequest) (*pbUser.LoginResponse, error) {
	token, err := s.userService.Login(ctx, req.GetUserId(), req.GetUserPass())
	if err != nil {
		return nil, err
	}

	return &pbUser.LoginResponse{
		Token: token,
	}, nil
}
