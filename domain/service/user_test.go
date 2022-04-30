package service

import (
	"ddd-practice-go/domain/model"
	mockss "ddd-practice-go/domain/repository/mock_repository"
	"ddd-practice-go/domain/repository/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
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

func TestExists2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mockss.NewMockUserRepository(ctrl)
	us := UserService{m}

	userID, err := model.NewUserID("id")
	require.NoError(t, err)
	userName, err := model.NewUserName("name")
	require.NoError(t, err)
	user, err := model.NewUser(*userID, *userName)
	require.NoError(t, err)

	t.Run("duplicate(exist)", func(t *testing.T) {
		m.EXPECT().FetchByName(gomock.Eq("name")).Return([]*model.User{{}}, nil)

		got, gotErr := us.Exists(user)
		assert.Equal(t, true, got)
		assert.NoError(t, gotErr)
	})

	t.Run("not exists", func(t *testing.T) {
		m.EXPECT().FetchByName(gomock.Eq("name")).Return([]*model.User{}, nil)

		got, gotErr := us.Exists(user)
		assert.Equal(t, false, got)
		assert.NoError(t, gotErr)
	})

	t.Run("error", func(t *testing.T) {
		m.EXPECT().FetchByName(gomock.Eq("name")).Return(nil, errors.New("unexpected error"))

		got, gotErr := us.Exists(user)
		assert.Equal(t, false, got)
		assert.Error(t, gotErr)
	})
}
