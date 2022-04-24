package service

import (
	"ddd-practice-go/domain/model"
	"ddd-practice-go/domain/repository"
)

type UserService struct {
	UserRepos repository.UserRepository
}

func (us *UserService) Exists(u *model.User) bool {
	username := u.Name()
	users, err := us.UserRepos.FetchByName(username.ToString())
	if err != nil {
		// TODO: error
	}
	return len(users) > 0
}
