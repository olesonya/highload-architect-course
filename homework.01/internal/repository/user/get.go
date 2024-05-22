package user

import (
	"context"

	logger "github.com/sirupsen/logrus"

	"github.com/olesonya/highload-architect-course/homework.01/internal/model"
	"github.com/olesonya/highload-architect-course/homework.01/internal/repository/converter"
	repoModel "github.com/olesonya/highload-architect-course/homework.01/internal/repository/model"
)

func (r *repository) Get(_ context.Context, uuid string) (*model.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	conn := r.db.GetGormDB()

	row := conn.Raw(
		"SELECT id, first_name, second_name, birthdate, biography, city "+
			"FROM postgres.public.\"users\" WHERE id=$1", uuid).Row()

	user := &repoModel.UserDB{
		Info: &repoModel.UserInfoDB{},
	}

	err := row.Scan(
		&user.UserId,
		&user.Info.FirstName,
		&user.Info.SecondName,
		&user.Info.Birthdate,
		&user.Info.Biography,
		&user.Info.City)
	if err != nil {
		logger.Errorf("row.Scan(...): %v\n", err)
		return nil, err
	}

	return converter.ToServiceUserFromRepoUser(user), nil
}
