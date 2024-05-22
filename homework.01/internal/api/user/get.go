package user

import (
	"context"

	"github.com/olesonya/highload-architect-course/homework.01/internal/service/converter"
	pbUser "github.com/olesonya/highload-architect-course/homework.01/pkg/grpc/user/v1"
)

func (s *Instance) Get(ctx context.Context, req *pbUser.GetRequest) (*pbUser.GetResponse, error) {
	user, err := s.userService.Get(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &pbUser.GetResponse{
		User: converter.ToPbUserFromServiceUser(user),
	}, nil
}
