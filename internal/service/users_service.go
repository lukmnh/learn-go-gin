package service

import (
	"context"
	"errors"
	"todo_rest_api/internal/apperrors"
	"todo_rest_api/internal/dto"
	"todo_rest_api/internal/models"
	"todo_rest_api/internal/repository"
	"todo_rest_api/internal/util/jwt"

	"golang.org/x/crypto/bcrypt"
)

type UsersService interface {
	Create(ctx context.Context, req dto.UserCreateRequest) (*dto.UserCreateResponse, error)
	Login(ctx context.Context, req dto.UserLoginRequest) (*dto.UserLoginResponse, error)
}

type usersService struct {
	repo repository.UsersRepository
}

func NewUsersService(repo repository.UsersRepository) UsersService {
	return &usersService{repo: repo}
}

func (svc *usersService) Create(ctx context.Context, req dto.UserCreateRequest) (*dto.UserCreateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	hashedPass, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, apperrors.ErrHashPassword
	}

	users, err := svc.repo.Create(
		ctx,
		req.Email,
		string(hashedPass),
	)

	if err != nil {
		return nil, err
	}

	return toCreateResponse(users), nil
}

func (svc *usersService) Login(ctx context.Context, req dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	users, err := svc.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return nil, apperrors.ErrInvalidCredentials
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(req.Password)); err != nil {
		return nil, apperrors.ErrInvalidCredentials
	}

	token, err := jwt.GenerateToken(users.ID, users.Email)
	if err != nil {
		return nil, err
	}

	return &dto.UserLoginResponse{Token: token}, nil
}

func toCreateResponse(m *models.Users) *dto.UserCreateResponse {
	return &dto.UserCreateResponse{
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
	}
}
