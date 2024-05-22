package user

import (
	"context"

	"github.com/olesonya/highload-architect-course/homework.01/internal/service/converter"
	pbUser "github.com/olesonya/highload-architect-course/homework.01/pkg/grpc/user/v1"
)

func (s *Instance) Register(ctx context.Context, req *pbUser.RegisterRequest) (*pbUser.RegisterResponse, error) {
	uuid, err := s.userService.Register(ctx, converter.ToServiceUserFromPbUser(req.UserInfo, req.UserPass))
	if err != nil {
		return nil, err
	}

	return &pbUser.RegisterResponse{
		UserId: uuid,
	}, nil
}
