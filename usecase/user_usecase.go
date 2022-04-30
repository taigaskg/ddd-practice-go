package usecase

import (
	"ddd-practice-go/domain/model"
	"ddd-practice-go/domain/repository"
	"ddd-practice-go/domain/service"
	"errors"
)

type UserUsecase interface {
	CreateUser(name string) error
}

type userUsecase struct {
	ur *repository.UserRepository
}

func NewUserUsecase(ur *repository.UserRepository) *userUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) CreateUser(userName string) error {

	name, _ := model.NewUserName(userName)
	user, err := model.NewUser(model.UserID{}, *name)
	if err != nil {
		// TODO: error
	}
	us := service.UserService{UserRepos: *uu.ur}
	if res, err := us.Exists(user); res {
		if err != nil {
			return err
		}
		return errors.New("duplicate user name")
	}

	err = (*uu.ur).Register(user)
	if err != nil {
		// TODO: error
	}
	return nil
}
