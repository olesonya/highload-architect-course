package converter

import (
	"github.com/olesonya/highload-architect-course/homework.01/internal/model"
	repoModel "github.com/olesonya/highload-architect-course/homework.01/internal/repository/model"
)

func ToServiceUserFromRepoUser(user *repoModel.UserDB) *model.User {
	return &model.User{
		UserId:   user.UserId,
		UserInfo: ToServiceUserInfoFromRepoUserInfo(user.Info),
	}
}

func ToServiceUserInfoFromRepoUserInfo(info *repoModel.UserInfoDB) *model.UserInfo {
	return &model.UserInfo{
		FirstName:  info.FirstName,
		SecondName: info.SecondName,
		Birthdate:  info.Birthdate,
		Biography:  info.Biography,
		City:       info.City,
	}
}

func ToRepoUserInfoFromServiceUserInfo(info *model.UserInfo) *repoModel.UserInfoDB {
	return &repoModel.UserInfoDB{
		FirstName:  info.FirstName,
		SecondName: info.SecondName,
		Birthdate:  info.Birthdate,
		Biography:  info.Biography,
		City:       info.City,
	}
}
