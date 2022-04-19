package user

import (
	"errors"
	"unicode/utf8"
)

type User struct {
	Id   UserId
	Name UserName
}

func NewUser(id UserId, name UserName) *User {
	return &User{Id: id, Name: name}
}

// TODO: このようなメソッドを作るのはGoらしいのか？直接プロパティ変更も可能だが。いらないのでは？
// func (u *User) ChangeName(name UserName) {
// 	u.Name = name
// }

type UserId string

func NewUserId(value string) (*UserId, error) {
	if value == "" {
		return nil, errors.New("value must not be empty")
	}
	u := UserId(value)
	return &u, nil
}

type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	if value == "" {
		return nil, errors.New("value must not be empty")
	}
	if utf8.RuneCountInString(value) < 3 {
		return nil, errors.New("value must be more than 3 characters.")
	}
	return &UserName{value}, nil
}
