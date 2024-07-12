package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/alkosmas92/platform-go-challenge/internal/app/handlers"
	mocks "github.com/alkosmas92/platform-go-challenge/internal/app/mocks"
	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFavoriteByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFavoriteService(ctrl)
	logger := logrus.New()
	handler := handlers.NewFavoriteHandler(mockService, logger)

	// Test data
	userID := "1"
	limit := 5
	offset := 0

	favorites := []*models.Favorite{
		{UserID: userID, AssetID: "101", AssetType: "chart", Description: "Favorite chart asset"},
	}

	favoritesJSON, _ := json.Marshal(favorites)

	mockService.EXPECT().GetFavoritesByUserID(gomock.Any(), userID, offset, limit).Return(favorites, nil)

	req, _ := http.NewRequest(http.MethodGet, "/favorites?user_id="+userID+"&limit=5&offset=0", nil)
	rr := httptest.NewRecorder()

	handler.GetFavoritesByUserID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, string(favoritesJSON), rr.Body.String())
}

func TestCreateFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFavoriteService(ctrl)
	logger := logrus.New()
	handler := handlers.NewFavoriteHandler(mockService, logger)

	// Test data
	favorite := &models.Favorite{
		UserID:      "1",
		AssetID:     "101",
		AssetType:   "chart",
		Description: "Favorite chart asset",
	}

	favoriteJSON, _ := json.Marshal(favorite)

	mockService.EXPECT().CreateFavorite(gomock.Any(), favorite).Return(nil)

	req, _ := http.NewRequest(http.MethodPost, "/favorites", bytes.NewBuffer(favoriteJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateFavorite(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUpdateFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFavoriteService(ctrl)
	logger := logrus.New()
	handler := handlers.NewFavoriteHandler(mockService, logger)

	// Test data
	favorite := &models.Favorite{
		UserID:      "1",
		AssetID:     "101",
		AssetType:   "insight",
		Description: "Updated favorite insight asset",
	}

	favoriteJSON, _ := json.Marshal(favorite)

	mockService.EXPECT().UpdateFavorite(gomock.Any(), favorite.UserID, favorite.AssetID, favorite).Return(nil)

	req, _ := http.NewRequest(http.MethodPut, "/favorites?user_id=1&asset_id=101", bytes.NewBuffer(favoriteJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.UpdateFavorite(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFavoriteService(ctrl)
	logger := logrus.New()
	handler := handlers.NewFavoriteHandler(mockService, logger)

	// Test data
	userID := "1"
	assetID := "101"

	mockService.EXPECT().DeleteFavorite(gomock.Any(), userID, assetID).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/favorites?user_id=1&asset_id=101", nil)
	rr := httptest.NewRecorder()

	handler.DeleteFavorite(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
