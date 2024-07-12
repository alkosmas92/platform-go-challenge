package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
)

type AssetRepository interface {
	GetAssetByID(ctx context.Context, assetID string) (*models.Asset, error)
	CreateAsset(ctx context.Context, asset *models.Asset) error
	UpdateAssetByID(ctx context.Context, id string, asset *models.Asset) error
	DeleteAssetByID(ctx context.Context, assetID string) error
}

type assetRepository struct {
	db *sql.DB
}

func NewAssetRepository(db *sql.DB) AssetRepository {
	return &assetRepository{db: db}
}

func (repo *assetRepository) GetAssetByID(ctx context.Context, id string) (*models.Asset, error) {
	var asset models.Asset
	var audienceData, chartData, insightData []byte

	query := `
        SELECT id, type, description, audience, chart, insight 
        FROM assets 
        WHERE id = ?`
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&asset.ID, &asset.Type, &asset.Description, &audienceData, &chartData, &insightData)
	if err != nil {
		return nil, err
	}

	switch asset.Type {
	case "audience":
		var audience models.Audience
		if err := json.Unmarshal(audienceData, &audience); err != nil {
			return nil, err
		}
		asset.Audience = &audience
	case "chart":
		var chart models.Chart
		if err := json.Unmarshal(chartData, &chart); err != nil {
			return nil, err
		}
		asset.Chart = &chart
	case "insight":
		var insight models.Insight
		if err := json.Unmarshal(insightData, &insight); err != nil {
			return nil, err
		}
		asset.Insight = &insight
	default:
		return nil, errors.New("invalid asset type")
	}

	return &asset, nil
}

func (repo *assetRepository) CreateAsset(ctx context.Context, asset *models.Asset) error {
	var audienceData, chartData, InsightData []byte
	var err error

	switch asset.Type {
	case "audience":
		audienceData, err = json.Marshal(asset.Audience)
	case "chart":
		chartData, err = json.Marshal(asset.Chart)
	case "insight":
		InsightData, err = json.Marshal(asset.Insight)
	default:
		return errors.New("unknown asset type")
	}
	if err != nil {
		return err
	}
	query := `
		INSERT INTO assets (id, type, description, audience, chart, insight)
		VALUES (?, ?, ?, ?, ?,?)
	`
	_, err = repo.db.ExecContext(ctx, query, asset.ID, asset.Type, asset.Description, audienceData, chartData, InsightData)
	return err
}

func (repo *assetRepository) UpdateAssetByID(ctx context.Context, id string, asset *models.Asset) error {
	var audienceData, chartData, InsightData []byte
	var err error
	switch asset.Type {
	case "audience":
		audienceData, err = json.Marshal(asset.Audience)
	case "chart":
		chartData, err = json.Marshal(asset.Chart)
	case "insight":
		InsightData, err = json.Marshal(asset.Insight)
	default:
		err = errors.New("unknown asset type")
	}
	if err != nil {
		return err
	}

	query := `
		UPDATE assets
		SET type = ?, description=?, audience=?, chart=?, insight=?
		WHERE id=?
		`
	_, err = repo.db.ExecContext(ctx, query, asset.Type, asset.Description, audienceData, chartData, InsightData, id)
	return err
}

func (repo *assetRepository) DeleteAssetByID(ctx context.Context, id string) error {

	query := `
	DELETE FROM assets WHERE id=?
	`
	_, err := repo.db.ExecContext(ctx, query, id)
	return err
}
