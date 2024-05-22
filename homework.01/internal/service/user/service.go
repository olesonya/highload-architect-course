package user

import (
	"github.com/olesonya/highload-architect-course/homework.01/internal/repository"
	def "github.com/olesonya/highload-architect-course/homework.01/internal/service"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

func NewService(
	userRepository repository.UserRepository,
) *service {
	return &service{
		userRepository: userRepository,
	}
}
