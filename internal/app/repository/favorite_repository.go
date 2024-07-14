package repository

import (
	"context"
	"database/sql"
	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
	"log"
)

type FavoriteRepository interface {
	GetFavoritesByUserID(ctx context.Context, userID string, limit, offset int) ([]*models.Favorite, error)
	CreateFavorite(ctx context.Context, favorite *models.Favorite) error
	UpdateFavorite(ctx context.Context, userID, assetID string, favorite *models.Favorite) error
	DeleteFavorite(ctx context.Context, userID, assetID string) error
}

type favoriteRepository struct {
	db *sql.DB
}

func NewFavoriteRepository(db *sql.DB) FavoriteRepository {
	return &favoriteRepository{db: db}
}

func (r *favoriteRepository) GetFavoritesByUserID(ctx context.Context, userID string, limit, offset int) ([]*models.Favorite, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		query := `
			SELECT user_id, asset_id, asset_type, description
			FROM favorites
			WHERE user_id = ?
			LIMIT ? OFFSET ?`

		log.Print("user   ", userID, "   li   ", limit, "   off   ", offset)

		rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var favorites []*models.Favorite
		for rows.Next() {
			var favorite models.Favorite
			if err := rows.Scan(&favorite.UserID, &favorite.AssetID, &favorite.AssetType, &favorite.Description); err != nil {
				return nil, err
			}
			favorites = append(favorites, &favorite)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}

		return favorites, nil
	}
}

func (r *favoriteRepository) CreateFavorite(ctx context.Context, favorite *models.Favorite) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		query := `
			INSERT INTO favorites (user_id, asset_id, asset_type, description)
			VALUES (?, ?, ?, ?)`
		_, err := r.db.ExecContext(ctx, query, favorite.UserID, favorite.AssetID, favorite.AssetType, favorite.Description)
		return err
	}
}

func (r *favoriteRepository) UpdateFavorite(ctx context.Context, userID, assetID string, favorite *models.Favorite) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		query := `
			UPDATE favorites
			SET description = ?
			WHERE user_id = ? AND asset_id = ?`
		_, err := r.db.ExecContext(ctx, query, favorite.Description, userID, assetID)
		return err
	}
}

func (r *favoriteRepository) DeleteFavorite(ctx context.Context, userID, assetID string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		query := "DELETE FROM favorites WHERE user_id = ? AND asset_id = ?"
		_, err := r.db.ExecContext(ctx, query, userID, assetID)
		return err
	}
}
