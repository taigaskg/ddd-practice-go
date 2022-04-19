package user

import (
	"errors"
	"unicode/utf8"
)

type User struct {
	id   UserId
	name UserName
}

func NewUser(id UserId, name UserName) *User {
	return &User{id: id, name: name}
}

func (u *User) SetName(name UserName) {
	u.name = name
}

func (u *User) GetId() UserId {
	return u.id
}

func (u *User) GetName() UserName {
	return u.name
}

type UserId struct {
	value string
}

func NewUserId(value string) (*UserId, error) {
	if value == "" {
		return nil, errors.New("value must not be empty")
	}
	return &UserId{value}, nil
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
