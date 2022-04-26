package service

import (
	"ddd-practice-go/domain/model"
	"ddd-practice-go/domain/repository"
	"errors"
)

type UserService struct {
	UserRepos repository.UserRepository
}

func (us *UserService) Exists(u *model.User) (bool, error) {
	username := u.Name()
	users, err := us.UserRepos.FetchByName(username.ToString())
	if err != nil {
		return false, errors.New("failure fetch from repository")
	}
	return len(users) > 0, nil
}
