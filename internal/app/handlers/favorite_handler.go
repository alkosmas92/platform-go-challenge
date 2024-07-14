package handlers

import (
	"encoding/json"
	appContext "github.com/alkosmas92/platform-go-challenge/internal/app/context"
	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
	"github.com/alkosmas92/platform-go-challenge/internal/app/services"
	"github.com/sirupsen/logrus"
	"log"
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

	userID, ok := r.Context().Value(appContext.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	favorite, err := handler.Service.GetFavoritesByUserID(r.Context(), userID, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Print("favorite", favorite)
	json.NewEncoder(w).Encode(favorite)
}

func (handler *FavoriteHandler) UpdateFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var favorite models.Favorite
	err := json.NewDecoder(r.Body).Decode(&favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	userID, ok := r.Context().Value(appContext.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	assetID := r.URL.Query().Get("asset_id")
	err = handler.Service.UpdateFavorite(r.Context(), userID, assetID, &favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *FavoriteHandler) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, ok := r.Context().Value(appContext.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	assetID := r.URL.Query().Get("asset_id")

	err := handler.Service.DeleteFavorite(r.Context(), userID, assetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
