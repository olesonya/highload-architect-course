package user

import (
	"context"

	logger "github.com/sirupsen/logrus"

	"github.com/olesonya/highload-architect-course/homework.01/internal/model"
	"github.com/olesonya/highload-architect-course/homework.01/internal/repository/converter"
	repoModel "github.com/olesonya/highload-architect-course/homework.01/internal/repository/model"
)

func (r *repository) Register(_ context.Context, info *model.User, hash string) error {
	r.m.Lock()
	defer r.m.Unlock()

	user := &repoModel.UserDB{
		UserId: info.UserId,
		Hash:   hash,
		Info:   converter.ToRepoUserInfoFromServiceUserInfo(info.UserInfo),
	}

	conn := r.db.GetGormDB()

	tx := conn.Exec(
		"INSERT INTO postgres.public.\"users\" VALUES ($1,$2,$3,$4,$5,$6,$7)",
		user.UserId,
		user.Hash,
		user.Info.FirstName,
		user.Info.SecondName,
		user.Info.Birthdate,
		user.Info.Biography,
		user.Info.City,
	)

	if tx.Error != nil {
		logger.Errorf("conn.Exec(...): %v\n", tx.Error)
		return tx.Error
	}

	return nil
}
