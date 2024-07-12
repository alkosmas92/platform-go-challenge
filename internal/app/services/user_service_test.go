package services_test

import (
	"context"
	"testing"

	"github.com/alkosmas92/platform-go-challenge/internal/app/mocks"
	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
	"github.com/alkosmas92/platform-go-challenge/internal/app/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockRepo)
	ctx := context.Background()

	user := models.NewUser("testuser", "password", "Test", "User")

	mockRepo.EXPECT().CreateUser(ctx, user).Return(nil)

	err := userService.RegisterUser(ctx, user)
	assert.NoError(t, err)
}

func TestAuthenticateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockRepo)
	ctx := context.Background()

	user := models.NewUser("testuser", "password", "Test", "User")

	mockRepo.EXPECT().GetUserByUsername(ctx, "testuser").Return(user, nil)

	result, err := userService.AuthenticateUser(ctx, "testuser", "password")
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

func TestAuthenticateUser_InvalidCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockRepo)
	ctx := context.Background()

	user := models.NewUser("testuser", "wrongpassword", "Test", "User")

	mockRepo.EXPECT().GetUserByUsername(ctx, "testuser").Return(user, nil)

	result, err := userService.AuthenticateUser(ctx, "testuser", "password")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "invalid credentials", err.Error())
}
