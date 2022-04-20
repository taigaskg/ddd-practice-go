package user

import (
	"errors"
	"unicode/utf8"
)

type User struct {
	id   UserId
	name UserName
}

func NewUser(id UserId, name UserName) (*User, error) {
	if id == (UserId{}) {
		return nil, errors.New("id must not be empty struct")
	}
	if name == (UserName{}) {
		return nil, errors.New("name must not be empty struct")
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

func (u *User) GetId() UserId {
	return u.id
}

func (u *User) GetName() UserName {
	return u.name
}

func (u *User) Equals(other User) bool {
	return u.id == other.id
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

//  ドメインサービスはインスタンス特有の状態をもたない。
// JavaなどのOOPの場合、UserServiceなどのドメインサービスのクラスを作るだろう。
// Goの場合はOOPではないので、クラスを作る代わりに構造体を定義するのは違和感がある。
// 同じパッケージの関数として定義するか別パッケージの関数にするか？どちらが良いのか？TODO:
func Exists(u *User) bool {
	// TODO:
	return false
}
