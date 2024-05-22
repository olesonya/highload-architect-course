package model

type User struct {
	UserId   string
	Password string
	UserInfo *UserInfo
}

type UserInfo struct {
	FirstName  string
	SecondName string
	Birthdate  string
	Biography  string
	City       string
}
