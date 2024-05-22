package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	logger "github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"github.com/olesonya/highload-architect-course/homework.01/internal/model"
)

func (s *service) Register(ctx context.Context, info *model.User) (string, error) {
	userUUID, err := uuid.NewUUID()
	if err != nil {
		logger.Errorf("uuid.NewUUID(): %v\n", err)
		return "", err
	}

	info.UserId = userUUID.String()
	hash := safePassword(info.Password)

	err = s.userRepository.Register(ctx, info, hash)
	if err != nil {
		logger.Errorf("s.userRepository.Register(...): %v\n", err)
		return "", err
	}

	return userUUID.String(), nil
}

func safePassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum(nil))
}
