package repository

import (
	"context"
	"database/sql"
	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}
type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (user_id , username, password,firstname, lastname)
		VALUES ($1, $2, $3, $4, $5)	`
	_, err := repo.db.Exec(query, query, user.Username, user.Password, user.FirstName, user.LastName)
	return err
}

func (repo *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `
		SELECT user_id, username, password, firstname, lastname
		FROM users
		WHERE users.username = ?
	`
	row := repo.db.QueryRow(query, username)
	var user models.User
	err := row.Scan(&user.UserID, &user.Username, &user.Password, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
