package user

import (
	"context"

	logger "github.com/sirupsen/logrus"
)

func (s *service) Login(ctx context.Context, uuid, pswd string) (string, error) {
	hash := safePassword(pswd)

	if err := s.userRepository.Login(ctx, uuid, hash); err != nil {
		logger.Errorf("s.userRepository.Login(...): %v\n", err)
		return "", err
	}

	return hash, nil
}
