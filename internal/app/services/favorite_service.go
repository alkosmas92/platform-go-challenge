package services

import (
	"context"
	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
	"github.com/alkosmas92/platform-go-challenge/internal/app/repository"
)

type FavoriteService interface {
	GetFavoritesByUserID(ctx context.Context, userID string, limit, offset int) ([]*models.Favorite, error)
	CreateFavorite(ctx context.Context, favorite *models.Favorite) error
	UpdateFavorite(ctx context.Context, userID, assetID string, favorite *models.Favorite) error
	DeleteFavorite(ctx context.Context, userID, assetID string) error
}

type favoriteService struct {
	repo repository.FavoriteRepository
}

func NewFavoriteService(repo repository.FavoriteRepository) FavoriteService {
	return &favoriteService{repo: repo}
}

func (s *favoriteService) GetFavoritesByUserID(ctx context.Context, userID string, limit, offset int) ([]*models.Favorite, error) {
	return s.repo.GetFavoritesByUserID(ctx, userID, limit, offset)
}

func (s *favoriteService) CreateFavorite(ctx context.Context, favorite *models.Favorite) error {
	return s.repo.CreateFavorite(ctx, favorite)
}

func (s *favoriteService) UpdateFavorite(ctx context.Context, userID, assetID string, favorite *models.Favorite) error {
	return s.repo.UpdateFavorite(ctx, userID, assetID, favorite)
}

func (s *favoriteService) DeleteFavorite(ctx context.Context, userID, assetID string) error {
	return s.repo.DeleteFavorite(ctx, userID, assetID)
}
