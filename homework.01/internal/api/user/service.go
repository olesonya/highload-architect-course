package user

import (
	"github.com/olesonya/highload-architect-course/homework.01/internal/service"
	pbUser "github.com/olesonya/highload-architect-course/homework.01/pkg/grpc/user/v1"
)

type Instance struct {
	pbUser.UnimplementedUserServiceServer
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Instance {
	return &Instance{
		userService: userService,
	}
}
