package repository

import "ddd-practice-go/domain/model"

type UserRepository interface {
	FetchByName(string) ([]*model.User, error)
	Register(*model.User) error
}
