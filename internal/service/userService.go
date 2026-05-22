package service

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
	"tracker/internal/config"
	"tracker/internal/models"
	"tracker/internal/repository"
)

type userService struct {
	repo *repository.Repository
	cfg  *config.Config
}

func NewUserService(repo *repository.Repository, cfg *config.Config) *userService {
	return &userService{
		repo: repo,
		cfg:  cfg,
	}
}

type UserRequest struct {
	Username string
	Password string
}

type UserResponse struct {
	Id           string `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (svc userService) Register(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	user, err := svc.repo.User.GetUserByUsername(ctx, req.Username)
	if err == nil && user != nil { // login user
		responseUser, err := svc.Login(ctx, req)
		if err != nil {
			return nil, err
		}
		return responseUser, nil
	}
	// register user
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost) // register user
	if err != nil {
		return nil, err
	}

	refreshToken, err := svc.generateRefreshToken()
	if err != nil {
		return nil, err
	}
	refreshTokenHash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)

	user = &models.User{
		Id:                     uuid.New().String(),
		Name:                   req.Username,
		PasswordHash:           string(passwordHash),
		RefreshTokenHash:       string(refreshTokenHash),
		RefreshTokenExpiryTime: time.Now().Add(svc.cfg.Token.RefreshTokenDuration),
	}

	accessToken, err := svc.generateAccessToken(user.Id)
	if err != nil {
		return nil, err
	}

	_, err = svc.repo.User.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	responseUser := UserResponse{
		Id:           user.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &responseUser, nil

}

func (svc userService) Login(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	user, err := svc.repo.User.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("error when getting user by username: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("error when comparing password: %w", err)
	}

	responseUser, err := svc.UpdateRefreshToken(ctx, user.Id)

	return responseUser, nil
}

func (svc userService) GetUserById(ctx context.Context, id string) (*models.User, error) {
	user, err := svc.repo.User.GetUserById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error when getting user by id: %w", err)
	}
	return user, nil
}

func (svc userService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := svc.repo.User.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("error when getting user by username: %w", err)
	}
	return user, nil
}

func (svc userService) UpdateRefreshToken(ctx context.Context, userId string) (*UserResponse, error) {
	user, err := svc.repo.User.GetUserById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("error when getting user by id: %w", err)
	}

	refreshToken, err := svc.generateRefreshToken()
	if err != nil {
		return nil, err
	}
	refreshTokenHash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)

	repoReq := repository.UpdateUserRefreshRequest{
		UserID:       userId,
		RefreshToken: string(refreshTokenHash),
		ExpiryTime:   time.Now().Add(svc.cfg.Token.RefreshTokenDuration),
	}

	err = svc.repo.User.UpdateRefreshToken(ctx, repoReq)
	if err != nil {
		return nil, fmt.Errorf("error when updating refresh token: %w", err)
	}
	accessToken, err := svc.generateAccessToken(user.Id)
	if err != nil {
		return nil, err
	}

	responseUser := UserResponse{
		Id:           user.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &responseUser, nil
}

func (svc userService) DeleteUser(ctx context.Context, userId string) error {
	err := svc.repo.User.DeleteUser(ctx, userId)
	if err != nil {
		return fmt.Errorf("error when deleting refresh token: %w", err)
	}
	return nil
}

func (svc userService) generateAccessToken(id string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(svc.cfg.Token.AccessTokenDuration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(svc.cfg.Token.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("token signing error: %w", err)
	}

	return tokenString, nil
}

func (svc userService) generateRefreshToken() (string, error) {
	refreshToken := uuid.New().String()

	return refreshToken, nil
}
