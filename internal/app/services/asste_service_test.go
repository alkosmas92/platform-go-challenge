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

func TestGetAssetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAssetRepository(ctrl)
	service := services.NewAssetService(mockRepo)

	asset := &models.Asset{
		ID:          "1",
		Type:        "chart",
		Description: "Chart asset",
		Chart:       &models.Chart{Title: "Test Chart"},
	}
	mockRepo.EXPECT().GetAssetByID(gomock.Any(), "1").
		Return(asset, nil)

	result, err := service.GetAssetByID(context.Background(), "1")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, asset, result)
}

func TestCreateAsset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAssetRepository(ctrl)
	service := services.NewAssetService(mockRepo)

	asset := &models.Asset{
		ID:          "2",
		Type:        "insight",
		Description: "Test insight asset",
		Insight:     &models.Insight{Text: "Insight Text"},
	}

	mockRepo.EXPECT().CreateAsset(gomock.Any(), asset).Return(nil)

	err := service.CreateAsset(context.Background(), asset)
	assert.NoError(t, err)
}

func TestUpdateAsset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockAssetRepository(ctrl)
	service := services.NewAssetService(mockRepo)
	asset := &models.Asset{
		ID:          "3",
		Type:        "audience",
		Description: "Test audience asset",
		Audience:    &models.Audience{Gender: "Female", AgeGroup: "18-34"},
	}
	mockRepo.EXPECT().UpdateAssetByID(gomock.Any(), "3", asset).Return(nil)
	err := service.UpdateAsset(context.Background(), "3", asset)
	assert.NoError(t, err)
}

func TestDeleteAsset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockAssetRepository(ctrl)
	service := services.NewAssetService(mockRepo)

	mockRepo.EXPECT().DeleteAssetByID(gomock.Any(), "3").Return(nil)
	err := service.DeleteAsset(context.Background(), "3")
	assert.NoError(t, err)
}
