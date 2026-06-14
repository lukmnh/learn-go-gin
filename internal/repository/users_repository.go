package repository

import (
	"context"
	"errors"
	"todo_rest_api/internal/apperrors"
	"todo_rest_api/internal/models"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type UsersRepository interface {
	Create(ctx context.Context, email string, password string) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db: db}
}

const pgUniqueViolation = "23505"

func (r *usersRepository) Create(ctx context.Context, email string, password string) (*models.Users, error) {
	user := models.Users{
		Email:    email,
		Password: password,
	}

	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgUniqueViolation {
			return nil, apperrors.ErrDuplicateEmail
		}
		return nil, err
	}

	return &user, nil
}

func (r *usersRepository) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	var user models.Users

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}
