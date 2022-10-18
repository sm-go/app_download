package service

import (
	"app-download/config"
	"app-download/model"
	"app-download/repository"
	"app-download/utils"
	"context"
)

type TokenService struct {
	tokenRepo *repository.TokenRepository
}

type TSConfig struct {
	TokenRepo *repository.TokenRepository
}

func NewTokenService(c *TSConfig) *TokenService {
	return &TokenService{
		tokenRepo: c.TokenRepo,
	}
}

func (s *TokenService) GenerateTokenPairs(ctx context.Context, user *model.User, prevToken string) (*model.TokenPair, error) {
	if prevToken != "" {
		if err := s.tokenRepo.DeleteRefreshToken(ctx, user, prevToken); err != nil {
			return nil, err
		}
	}

	accessToken, err := utils.GenerateUserAccessToken(user, config.PrivateKey)
	if err != nil {
		return nil, err
	}

	refreshData, err := utils.GenerateUserRefreshToken(user, config.RefreshSecret)
	if err != nil {
		return nil, err
	}

	if err := s.tokenRepo.StoreRefreshToken(ctx, user, refreshData); err != nil {
		return nil, err
	}

	return &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshData.SS,
	}, nil
}
