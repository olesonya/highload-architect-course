package user

import (
	"context"

	logger "github.com/sirupsen/logrus"

	"github.com/olesonya/highload-architect-course/homework.01/internal/model"
)

func (s *service) Get(ctx context.Context, uuid string) (*model.User, error) {
	user, err := s.userRepository.Get(ctx, uuid)
	if err != nil {
		logger.Errorf("s.userRepository.Get(...): %v\n", err)
		return nil, err
	}
	if user == nil {
		logger.Errorf("user with uuid %s not found\n", uuid)
		return nil, model.ErrorUserNotFound
	}

	return user, nil
}
