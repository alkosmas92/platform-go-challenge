package models

type Favorite struct {
	UserID      string `json:"user_id"`
	AssetID     string `json:"asset_id"`
	AssetType   string `json:"asset_type"` // "chart", "insight", "audience"
	Description string `json:"description"`
}
