package services_test

import (
	"context"
	mocks "github.com/alkosmas92/platform-go-challenge/internal/app/mocks"
	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
	"github.com/alkosmas92/platform-go-challenge/internal/app/services"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetFavoritesByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockFavoriteRepository(ctrl)
	service := services.NewFavoriteService(mockRepo)

	favorites := []*models.Favorite{
		{
			UserID:      "1",
			AssetID:     "101",
			AssetType:   "chart",
			Description: "chart asset",
		},
		{
			UserID:      "1",
			AssetID:     "102",
			AssetType:   "insight",
			Description: "insight asset",
		},
	}

	mockRepo.EXPECT().GetFavoritesByUserID(gomock.Any(), "1", 2, 0).Return(favorites, nil)

	result, err := service.GetFavoritesByUserID(context.Background(), "1", 2, 0)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "101", result[0].AssetID)
	assert.Equal(t, "102", result[1].AssetID)
}

func TestCreateFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockFavoriteRepository(ctrl)
	service := services.NewFavoriteService(mockRepo)

	favorite := &models.Favorite{
		UserID:      "2",
		AssetID:     "405",
		AssetType:   "insight",
		Description: "insight asset",
	}

	mockRepo.EXPECT().CreateFavorite(gomock.Any(), favorite).Return(nil)

	err := service.CreateFavorite(context.Background(), favorite)
	assert.NoError(t, err)
}

func TestUpdateFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockFavoriteRepository(ctrl)
	service := services.NewFavoriteService(mockRepo)

	favorite := &models.Favorite{
		UserID:      "2",
		AssetID:     "202",
		AssetType:   "insight",
		Description: "insight asset",
	}

	mockRepo.EXPECT().UpdateFavorite(gomock.Any(), "2", "202", favorite).Return(nil)

	err := service.UpdateFavorite(context.Background(), "2", "202", favorite)
	assert.NoError(t, err)
}

func TestDeleteFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockFavoriteRepository(ctrl)
	service := services.NewFavoriteService(mockRepo)

	mockRepo.EXPECT().DeleteFavorite(gomock.Any(), "2", "202").Return(nil)

	err := service.DeleteFavorite(context.Background(), "2", "202")
	assert.NoError(t, err)
}
