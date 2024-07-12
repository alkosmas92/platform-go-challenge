package handlers

import (
	"encoding/json"
	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
	"github.com/alkosmas92/platform-go-challenge/internal/app/services"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type FavoriteHandler struct {
	Service services.FavoriteService
	Logger  *logrus.Logger
}

func NewFavoriteHandler(service services.FavoriteService, logger *logrus.Logger) *FavoriteHandler {
	return &FavoriteHandler{Service: service, Logger: logger}
}

func (handler *FavoriteHandler) CreateFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var favorite models.Favorite
	err := json.NewDecoder(r.Body).Decode(&favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.Service.CreateFavorite(r.Context(), &favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (handler *FavoriteHandler) GetFavoritesByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.URL.Query().Get("user_id")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10 // default limit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0 // default offset
	}

	favorite, err := handler.Service.GetFavoritesByUserID(r.Context(), userID, offset, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(favorite)
}

func (handler *FavoriteHandler) UpdateFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var favorite models.Favorite
	err := json.NewDecoder(r.Body).Decode(&favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	userID := r.URL.Query().Get("user_id")
	assetID := r.URL.Query().Get("asset_id")
	err = handler.Service.UpdateFavorite(r.Context(), userID, assetID, &favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *FavoriteHandler) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.URL.Query().Get("user_id")
	assetID := r.URL.Query().Get("asset_id")

	err := handler.Service.DeleteFavorite(r.Context(), userID, assetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
