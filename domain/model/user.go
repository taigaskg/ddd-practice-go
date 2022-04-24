package model

import (
	"errors"
	"unicode/utf8"

	"github.com/google/uuid"
)

type User struct {
	id   UserID
	name UserName
}

func NewUser(id UserID, name UserName) (*User, error) {
	if name == (UserName{}) {
		return nil, errors.New("name must not be empty struct")
	}

	if id == (UserID{}) {
		return &User{id: UserID{uuid.NewString()}, name: name}, nil
	}
	return &User{id: id, name: name}, nil
}

func (u *User) SetName(name UserName) error {
	if name == (UserName{}) {
		return errors.New("name must not be empty struct")
	}
	u.name = name
	return nil
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Name() UserName {
	return u.name
}

func (u *User) Equals(other User) bool {
	return u.id == other.id
}

type UserID struct {
	value string
}

func NewUserID(value string) (*UserID, error) {
	if value == "" {
		return nil, errors.New("value must not be empty")
	}
	return &UserID{value}, nil
}

func (ui *UserID) ToString() string {
	return ui.value
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

func (un *UserName) ToString() string {
	return un.value
}
