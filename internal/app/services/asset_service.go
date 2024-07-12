package services

import (
	"context"
	"github.com/alkosmas92/platform-go-challenge/internal/app/repository"

	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
)

type AssetService interface {
	GetAssetByID(ctx context.Context, id string) (*models.Asset, error)
	CreateAsset(ctx context.Context, asset *models.Asset) error
	UpdateAsset(ctx context.Context, id string, asset *models.Asset) error
	DeleteAsset(ctx context.Context, id string) error
}

type assetService struct {
	repo repository.AssetRepository
}

func NewAssetService(repo repository.AssetRepository) AssetService {
	return &assetService{repo: repo}
}

func (s *assetService) GetAssetByID(ctx context.Context, id string) (*models.Asset, error) {
	return s.repo.GetAssetByID(ctx, id)
}

func (s *assetService) CreateAsset(ctx context.Context, asset *models.Asset) error {
	return s.repo.CreateAsset(ctx, asset)
}

func (s *assetService) UpdateAsset(ctx context.Context, id string, asset *models.Asset) error {
	return s.repo.UpdateAssetByID(ctx, id, asset)
}
func (s *assetService) DeleteAsset(ctx context.Context, id string) error {
	return s.repo.DeleteAssetByID(ctx, id)
}
