package service

import (
	"ddd-practice-go/domain/model"
	"ddd-practice-go/domain/repository/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestExists(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	us := UserService{mockUserRepository}

	userID, err := model.NewUserID("id")
	require.NoError(t, err)
	userName, err := model.NewUserName("name")
	require.NoError(t, err)
	user, err := model.NewUser(*userID, *userName)
	require.NoError(t, err)

	t.Run("duplicate(exist)", func(t *testing.T) {
		mockUserRepository.On("FetchByName", mock.AnythingOfType("string")).Return([]*model.User{{}}, nil).Once()
		got, gotErr := us.Exists(user)
		assert.NoError(t, gotErr)
		assert.Equal(t, true, got)
	})

	t.Run("not exits", func(t *testing.T) {
		mockUserRepository.On("FetchByName", mock.AnythingOfType("string")).Return([]*model.User{}, nil).Once()
		got, gotErr := us.Exists(user)
		assert.NoError(t, gotErr)
		assert.Equal(t, false, got)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("FetchByName", mock.AnythingOfType("string")).Return(nil, errors.New("")).Once()
		got, gotErr := us.Exists(user)
		assert.Error(t, gotErr)
		assert.Equal(t, false, got)
	})
}
