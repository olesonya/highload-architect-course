package user

import (
	"context"

	logger "github.com/sirupsen/logrus"

	repoModel "github.com/olesonya/highload-architect-course/homework.01/internal/repository/model"
)

func (r *repository) Login(_ context.Context, uuid, hash string) error {
	r.m.RLock()
	defer r.m.RUnlock()

	conn := r.db.GetGormDB()
	row := conn.Raw(
		"SELECT id, hash FROM postgres.public.\"users\" WHERE id=$1 AND hash=$2", uuid, hash).Row()

	user := &repoModel.UserDB{
		Info: &repoModel.UserInfoDB{},
	}

	err := row.Scan(
		&user.UserId,
		&user.Hash,
	)
	if err != nil {
		logger.Errorf("row.Scan(...): %v\n", err)
		return err
	}

	return nil
}
