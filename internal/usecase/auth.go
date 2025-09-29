package usecase

import (
	"context"
	"errors"
	"time"

	"wallet/internal/entity"
	"wallet/internal/repo"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserRepo interface {
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Create(ctx context.Context, u *entity.User) (int64, error)
}

type RefreshTokenRepo interface {
	Store(ctx context.Context, token *entity.RefreshToken) error
}

type TokenService interface {
	GenerateAccessToken(userID int64, duration time.Duration) (string, error)
	GenerateRefreshToken(userID int64, duration time.Duration) (string, error)
}

type AuthUsecase struct {
	userRepo         UserRepo
	refreshTokenRepo RefreshTokenRepo
	tokenService     TokenService
}

func NewAuthUsecase(userRepo UserRepo, refreshTokenRepo RefreshTokenRepo, tokenService TokenService) *AuthUsecase {
	return &AuthUsecase{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		tokenService:     tokenService,
	}
}

func (a *AuthUsecase) Login(ctx context.Context, username, password string) (accessToken string, refreshToken string, err error) {
	user, err := a.userRepo.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, repo.ErrUserNotFound) {
			return "", "", ErrInvalidCredentials
		}
		return "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", "", ErrInvalidCredentials
	}

	accessToken, err = a.tokenService.GenerateAccessToken(user.ID, time.Minute*15)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = a.tokenService.GenerateRefreshToken(user.ID, time.Hour*24*7)
	if err != nil {
		return "", "", err
	}

	rt := &entity.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	}

	if err := a.refreshTokenRepo.Store(ctx, rt); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *AuthUsecase) Register(ctx context.Context, u *entity.User) (int64, error) {
	if existingUser, _ := a.userRepo.GetByUsername(ctx, u.Username); existingUser != nil {
		return 0, errors.New("username already taken")
	}

	id, err := a.userRepo.Create(ctx, u)
	if err != nil {
		return 0, err
	}

	return id, nil
}
