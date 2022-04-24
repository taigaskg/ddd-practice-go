package usecase

import (
	"ddd-practice-go/domain/model"
	"ddd-practice-go/domain/repository"
	"ddd-practice-go/domain/service"
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
	if us.Exists(user) {
		// TODO: error
	}

	err = (*uu.ur).Register(user)
	if err != nil {
		// TODO: error
	}
	return nil
}
