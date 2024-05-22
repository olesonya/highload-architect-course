package model

type UserDB struct {
	UserId string
	Hash   string
	Info   *UserInfoDB
}

type UserInfoDB struct {
	FirstName  string
	SecondName string
	Birthdate  string
	Biography  string
	City       string
}
