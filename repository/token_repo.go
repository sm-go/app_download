package repository

import (
	"app-download/ds"
	"app-download/model"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type TokenRepository struct {
	RDB *redis.Client
}

func NewTokenRepository(ds *ds.DataSource) *TokenRepository {
	return &TokenRepository{
		RDB: ds.RDB,
	}
}

func (r *TokenRepository) StoreRefreshToken(ctx context.Context, user *model.User, token *model.RefreshTokenData) error {
	key := fmt.Sprintf("user:%s:%v", token.ID, user.Id)
	if err := r.RDB.Set(ctx, key, 0, token.ExpiresIn).Err(); err != nil {
		return err
	}
	return nil
}

func (r *TokenRepository) DeleteRefreshToken(ctx context.Context, user *model.User, token string) error {
	key := fmt.Sprintf("user:%s:%v", token, user.Id)
	if err := r.RDB.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
