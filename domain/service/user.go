package service

import (
	"ddd-practice-go/domain/model"
	"ddd-practice-go/domain/repository"
	"log"
)

type UserService struct {
	UserRepos repository.UserRepository
}

func (us *UserService) Exists(u *model.User) bool {
	username := u.Name() // TODO: Go wayな書き方ではgetterにGetをつけない
	users, err := us.UserRepos.FetchByName(username.ToString())
	if err != nil {
		log.Fatal(err)
	}
	return len(users) > 0
}
