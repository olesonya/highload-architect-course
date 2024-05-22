package converter

import (
	"github.com/olesonya/highload-architect-course/homework.01/internal/model"
	pbUser "github.com/olesonya/highload-architect-course/homework.01/pkg/grpc/user/v1"
)

func ToPbUserFromServiceUser(user *model.User) *pbUser.User {
	return &pbUser.User{
		UserId:   user.UserId,
		UserInfo: ToPbUserInfoFromServiceUserInfo(user.UserInfo),
	}
}

func ToServiceUserFromPbUser(pbUser *pbUser.UserInfo, pbPass string) *model.User {
	return &model.User{
		UserInfo: ToServiceUserInfoFromPbUserInfo(pbUser),
		Password: pbPass,
	}
}

func ToPbUserInfoFromServiceUserInfo(svcInfo *model.UserInfo) *pbUser.UserInfo {
	return &pbUser.UserInfo{
		FirstName:  svcInfo.FirstName,
		SecondName: svcInfo.SecondName,
		Birthdate:  svcInfo.Birthdate,
		Biography:  svcInfo.Biography,
		City:       svcInfo.City,
	}
}

func ToServiceUserInfoFromPbUserInfo(pbInfo *pbUser.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		FirstName:  pbInfo.FirstName,
		SecondName: pbInfo.SecondName,
		Birthdate:  pbInfo.Birthdate,
		Biography:  pbInfo.Biography,
		City:       pbInfo.City,
	}
}
